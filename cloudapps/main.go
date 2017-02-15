package main

import (
	"strconv"
	"github.com/go-martini/martini"
	"cloudapps/service"
	//"cloudapps/model"
	//"cloudapps/helpers"
)

func main() {

	m := martini.Classic()

	m.Get("/apps", func() string {
		client := service.GetClient()
		apps, _ := client.ListApps()
		return "Number of apps: " + strconv.Itoa(len(apps))
	})

	m.Get("/", func() string {
		return "Hello username!"
	})

	m.Run()

}
