package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
)

type application struct {
	config config
}

type config struct {	
	port int
}

func main() {
	var config config
	config.port = 8080;

	app := &application{
		config: config,
	}

	app.serve()
}

func (app *application) serve() (error) {
	
	server := &http.Server{
		Addr: fmt.Sprintf(`:%d`, app.config.port),
	}

	iris := iris.New().Run(iris.Server(server))

	return iris
}
