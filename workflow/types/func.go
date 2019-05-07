package types

import "konekko.me/gosion/commons/dto"

type StateCallback func() (*gs_commons_dto.State, error)
