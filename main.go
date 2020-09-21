package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/reujab/wallpaper"
)

/// ~ DEFINE VARIABLES ~ ////
var (
	layout  string
	running bool
	picture string
	path    string
)

func main() {

	//// ~ SET VARIABLES ~ ////
	layout = "15:04"
	running = true
	path = "C:\\Users\\User\\Documents\\goAnimatedScreen\\source"

	//// ~ on for loop ~ ////
	for running {

		//// ~ go get file ~ ////
		files, err := ioutil.ReadDir("./source")
		if err != nil {
			log.Fatal(err)
		}

		//// ~ change picture ~ ////
		for _, file := range files {
			picture = path + "\\" + file.Name()
			wallpaper.SetFromFile(picture)
			fmt.Println("Wallpaper changed by " + picture)

		}
	}
}
