package with_alias_import

import alias "github.com/spf13/viper"

type config struct{}

var C config

func withAliasImport() {
  alias.Sub("key.subkey")                  // want "avoid usage of viper.Sub as it doesn't work well with environment variables"
  _ = alias.UnmarshalKey("key.subkey", &C) // want "avoid usage of viper.UnmarshalKey as it doesn't work well with environment variables"
}
