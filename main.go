package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"runtime"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Weather struct {
	Water  int
	Wind   int
	Status string
}

var weather Weather

func weatherData() {
	for {
		weather.Water = rand.Intn(20)
		weather.Wind = rand.Intn(20)

		//Condition Weather Status
		if weather.Water <= 5 || weather.Wind <= 6 {
			weather.Status = "Aman"
		}
		if (weather.Water >= 6 && weather.Water <= 8) || (weather.Wind >= 7 && weather.Wind <= 15) {
			weather.Status = "Siaga"
		}
		if weather.Water >= 8 || weather.Wind >= 15 {
			weather.Status = "Bahaya"
		}
		fmt.Println(weather.Water, weather.Wind, weather.Status)
		structWeather, _ := json.Marshal(weather)
		ioutil.WriteFile("weather.json", structWeather, os.ModePerm)
		time.Sleep(2 * time.Second)
	}
}

func handlerWeather(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(weather.Status))
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// fmt.Println(err)
		// os.Exit(1)
		return
	}
	// fmt.Println(tmpl)

	err = tmpl.Execute(w, weather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// fmt.Println(err)
	}
}

func main() {
	fmt.Println("Total Go routine", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	go weatherData()
	r := mux.NewRouter()
	r.HandleFunc("/", handlerWeather)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
