package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"github.com/t-tomalak/logrus-easy-formatter"
)

var (
	Log *logrus.Entry
)

func init() {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% %logger_name% %lvl% %filename% %msg%\n",
		},
	}
	Log = logger.WithFields(logrus.Fields{"filename": "tcf","logger_name": "tcf"})
}



