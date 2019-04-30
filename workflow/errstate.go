package workflow

import "konekko.me/gosion/commons/dto"

var (
	ErrNoOperatingConditions   = &gs_commons_dto.State{}
	ErrScript                  = &gs_commons_dto.State{}
	ErrNoStartEvent            = &gs_commons_dto.State{}
	ErrNoEndEvent              = &gs_commons_dto.State{}
	ErrProcessName             = &gs_commons_dto.State{}
	ErrProcessKey              = &gs_commons_dto.State{}
	ErrProcessExists           = &gs_commons_dto.State{}
	ErrInstanceAlreadyFinished = &gs_commons_dto.State{}
	ErrInvalid                 = &gs_commons_dto.State{}
)
