package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tuvok"
	app.Usage = "Run tuvok"
	app.Action = healthcheck
}

func healthcheck(context *cli.Context) {
	for {
		// healthy, err := Healthcheck()
		// panicIfError(err)
		//
		// if !healthy {
		// 	err = Unregister()
		// } else {
		// 	err = Register()
		// }
		// panicIfError(err)
	}
}

func panicIfError(err error) {
	if err != nil {
		log.Fatalf("Error occured: %v", err.Error())
	}
}
