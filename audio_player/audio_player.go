package audio_player

import (
	"github.com/hajimehoshi/oto"
	"fmt"
	"io"
	"playgo/structures"
	"sync/atomic"
	"time"
)



func PlayAudioFile(audiofile structures.AudioFile, volume *atomic.Value, pause *atomic.Value) {
    ctx, err := oto.NewContext(audiofile.SampleRate, 2, 2, 5000)
    if err != nil {
        fmt.Println("oto error")
    }

    player := ctx.NewPlayer()

    go func() {
        defer player.Close()
        buf := make([]byte, 4096)

        for {
			vol := volume.Load().(float64)
			pau := pause.Load().(bool)
			if pau {			
				n, err := audiofile.Decoder.Read(buf)
				if n > 0 {
					for i := 0; i < len(buf); i += 2 { 
						sample := int16(buf[i]) | int16(buf[i+1])<<8
						adjusted := int16(float64(sample) * vol)
						buf[i] = byte(adjusted & 0xFF)
						buf[i+1] = byte((adjusted >> 8) & 0xFF)
					}
					player.Write(buf[:n])
				}
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("DECODER ERROR")
					break
				}
			} else {
				time.Sleep(10 * time.Millisecond)
			}
        }
    }()
}

