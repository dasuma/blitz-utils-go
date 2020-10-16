package logs

import (
	"github.com/dasuma/blitz-utils-go/utils"
	"github.com/logdna/logdna-go/logger"
)

var (
	api    = utils.GetEnv("API", "deev")
	Logger *logger.Logger
)

//Start start
func Start() {
	options := logger.Options{}
	options.Level = "fatal"
	options.App = api
	Logger, _ = logger.NewLogger(options, "e038de664579453d90769648448b0e6b")

}
