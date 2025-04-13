package user

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
)

func CreateUser(userNumber int, ctx context.Context, userWaitGroup *sync.WaitGroup) {

	log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker started")
	defer userWaitGroup.Done()
	defer log.Debug().Str("worker-type", "user").Int("worker-number", userNumber).Msg("Worker finished and exiting.")
	select {
	case <-ctx.Done():
		log.Info().Str("woker-type", "user").Int("worker-number", userNumber).Msg("stopping due to context cancellation")
	default:
		log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker doing its job")
	}
	log.Info().Str("worker-type", "user").Int("worker-number", userNumber).Msg("Worker completed all its tasks naturally.")
}
