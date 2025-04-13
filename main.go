package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	osInterruptChan := make(chan os.Signal, 1)
	signal.Notify(osInterruptChan, syscall.SIGINT, syscall.SIGTERM)

	var userPoolWaitGroup sync.WaitGroup
	userPool := CONFIG.Int("app.user-pool")
	for i := range userPool {
		userPoolWaitGroup.Add(1)
		go user.CreateUser(i, ctx, &userPoolWaitGroup)
	}

	go func() {
		<-osInterruptChan
		log.Info().Msg("os interrupt recieved, sending signal to workers to stop immidietly...")
		cancel()
	}()
	userPoolWaitGroup.Wait()
	log.Info().Msg("all workers finished, exiting program normally")
}
