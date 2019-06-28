package flowstate

import (
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/workflow/types"
)

var (
	NextFlow                  = &commons.State{Ok: true, Code: types.NFNextFlow}
	FlowScriptTrue            = &commons.State{Ok: true, Code: types.NFScriptTrue}
	FlowScriptFalse           = &commons.State{Ok: true, Code: types.NFScriptFalse}
	ErrInvalidInstance        = &commons.State{}
	ErrUnsupportedConnectType = &commons.State{}
	ErrNode                   = &commons.State{}
	ErrSubmitFormFieldNil     = &commons.State{}
	ErrSubmitFormFieldType    = &commons.State{}
	ErrSubmitFormFieldValue   = &commons.State{}
	ErrSubmitFormFieldRegex   = &commons.State{}
)
