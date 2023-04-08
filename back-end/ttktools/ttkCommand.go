package ttktools

import (
	"fmt"
	"reflect"
	"strings"
)

// Command is used for custom commands in form of Go Functions
type Command struct{}

type ttkcommand struct {
}

//Cmd Returns an instance of ttktools Cmd
func Cmd() *ttkcommand {
	t := ttkcommand{}
	return &t
}

// Call a function with its args
func (t *ttkcommand) Call(CmdType interface{}, cmd string, args ...interface{}) ([]reflect.Value, error) {
	cmdo := strings.ToUpper(cmd[0:1]) + strings.ToLower(cmd[1:])
	return t.run(CmdType, cmdo, args)
}

func (t *ttkcommand) run(CmdType interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	// var cmds Command
	class := reflect.ValueOf(CmdType)
	method := class.MethodByName(name)
	if !method.IsValid() {
		return result, fmt.Errorf("%s", "Method "+name+" does not exist")
	}
	if len(params) != method.Type().NumIn() {
		return result, fmt.Errorf("Incorrect parameters for function %s - Expected: %d Received: %d", name, method.Type().NumIn(), len(params))
	}
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	return method.Call(in), nil
}
