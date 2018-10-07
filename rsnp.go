package main

import (
	"flag"
	"time"
)

var LAST_SONG map[string]interface{}

func main() {
	var help = flag.Bool("h", false, "lists all commands")
	var notify = flag.Bool("n", false, "show notification for now playing")

	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		return
	}

	for {
		jsonFromLink := GetJSONFromLink()
		songs := SaveJSONToDataStruct(jsonFromLink)

		if *notify == true {
			if isSongNew(songs[0]) {
				DisplayNotification(songs[0])
			}
			saveSong(songs[0])
			time.Sleep(30 * time.Second)
		} else {
			PrintAllSongs(songs)
			return
		}
	}
}
func saveSong(song map[string]interface{}) {
	LAST_SONG = song
}
func isSongNew(song map[string]interface{}) bool {
	if LAST_SONG["played_song"] == song["played_song"] {
		return false
	}
	return true
}
