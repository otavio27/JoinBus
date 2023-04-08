package contracts

import (
	"context"

	"github.com/vingarcia/krest"
)

type LogProvider interface {
	Msg(ctx context.Context, module, message string, level int)
}

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
