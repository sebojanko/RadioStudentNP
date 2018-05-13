package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var help = flag.Bool("h", false, "lists all commands")
	var show_all = flag.Bool("all", false, "shows last 25 songs")

	flag.Parse()

	if *help == true {
		fmt.Println("-h\tfor showing this message")
		fmt.Println("-all\tfor showing last 25 songs")
		return
	}

	for {
		jsonFromLink := GetJSONFromLink()
		songs := saveJSONToDataStruct(jsonFromLink)

		if *show_all == true {
			PrintAllSongs(songs)
			return
		}

		DisplayNotification(songs[0])
		time.Sleep(4 * time.Minute)
	}
}
