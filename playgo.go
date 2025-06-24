package main

import (
	"fmt"
	"playgo/user_interface"
	"playgo/audio_player"
	"playgo/structures"
	// "time"
)



func main() {

	// DIRECTORY SEARCH
	var folders []structures.AudioFolder = audio_player.ListAudioFolders("/home/milosz/Music")
	for i, _ := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles("/home/milosz/Music",&folders[i])
		folders[i].UpdateAudioFiles(files)
		fmt.Println(folders[i].Repr())
		for _, file := range folders[i].AudioFiles{
			fmt.Println(file.Repr())
		}
	}
	
	var commandChannel chan structures.PlayerCommand
	commandChannel = make(chan structures.PlayerCommand)

	audio_player.InitializePlayer(commandChannel)
	user_interface.InitializeUI(&folders, commandChannel)
	// commandChannel <- structures.PlayerCommand{ Action: structures.ActionSetTrack, Track: &folders[0].AudioFiles[0] }
	// time.Sleep(time.Second * 1)
	// commandChannel <- structures.PlayerCommand{ Action: structures.ActionSetTrack, Track: &folders[1].AudioFiles[0] }
	// time.Sleep(time.Second * 1)


}