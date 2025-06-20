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
	for i, folder := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles("/home/milosz/Music",folder)
		folders[i].UpdateAudioFiles(files)
		fmt.Println(folders[i].Repr())
		for _, file := range folders[i].AudioFiles{
			fmt.Println(file.Repr())
		}
	}

}