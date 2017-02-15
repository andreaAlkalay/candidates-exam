package main

import (
	"github.com/cloudfoundry-community/go-cfclient"
	//"strconv"
	"github.com/go-martini/martini"
	"cloudapps/service"
	//"cloudapps/model"
	//"cloudapps/helpers"
	"math/rand"
	"encoding/json"
	"time"
	"sort"
)

type NameSorter []cfclient.App

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }


func shuffle(arr []cfclient.App) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond())) // no shuffling without this line
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {

	m := martini.Classic()

	m.Get("/apps", func() ([]byte, error){
		client := service.GetClient()
		apps, _ := client.ListApps()
		return json.MarshalIndent(apps, "", "  ")
	})

	m.Get("/apps/sort?shuffle", func() ([]byte, error){
		client := service.GetClient()
		apps, _ := client.ListApps()
		shuffledList := apps
		shuffle(shuffledList)
		return json.MarshalIndent(shuffledList, "", "  ")
	})

	m.Get("/apps/sort?byName", func() ([]byte, error){
		client := service.GetClient()
		apps, _ := client.ListApps()
		sortedList := apps
		sort.Sort(NameSorter(sortedList))
		return json.MarshalIndent(sortedList, "", "  ")
	})

	m.Get("/", func() string {
		current_time := time.Now().Local()
		str:= "Returning a warm welcome :-) !! The current date and time is " + string(current_time.Format("2006-01-02 15:04:05 +0800"))
		return str
	})

	m.Run()

}
