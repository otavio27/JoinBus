package ttktools

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//DBStats ..
type DBStats struct {
	Reads   int
	Writes  int
	Deletes int
}

//TTKDB ...
type ttkcdb struct {
	module   string
	Server   string
	Password string
	Control  string
	Stats    DBStats
	ctx      context.Context
	log      *ttklog
}

// TTKDSC ...
type TTKDSC struct {
	Key   string
	Value string
	Rev   string `json:"_rev,omitempty"`
}

// //DBList ..
// type DBList struct {
// 	Next string
// 	Data interface{}
// }

// CDB ...
func CDB(ctx context.Context, Server string, Password string, TablePrefix string, logger *ttklog) *ttkcdb {
	var err error
	t := ttkcdb{}
	t.module = "ttkcdb"
	t.ctx = ctx
	t.log = logger
	t.Server = Server
	t.Password = Password
	t.Control = TablePrefix + "_control"
	if err != nil {
		// t.log.Msg(ctx, t.module, err.Error(), DEBUG)
	}
	return &t
}

func (t *ttkcdb) CtrlPut(key string, value string) {
	data := TTKDSC{}
	err := t.Read(t.Control, key, &data)
	if err != nil {
		// t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	data.Key = key
	data.Value = value
	err = t.Write(t.Control, key, &data)
	if err != nil {
		// t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	t.Stats.Writes++
}

func (t *ttkcdb) CtrlGet(key string) string {
	ret := TTKDSC{}
	err := t.Read(t.Control, key, &ret)
	if err != nil {
		// t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	t.Stats.Reads++
	return ret.Value
}

// THIS IS CORRECT AND WORKING
func (t *ttkcdb) Write(kind string, key string, data interface{}) error {
	type httpRet struct {
		Ok  bool   `json:"ok"`
		Rev string `json:"rev"`
	}
	ret := httpRet{}
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(d))
	req, err := http.NewRequest("PUT", t.Server+"/"+kind+"/"+key, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+t.Password)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode == 401 {
		t.log.Msg(t.ctx, t.module, "Authentication error", DEBUG)
		return fmt.Errorf("%s", "Authentication error")
	}
	if err != nil {
		return err
	}
	t.Stats.Writes++
	if res.StatusCode == 404 {
		req, _ = http.NewRequest("PUT", t.Server+"/"+kind, nil)
		req.Header.Add("Authorization", "Basic "+t.Password)
		res, err := http.DefaultClient.Do(req)
		t.Stats.Writes++
		if err != nil {
			t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		}
		// body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode == 201 {
			payload = strings.NewReader(string(d))
			req, err = http.NewRequest("PUT", t.Server+"/"+kind+"/"+key, payload)
			req.Header.Add("Authorization", "Basic "+t.Password)
			req.Header.Add("Content-Type", "application/json")
			// req.Header.Add("cache-control", "no-cache")
			res, err = http.DefaultClient.Do(req)
			t.Stats.Writes++
			// body, _ := ioutil.ReadAll(res.Body)
			if err != nil {
				t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
				return err
			}
		}
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	_, np := reflect.TypeOf(data).Elem().FieldByName("Rev")
	if np {
		reflect.ValueOf(data).Elem().FieldByName("Rev").Set(reflect.ValueOf(ret.Rev))
	}
	if res.StatusCode > 299 {
		t.log.Msg(t.ctx, t.module, strconv.Itoa(res.StatusCode)+" - "+string(body), DEBUG)
	} else {
		t.log.Msg(t.ctx, t.module, "Record written to database", DEBUG)
	}
	return nil
}

func (t *ttkcdb) Delete(kind string, key string) error {
	type httpRet struct {
		Rev string `json:"_rev"`
	}
	ret := httpRet{}
	req, err := http.NewRequest("GET", t.Server+"/"+kind+"/"+key, nil)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	req.Header.Add("Authorization", "Basic "+t.Password)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	t.Stats.Reads++
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	req, err = http.NewRequest("DELETE", t.Server+"/"+kind+"/"+key+"?rev="+ret.Rev, nil)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	res, err = http.DefaultClient.Do(req)
	t.Stats.Deletes++
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	defer res.Body.Close()
	return nil
}

func (t *ttkcdb) Read(kind string, key string, data interface{}) error {
	if len(key) == 0 {
		return fmt.Errorf("Invalid key")
	}
	req, err := http.NewRequest("GET", t.Server+"/"+kind+"/"+key, nil)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	req.Header.Add("Authorization", "Basic "+t.Password)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	t.Stats.Reads++
	// ttklog.msg(ctx, t.module, "["+kind+":"+key+"]"+res.Status, DEBUG)
	if res.StatusCode == 401 || res.StatusCode == 404 {
		// t.log.Msg(t.ctx, t.module, res.Status, DEBUG)
		return fmt.Errorf(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return err
	}
	return nil
}

func (t *ttkcdb) List(kind string, dst interface{}, pageSize int, cursorStr string, order string, filter ...interface{}) (string, error) {
	type Ret struct {
		Docs           []interface{} `json:"docs"`
		Bookmark       string        `json:"bookmark"`
		ExecutionStats struct {
			TotalKeysExamined       int     `json:"total_keys_examined"`
			TotalDocsExamined       int     `json:"total_docs_examined"`
			TotalQuorumDocsExamined int     `json:"total_quorum_docs_examined"`
			ResultsReturned         int     `json:"results_returned"`
			ExecutionTimeMs         float32 `json:"execution_time_ms"`
		} `json:"execution_stats"`
		Warning string `json:"warning"`
	}
	ret := Ret{}
	fx := ""
	selector := "{\"selector\": {"
	for i := 0; i < len(filter); i++ {
		sel := make([]string, 0)
		if strings.Contains(filter[i].(string), "[]") { // It is an array
			vet := strings.Split(filter[i].(string), "[]")
			sel = strings.Split(vet[1][1:], " ")
			if len(sel) > 1 {
				if sel[1] == "=" {
					fx = "$eq"
				}
			} else {
				fx = "$regex"
			}
			selector += "\"" + vet[0] + "\" : {\"$elemMatch\": {\"" + sel[0] + "\": {\"" + fx + "\": " + t.strwrap(filter[i+1]) + "}}}"
		} else { // It is not an array
			sel = strings.Split(filter[i].(string), " ")

			if len(sel) > 1 {
				if sel[1] == "=" {
					fx = "$eq"
				}
			} else {
				fx = "$regex"
			}
			selector += "\"" + sel[0] + "\" : { \"" + fx + "\": " + t.strwrap(filter[i+1]) + "}"
		}
		i++
		if i+1 < len(filter) {
			selector += ","
		}
	}
	selector += "},"
	limit := ""
	if pageSize > 0 {
		limit = "\"limit\": " + strconv.Itoa(pageSize) + ","
	}
	bookmark := "\"bookmark\": \"" + cursorStr + "\", \"skip\": 0,\"execution_stats\": true"
	query := selector + limit + bookmark + "}"
	payload := strings.NewReader(query)
	req, _ := http.NewRequest("POST", t.Server+"/"+kind+"/_find", payload)
	req.Header.Add("Authorization", "Basic "+t.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	t.Stats.Reads++
	if res.StatusCode != 200 {
		return "", fmt.Errorf(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}
	r, _ := json.Marshal(ret.Docs)
	err = json.Unmarshal(r, dst)
	if err != nil {
		return "", err
	}
	return ret.Bookmark, nil
}

func (t *ttkcdb) strwrap(x interface{}) string {
	if reflect.TypeOf(x).String() == "int" {
		return strconv.Itoa(x.(int))
	}
	return "\"" + x.(string) + "\""
}
