package modules

import "konekko.me/gosion/workflow/flowerr"

type IStore interface {
	Set(status int64, keys map[string]interface{}) (bool, *flowerr.Error)

	Get(keys map[string]interface{}) (int64, *flowerr.Error)

	Clear(keys map[string]interface{}) *flowerr.Error
}
