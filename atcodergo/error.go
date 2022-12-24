package atcodergo

import (
	"fmt"
)

type NeedAuthError struct {
	msg string
}

func newNeedAuthError(funcName string) *NeedAuthError {
	return &NeedAuthError{msg: fmt.Sprintf("%s needs login. must call after Client.Login().", funcName)}
}

func (e *NeedAuthError) Error() string {
	return e.msg
}
