// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"config"
	"engine"
	"fmt"
	"github.com/nporsche/np-golang-logging"
	"handler"
	"log/syslog"
	"net/http"
	"os"
	"runtime"
)

func main() {
	config.Init()
	/* logger initiailization */
	backend, err := logging.NewSyslogBackendPriority("bean_game", syslog.LOG_LOCAL3)
	if err != nil {
		fmt.Printf("logging init error=[%s]", err.Error())
		os.Exit(1)
	}
	format := logging.MustStringFormatter(
		"%{color}[%{module}.%{shortfunc}][%{level:.4s}]%{color:reset}%{message}",
	)
	logging.SetBackend(logging.NewBackendFormatter(backend, format))
	/* end logger initiailization */
	runtime.GOMAXPROCS(config.This.Processors)

	handlers := map[string]http.HandlerFunc{
		"/player/report":   handler.PlayerReport,
		"/bean/manipulate": handler.BeanManipulate,
		"/clean":           handler.Clean,
	}

	engine.Run(handlers)
}
