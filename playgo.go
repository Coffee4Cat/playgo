package main

import (
	"log"
	"flag"
	"playgo/user_interface"
	"playgo/audio_player"
	"playgo/structures"
	"time"
)



func main() {
	audiopath := flag.String("f", "/home/milosz/Music", "Absolute path to the folder with music")
	flag.Parse()
	
	log.Println("[PLAYGO][VERSION 1.0]")
	time.Sleep(time.Millisecond * 100)
	log.Println("[PLAYGO][PATH:" + *audiopath + "]")
	time.Sleep(time.Millisecond * 100)
	log.Println("[PLAYGO][HAVE FUN!]")
	time.Sleep(time.Millisecond * 500)
	
	// DIRECTORY SEARCH
	var folders []structures.AudioFolder = audio_player.ListAudioFolders(*audiopath)
	for i, _ := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles(*audiopath,&folders[i])
		folders[i].UpdateAudioFiles(files)
	}
	


	var commandChannel chan structures.PlayerCommand
	
	commandChannel = make(chan structures.PlayerCommand)
	
	
	audio_player.InitializePlayer(commandChannel)
	user_interface.InitializeUI(&folders, commandChannel)

}