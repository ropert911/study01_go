package 日志

import (
	"fmt"
)

var (
	log LoggingClient
)

func initLog(serverName string, logLevel string, EnableRemote bool) {
	var logTarget string

	if EnableRemote {
		logTarget = "http://192.168.3.2/log/" + serverName
		fmt.Println("EnableRemote is true, using remote logging service")
	} else {
		logTarget = serverName + ".log"
		fmt.Println("EnableRemote is false, using local log file")
	}
	log = NewClient(serverName, false, logTarget, logLevel)
}

func main() {
	initLog("receiver", "INFO", false)

	log.Info("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
