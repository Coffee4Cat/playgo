package main

import (
	"fmt"
	// "playgo/user_interface"
	"playgo/audio_player"
	"playgo/structures"
)



func main() {
	// user_interface.InitializeUI();
	var folders []structures.AudioFolder = audio_player.ListAudioFolders("/home/milosz/Music")
	for _, folder := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles("/home/milosz/Music",folder)
		folder.UpdateAudioFiles(files)
		fmt.Println(folder.Repr())
		for _, file := range files{
			fmt.Println(file.Repr())
		}
	}


}