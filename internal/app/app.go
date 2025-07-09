package app

import ( 
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger 
}

func NewApplication (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: Logger,
	}

	return app, nil
}
