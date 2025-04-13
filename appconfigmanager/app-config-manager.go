package appconfigmanager

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func InitAppConfig(k *koanf.Koanf) {
	// Loading base config
	err := k.Load(file.Provider("resources/application.yaml"), yaml.Parser())
	if err != nil {
		log.Fatal().Msg("error loading base configuration")
	}
	// Loading specific config based on flag
	profile := flag.String("profile", "default", "sets a particular profile that overrides base configuration")
	flag.Parse()
	err = k.Load(file.Provider("resources/application-"+*profile+".yaml"), yaml.Parser())
	if err != nil {
		log.Error().Str("active profile", *profile).Msg("error loading profile configuration")
	}
	log.Info().Str("active profile", *profile).Msg("configurations Loaded")
}
