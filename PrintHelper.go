package main

import (
	"fmt"
	"strings"
	"os/exec"
)

func DisplayNotification(data map[string]interface{}) {
	fmt.Println(data["played_song"])
	exec.Command("notify-send", "-i", "/", "Radio Student NowPlaying", data["played_song"].(string)).Run()
}

func PrintAllSongs(songs []map[string]interface{}) {
	for _, v := range songs {
		time := strings.Split(v["played_time"].(string), " ")[1]
		fmt.Println(time, " - ", v["played_song"])
	}
}
