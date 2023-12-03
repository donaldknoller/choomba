package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	SQLITE_PATH = "sqlite-path"
)

func InitConfig() {
	envPrefix := "CHOOMBA"
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "-"))
	viper.AutomaticEnv()
	prepareEnv()
}

func prepareEnv() {
	// Currently this snippet does not work with viper lib
	// https://github.com/spf13/viper/issues/188
	envPrefix := "CHOOMBA"
	underScoreReplacer := strings.NewReplacer("_", "-")
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(underScoreReplacer)
	viper.AutomaticEnv()
	// unfortunate workaround
	for _, envEntry := range os.Environ() {
		envVar := strings.Split(envEntry, "=")
		if strings.HasPrefix(envVar[0], envPrefix) {
			stripPrefix := fmt.Sprintf("%s_", envPrefix)
			newKey := strings.TrimPrefix(envVar[0], stripPrefix)
			newKey = underScoreReplacer.Replace(newKey)
			viper.Set(newKey, envVar[1])
		}
	}
}
