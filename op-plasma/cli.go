package plasma

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/urfave/cli/v2"
)

const (
	EnabledFlagName    = "plasma.enabled"
	DaRedisUrlFlagName = "plasma.redis-url"
)

func plasmaEnv(envprefix, v string) []string {
	return []string{envprefix + "_PLASMA_" + v}
}

func CLIFlags(envPrefix string, category string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     EnabledFlagName,
			Usage:    "Enable plasma mode",
			Value:    false,
			EnvVars:  plasmaEnv(envPrefix, "ENABLED"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     DaRedisUrlFlagName,
			Usage:    "Redis URL",
			EnvVars:  plasmaEnv(envPrefix, "REDIS_URL"),
			Category: category,
		},
	}
}

type CLIConfig struct {
	Enabled    bool
	DARedisURL string
}

func (c CLIConfig) Check() error {
	if !c.Enabled {
		return fmt.Errorf("plasma must be enabled")
	}

	if c.DARedisURL == "" {
		return fmt.Errorf("DA server URL is required when plasma da is enabled")
	}

	return nil
}

func (c CLIConfig) NewDAClient() *DAClient {
	return &DAClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     c.DARedisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func ReadCLIConfig(c *cli.Context) CLIConfig {
	return CLIConfig{
		Enabled:    c.Bool(EnabledFlagName),
		DARedisURL: c.String(DaRedisUrlFlagName),
	}
}
