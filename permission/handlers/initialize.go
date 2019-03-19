package permission_handers

import "konekko.me/gosion/commons/config"

func Initialize() gs_commons_config.OnConfigNodeChanged {
	return func(config *gs_commons_config.GosionInitializeConfig) {

	}
}
