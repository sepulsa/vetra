package with_viper

import "github.com/spf13/viper"

type config struct{}

var C config

func withViper() {
  viper.Sub("key.subkey")                  // want "avoid usage of viper.Sub as it doesn't work well with environment variables"
  _ = viper.UnmarshalKey("key.subkey", &C) // want "avoid usage of viper.UnmarshalKey as it doesn't work well with environment variables"
}
