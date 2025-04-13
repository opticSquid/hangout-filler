package main

import (
	"github.com/knadh/koanf/v2"
	"hangoutsb.in/filler-service/appconfigmanager"
	"hangoutsb.in/filler-service/loggerconfig"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var CONFIG = koanf.New(".")

func main() {
	loggerconfig.InitLogger()
	appconfigmanager.InitAppConfig(CONFIG)
	loggerconfig.SetGlobalLogLevel(CONFIG)
}
