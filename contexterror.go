package contexterror

import (
	"encoding/json"
	"fmt"
)

type Context map[string]interface{}

type ContextError struct {
	ctx Context
}

func (e ContextError) Error() string {
	message, err := e.ctx.GetJsonString()
	if err != nil {
		return fmt.Sprintf("Error: GetJsonString %v", e.ctx)
	}
	return message
}

func (e *ContextError) Init(ctx Context, prev error) {
	err, ok := prev.(ContextError)
	if ok {
		e.ctx = ctx.Merge(err.ctx)
	} else {
		e.ctx = ctx.Merge(Context{
			"message": prev.Error(),
		})
	}
}

func NewContextError(ctx Context, prev error) ContextError {
	err := new(ContextError)
	err.Init(ctx, prev)
	return *err
}

func (ctx Context) Merge(ctx2 Context) Context {
	for k, v := range ctx2 {
		ctx[k] = v
	}
	return ctx
}

func (ctx Context) GetJsonString() (string, error) {
	bytes, err := json.Marshal(ctx)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
