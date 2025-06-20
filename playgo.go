package main

import (
	"fmt"
	// "playgo/user_interface"
	"playgo/audio_player"
	"playgo/structures"
	"time"
	"sync/atomic"
)



func main() {
	// user_interface.InitializeUI();
	// FILE SCAN (yep. i will put it into sum function in some time...)
	var folders []structures.AudioFolder = audio_player.ListAudioFolders("/home/milosz/Music")
	for i, folder := range folders {
		var files []structures.AudioFile = audio_player.ListAudioFiles("/home/milosz/Music",folder)
		folders[i].UpdateAudioFiles(files)
		fmt.Println(folders[i].Repr())
		for _, file := range folders[i].AudioFiles{
			fmt.Println(file.Repr())
		}
	}
	


	// POINTER-FLAG STUFF TEST
	var volume atomic.Value
	var pause atomic.Value
	volume.Store(0.1) 
	pause.Store(true)


	audio_player.PlayAudioFile(folders[0].AudioFiles[0], &volume, &pause)
	time.Sleep(1 * time.Second)
	volume.Store(0.05)
	pause.Store(false)
	time.Sleep(3 * time.Second)
	pause.Store(true)
	time.Sleep(5 * time.Second)


}