package application

import (
	"net/http"
)

type Application struct {
	Server Server
}

type Server struct {
	Port int
	Mux *http.ServeMux
}

func (app *Application) Init() (Application, error) {
	app.Server = Server {
		Port : 8080,
		Mux : http.NewServeMux(),
	}

	return *app, nil
}
