package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
	"strings"
	"time"
)

var NewConfig = fx.Provide(newConfig)

type IConfig interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetDuration(key string) time.Duration
}

type config struct {
	cfg *viper.Viper
}

func newConfig() IConfig {

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("json")

	cfg.AddConfigPath(getConfigPath())

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg.WatchConfig()

	return &config{cfg: cfg}
}

func (c *config) Get(key string) interface{} {
	return c.cfg.Get(key)
}

func (c *config) GetBool(key string) bool {
	return c.cfg.GetBool(key)
}

func (c *config) GetFloat64(key string) float64 {
	return c.cfg.GetFloat64(key)
}

func (c *config) GetInt(key string) int {
	return c.cfg.GetInt(key)
}

func (c *config) GetIntSlice(key string) []int {
	return c.cfg.GetIntSlice(key)
}

func (c *config) GetString(key string) string {
	return c.cfg.GetString(key)
}

func (c *config) GetStringSlice(key string) []string {
	return c.cfg.GetStringSlice(key)
}

func (c *config) GetDuration(key string) time.Duration {
	return c.cfg.GetDuration(key)
}

func getConfigPath() (path string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	slice := strings.Split(wd, "b2b-api")
	path = slice[0] + "b2b-api/internal/config"
	return
}
