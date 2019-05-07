package flowstate

import (
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/workflow/types"
)

var (
	NextFlow                  = &gs_commons_dto.State{Ok: true, Code: types.NFNextFlow}
	FlowScriptTrue            = &gs_commons_dto.State{Ok: true, Code: types.NFScriptTrue}
	FlowScriptFalse           = &gs_commons_dto.State{Ok: true, Code: types.NFScriptFalse}
	ErrInvalidInstance        = &gs_commons_dto.State{}
	ErrUnsupportedConnectType = &gs_commons_dto.State{}
	ErrNode                   = &gs_commons_dto.State{}
	ErrSubmitFormFieldNil     = &gs_commons_dto.State{}
	ErrSubmitFormFieldType    = &gs_commons_dto.State{}
	ErrSubmitFormFieldValue   = &gs_commons_dto.State{}
	ErrSubmitFormFieldRegex   = &gs_commons_dto.State{}
)
