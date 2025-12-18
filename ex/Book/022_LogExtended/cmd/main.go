package main

import (
	"fmt"
	"logextended/pkg/logex"
)

func main() {
	fmt.Println("LogExtended")

	logger := logex.NewLogExtended()
	logger.SetLogLevel(logex.LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")

}
