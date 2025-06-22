package structures

import (
	"fmt"
	"os"
)


type AudioFolder struct {
	Entry os.DirEntry
	AudioFiles []AudioFile
	Duration int
}

func NewAudioFolder(entry os.DirEntry) AudioFolder {
	return AudioFolder{Entry: entry}
}

func (audiofolder *AudioFolder) UpdateEntry(new_entry os.DirEntry){
	audiofolder.Entry = new_entry
}

func (audiofolder *AudioFolder) UpdateAudioFiles(new_audio_files []AudioFile){
	audiofolder.AudioFiles = new_audio_files
}

func (audiofolder *AudioFolder) Repr() string {
	var ret string = fmt.Sprintf("+ %-60s track count: %-20d total time: %d",audiofolder.Entry.Name(),len(audiofolder.AudioFiles),audiofolder.Duration)
	return ret
}



