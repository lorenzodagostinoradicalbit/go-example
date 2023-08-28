package main

import (
	"github.com/lorenzodagostinoradicalbit/go-example/config"
	"github.com/lorenzodagostinoradicalbit/go-example/service"
)

func main() {
	config.InitViper()
	service.InitServer()
}
