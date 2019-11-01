package settings

import (
	"gopkg.in/ini.v1"
	"tcf/vars"
	"tcf/logger"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		logger.Log.Panicln(err)
	}

	vars.HTTP_HOST = Cfg.Section("").Key("HTTP_HOST").MustString("127.0.0.1")
	vars.HTTP_PORT = Cfg.Section("").Key("HTTP_PORT").MustInt(8000)

	vars.REPO_PATH = Cfg.Section("").Key("REPO_PATH").MustString("repos")
	vars.MAX_INDEXERS = Cfg.Section("").Key("MAX_INDEXERS").MustInt(vars.DefaultMaxConcurrentIndexers)
	vars.MAX_Concurrency_REPOS = Cfg.Section("").Key("MAX_Concurrency_REPOS").MustInt(100)

}
