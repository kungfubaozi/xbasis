package serviceconfiguration

import (
	"fmt"
	"konekko.me/gosion/commons/config"
)

var configuration *gs_commons_config.GosionConfiguration

func Configuration() gs_commons_config.OnGosionConfigurationChanged {
	return func(config *gs_commons_config.GosionConfiguration) {
		fmt.Println("service update new config ok")
		configuration = config
	}
}

func Get() *gs_commons_config.GosionConfiguration {
	if configuration == nil {
		panic("error gosion configuration.")
	}
	return configuration
}
