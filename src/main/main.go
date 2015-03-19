// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"config"
	"engine"
	"github.com/nporsche/np-golang-logger"
	"handler"
	"net/http"
	"runtime"
)

func main() {
	config.Init()
	logger.Init("bean_game", config.This.Debug)
	runtime.GOMAXPROCS(config.This.Processors)

	handlers := map[string]http.HandlerFunc{
		"/player/report":   handler.PlayerReport,
		"/bean/manipulate": handler.BeanManipulate,
		"/clean":           handler.Clean,
	}

	engine.Run(handlers)
}
