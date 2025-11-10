package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"playgo/audio_player"
	"playgo/structures"
	"playgo/user_interface"
	"time"
)

func main() {
	home, _ := os.UserHomeDir()
	audiopath := flag.String("f", filepath.Join(home, "Music"), "Absolute path to the folder with music")
	speed := flag.Float64("s", 1.0, "Audio speed multiplier")
	flag.Parse()

	log.Println("[PLAYGO][VERSION 1.0]")
	time.Sleep(time.Millisecond * 100)
	log.Println("[PLAYGO][PATH: " + *audiopath + "]")
	log.Println("[PLAYGO][SAMPLERATE MULTIPLIER: " + fmt.Sprintf("%.2f", *speed) + "]")
	time.Sleep(time.Millisecond * 100)
	log.Println("[PLAYGO][HAVE FUN!]")
	time.Sleep(time.Millisecond * 500)

	// DIRECTORY SEARCH
	var folders []structures.AudioFolder = audio_player.ListAudioFolders(*audiopath)
	for i, _ := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles(*audiopath, &folders[i])
		folders[i].UpdateAudioFiles(files)
	}

	var commandChannel chan structures.PlayerCommand
	var feedbackChannel chan structures.PlayerCommand

	commandChannel = make(chan structures.PlayerCommand)
	feedbackChannel = make(chan structures.PlayerCommand)

	audio_player.InitializePlayer(commandChannel, feedbackChannel, *speed)
	user_interface.InitializeUI(&folders, commandChannel, feedbackChannel)

}
