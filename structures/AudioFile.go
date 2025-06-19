package structures

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	// "github.com/hajimehoshi/go-mp3"
)


type AudioFile struct {
	Entry os.DirEntry
	Name string
}

func NewAudioFile(entry os.DirEntry) AudioFile {
	var name string = strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
	// decoder, _ := mp3.NewDecoder()

	return AudioFile{Entry: entry, Name: name}
}

func (audiofile *AudioFile) UpdateEntry(new_entry os.DirEntry){
	*audiofile = NewAudioFile(audiofile.Entry)
}


func (audiofile *AudioFile) Repr() string {
	var ret string = fmt.Sprintf("--- %s",audiofile.Name)
	return ret
}