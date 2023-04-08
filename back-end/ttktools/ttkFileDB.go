package ttktools

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Memdb ...
type filedb struct {
	memdb map[string][]byte
	log   *ttklog
	ctx   context.Context
}

// FileDB returns a MemDB instance
func FileDB() *filedb {
	t := filedb{}
	t.memdb = make(map[string][]byte)
	return &t
}

// Write a key in the memory database
func (t *filedb) Write(key string, value []byte) {
	t.memdb[key] = value
}

// Read a key from the memory database
func (t *filedb) Read(key string) []byte {
	return t.memdb[key]
}

// Delete a key from memory database
func (t *filedb) Delete(key string) {
	delete(t.memdb, key)
}

func (t *filedb) List() []string {
	ret := make([]string, 0)
	for k := range t.memdb {
		ret = append(ret, k)
	}
	return ret
}

func (t *filedb) Init(ctx context.Context, path string, ext string, log *ttklog) error {
	c := 0
	if log == nil {
		t.log.Msg(t.ctx, "FILEDB", "Log instance inexistent", DEBUG)
		return fmt.Errorf("%s", "Log instance inexistent")
	}
	t.log = log
	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.log.Msg(t.ctx, "FILEDB", err.Error(), DEBUG)
		return fmt.Errorf("%s", err.Error())
	}
	log.Msg(t.ctx, "FILEDB", "Loading files", DEBUG)
	for _, f := range files {
		if filepath.Ext(f.Name()) == "."+ext {
			c++
			key := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			f, _ := ioutil.ReadFile(path + "/" + f.Name())
			t.Write(key, f)
		}
	}
	log.Msg(t.ctx, "FILEDB", fmt.Sprintf("%d files loaded in memory", c), DEBUG)
	return nil
}
