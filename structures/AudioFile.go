package structures

import (
	"fmt"
	"os"
	// "path/filepath"
	// "strings"
	"github.com/hajimehoshi/go-mp3"
	// "time"
)


type AudioFile struct {
	FullPath string
	Name string
	DataLength int64
	SampleRate int
	Duration int
	Decoder *mp3.Decoder
}

func NewAudioFile(name string, fullpath string) AudioFile {
	f, err := os.Open(fullpath)
	if err != nil { panic(err) }

	decoder, err := mp3.NewDecoder(f)
	if err != nil { panic(err) }
	
	var data_length int64 = decoder.Length()
	var sample_rate int = decoder.SampleRate()
	var duration int = int(float64(data_length) / float64(sample_rate * 4))

	return AudioFile{FullPath: fullpath, Name: name, Duration: duration,
					DataLength: data_length, SampleRate: sample_rate, Decoder: decoder}
}

func (audiofile *AudioFile) UpdateData(new_name string, new_fullpath string){
	*audiofile = NewAudioFile(new_name, new_fullpath)
}


func (audiofile *AudioFile) Repr() string {
	var ret string = fmt.Sprintf("--- %s   :   %d seconds",audiofile.Name, audiofile.Duration)
	return ret
}
