package serviceconfiguration

import (
	"fmt"
	"konekko.me/xbasis/commons/config"
)

var configuration *xbasisconfig.GosionConfiguration

func Configuration() xbasisconfig.OnGosionConfigurationChanged {
	return func(config *xbasisconfig.GosionConfiguration) {
		fmt.Println("service update new config ok")
		configuration = config
	}
}

func Get() *xbasisconfig.GosionConfiguration {
	if configuration == nil {
		panic("error gosion configuration.")
	}
	return configuration
}
