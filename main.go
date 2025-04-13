package main

import (
	"context"
	"time"

	"github.com/knadh/koanf/v2"
	"hangoutsb.in/filler-service/appconfigmanager"
	"hangoutsb.in/filler-service/loggerconfig"
	"hangoutsb.in/filler-service/user"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var CONFIG = koanf.New(".")

func main() {
	loggerconfig.InitLogger()
	appconfigmanager.InitAppConfig(CONFIG)
	loggerconfig.SetGlobalLogLevel(CONFIG)
	userPool := CONFIG.Int("app.user-pool")
	context, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	for i := range userPool {
		go user.CreateUser(i, context)
	}
}
