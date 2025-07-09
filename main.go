package main

import (
	"fem_project/internal/app"	
)

func main(){
	app, err := app.NewApplication()
	if (err != nil) {
		panic(err)
	}

	app.Logger.Println("we are running our app")
}
