package main

import (
	"flag"
	"time"
)

func main() {
	var help = flag.Bool("h", false, "lists all commands")
	var notify = flag.Bool("n", false, "show notification for now playing every 210secs")

	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		return
	}

	for {
		jsonFromLink := GetJSONFromLink()
		songs := saveJSONToDataStruct(jsonFromLink)

		if *notify == true {
			DisplayNotification(songs[0])
			time.Sleep(210 * time.Second)
		} else {
			PrintAllSongs(songs)
			return
		}
	}
}
