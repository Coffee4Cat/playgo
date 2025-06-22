package audio_player

import(
	"fmt"
	"strings"
	"os"
	"path/filepath"
	// "github.com/hajimehoshi/go-mp3"
	"playgo/structures"
)




func ListAudioFolders(audiopath string) []structures.AudioFolder {
	var audiofolders []structures.AudioFolder
	entries, err := os.ReadDir(audiopath)

	if err != nil {
		fmt.Println("Couldn't open the directory")
		return audiofolders
	}

	for _,entry := range entries {
		var audiofolder structures.AudioFolder = structures.NewAudioFolder(entry)
		audiofolders = append(audiofolders, audiofolder) 
	}

	return audiofolders
}

func ListAudioFiles(audiopath string, audiofolder *structures.AudioFolder) []structures.AudioFile {
	var audiofiles []structures.AudioFile
	var fullpath string = audiopath + "/" + audiofolder.Entry.Name()
	var totaltime int = 0
	audiofolder.Duration = totaltime
	entries, err := os.ReadDir(fullpath)


	if err != nil {
		fmt.Println("Couldn't list the directory")
		return audiofiles
	}

	for _,entry := range entries {
		var name string = strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		var fp string = fullpath + "/" + entry.Name()
		var audiofile structures.AudioFile = structures.NewAudioFile(name,fp)
		fmt.Println("%d",audiofile.Duration)
		totaltime += audiofile.Duration
		audiofiles = append(audiofiles, audiofile) 
	}
	audiofolder.AudioFiles = audiofiles
	audiofolder.Duration = totaltime

	return audiofiles
}