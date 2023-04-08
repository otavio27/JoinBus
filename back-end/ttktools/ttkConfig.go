/*
 MODULE: ttkconfig.go
 AUTHOR: Leo Schneider <schleo@outlook.com>
 DATE  : August 2017
 INFO  : This module handles application parameters via command line or config file
*/

package ttktools

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

// ttkconfig ...
type ttkconfig struct {
	cfg map[string]interface{}
}

// Config ..
func Config(cfgfile string) *ttkconfig {
	t := ttkconfig{}
	t.init(cfgfile)
	return &t
}

func (t *ttkconfig) init(cfgfile string) {
	t.cfg = make(map[string]interface{})
	config := flag.String("c", cfgfile, "Alternate configuration")
	listcfg := flag.Bool("l", false, "List Configuration")
	flag.Parse()
	t.Set("config", *config)
	err := json.Unmarshal([]byte(FileRead(*config)), &t.cfg)
	if err != nil {
		PrintError("Configuration file not found [" + *config + "]")
		os.Exit(1)
	}
	for key, value := range t.cfg {
		t.Set(key, value)
	}
	if *listcfg {
		t.List()
	}
}

// Get ...
func (t *ttkconfig) Get(key string) interface{} {
	if t.cfg[key] == nil {
		PrintError("Configuration not found [" + key + "]")
		os.Exit(1)
	}
	return t.cfg[key]
}

// GetString ...
func (t *ttkconfig) GetString(key string) string {
	return t.Get(key).(string)
}

// GetBool ...
func (t *ttkconfig) GetBool(key string) bool {
	return t.Get(key).(bool)
}

// GetInt ...
func (t *ttkconfig) GetInt(key string) int {
	return t.Get(key).(int)
}

// GetInt32 ...
func (t *ttkconfig) GetInt32(key string) int32 {
	return t.Get(key).(int32)
}

// GetInt64 ...
func (t *ttkconfig) GetInt64(key string) int64 {
	return t.Get(key).(int64)
}

// GetFloat32 ...
func (t *ttkconfig) GetFloat32(key string) float32 {
	return t.Get(key).(float32)
}

// GetFloat64 ...
func (t *ttkconfig) GetFloat64(key string) float64 {
	return t.Get(key).(float64)
}

// Set ...
func (t *ttkconfig) Set(key string, value interface{}) {
	t.cfg[key] = value

}

// List ...
func (t *ttkconfig) List() {
	var keys []string
	for k := range t.cfg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		PrintData(CYAN, fmt.Sprintf("%-19s", k), fmt.Sprintf("%s", t.cfg[k]))
	}
	fmt.Println("")
}
