package logs

import (
	"github.com/dasuma/blitz-utils-go/utils"
	"github.com/evalphobia/go-logdna/logdna"
)

var (
	api = utils.GetEnv("API", "deev")
	//Logger
	Logger *logdna.Client
)

//Start start
func Start() {
	conf := logdna.Config{
		APIKey:       "e038de664579453d90769648448b0e6b",
		App:          api,
		Env:          "production",
		MinimumLevel: logdna.LogLevelInfo,
		Sync:         false,
		Debug:        true,
	}
	Logger, _ = logdna.New(conf)

}
