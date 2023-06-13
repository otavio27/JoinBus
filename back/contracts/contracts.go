package contracts

import (
	"context"

	"github.com/vingarcia/krest"
)

type LogProvider interface {
	Debug(ctx context.Context, title string, valueMaps ...LogBody)
	Info(ctx context.Context, title string, valueMaps ...LogBody)
	Warn(ctx context.Context, title string, valueMaps ...LogBody)
	Error(ctx context.Context, title string, valueMaps ...LogBody)
	Fatal(ctx context.Context, title string, valueMaps ...LogBody)
}

type LogBody = map[string]any

type CfgProvider interface {
	GetString(key string) string
}

type MsgProvider interface {
	Read(key string) []byte
	Write(key string, value []byte)
}

type UnpackerProvider interface {
	UnpackLines(file []byte) interface{}
	UnpackStations(file []byte) interface{}
	SendMessageError(ctx context.Context, resp krest.Response) (interface{}, string)
}
