package structures

import (
	"fmt"
	"os"
)

type AudioFolder struct {
	Entry         os.DirEntry
	AudioFiles    []AudioFile
	Duration      int
	AutoplayIndex int
}

func NewAudioFolder(entry os.DirEntry) AudioFolder {
	return AudioFolder{Entry: entry, AutoplayIndex: 0}
}

func (audiofolder *AudioFolder) UpdateEntry(new_entry os.DirEntry) {
	audiofolder.Entry = new_entry
}

func (audiofolder *AudioFolder) UpdateAudioFiles(new_audio_files []AudioFile) {
	audiofolder.AudioFiles = new_audio_files
}

func (audiofolder *AudioFolder) GetNextTrack() *AudioFile {
	audiofolder.AutoplayIndex = (audiofolder.AutoplayIndex + 1) % len(audiofolder.AudioFiles)
	return &audiofolder.AudioFiles[audiofolder.AutoplayIndex]
}

func (audiofolder *AudioFolder) Repr() string {
	var ret string = fmt.Sprintf("+ %-30s Tracks Count: %-20d Total Duration: %02d:%02d", audiofolder.Entry.Name(), len(audiofolder.AudioFiles), int(float64(audiofolder.Duration)/60000.0), int(float64(audiofolder.Duration)/1000)%60)
	return ret
}
