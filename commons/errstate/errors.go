package errstate

import "konekko.me/gosion/commons/dto"

var (
	Success                     = &gs_commons_dto.State{Ok: false, Code: 0, Message: "success"}
	ErrRequest                  = &gs_commons_dto.State{Ok: false, Code: 20001, Message: "err request"}
	ErrInvalidStateField        = &gs_commons_dto.State{Ok: false, Code: 20002, Message: "err return type, no found 'State' file"}
	ErrRoleAlreadyExists        = &gs_commons_dto.State{Ok: false, Code: 20003, Message: "err role already exists"}
	ErrGroupAlreadyExists       = &gs_commons_dto.State{Ok: false, Code: 20004, Message: "err group already exists"}
	ErrBlacklistAlreadyExists   = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err blacklist already exists"}
	ErrApiAlreadyExists         = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err api already exists"}
	ErrFunctionBindGroupId      = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err no found function groupId"}
	ErrFunctionAuthType         = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err function auth type"}
	ErrApplicationAlreadyExists = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err application already exists"}
	ErrApplicationClosed        = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err application closed"}
)
