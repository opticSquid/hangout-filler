package user

import (
	"context"

	"github.com/rs/zerolog/log"
)

func CreateUser(userNumber int, ctx context.Context) {

	log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker started")
	defer log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker stopped")
	select {
	case <-ctx.Done():
		log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker finished job")
	default:
		log.Debug().Str("woker-type", "user").Int("worker-number", userNumber).Msg("worker doing its job")
	}
}
