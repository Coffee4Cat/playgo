package audio_player

import (
	"fmt"
	"io"
	"playgo/structures"
	"sync"

	"github.com/hajimehoshi/oto"
)

var playbutton_mutex sync.Mutex
var playspeed int = 44100

func InitializePlayer(command chan structures.PlayerCommand, feedback chan structures.PlayerCommand) {
	var audiofile *structures.AudioFile
	var ctx *oto.Context
	var player *oto.Player
	var playing bool
	var volume float64 = 0.5

	go func() {
		buf := make([]byte, 4096)
		ctx, _ = oto.NewContext(playspeed, 2, 2, 4096) // Well... Cannot think bout better solution right now
		for {
			select {
			case cmd := <-command:
				switch cmd.Action {
				case structures.ActionSetPlay:
					playbutton_mutex.Lock()
					playing = !playing
					playbutton_mutex.Unlock()
				case structures.ActionSetTrack:
					playbutton_mutex.Lock()
					if player != nil {
						player.Close()
					}
					if audiofile != nil {
						audiofile.CurrentTime = 0
					}
					audiofile = cmd.Track
					audiofile.ResetDecoder()
					player = ctx.NewPlayer()
					playing = true
					playbutton_mutex.Unlock()
				case structures.ActionSetLevel:
					if cmd.Level {
						volume += 0.1
					} else {
						volume -= 0.1
					}

					if volume < 0.0 {
						volume = 0.0
					} else if volume > 1.0 {
						volume = 1.0
					}
				}

			default:
				if playing && ctx != nil && audiofile != nil {
					n, err := audiofile.Decoder.Read(buf)
					if n > 0 {
						for i := 0; i < n; i += 2 {
							sample := int16(buf[i]) | int16(buf[i+1])<<8
							adjusted := int16(float64(sample) * volume)
							buf[i] = byte(adjusted & 0xFF)
							buf[i+1] = byte((adjusted >> 8) & 0xFF)
						}
						player.Write(buf[:n])
						audiofile.CurrentTime += int(float64(n) * 1000 / (4.0 * float64(playspeed)))
						if audiofile.CurrentTime > audiofile.Duration {
							audiofile.CurrentTime = audiofile.Duration
						}
					}
					if err == io.EOF {
						playbutton_mutex.Lock()
						playing = false
						playbutton_mutex.Unlock()
						audiofile.CurrentTime = 0
						feedback <- structures.PlayerCommand{Action: structures.ActionSetPlayFeedback, Play: playing}
						continue
					}
					if err != nil {
						fmt.Println("Decoder error:", err)
						playbutton_mutex.Lock()
						playing = false
						playbutton_mutex.Unlock()
						feedback <- structures.PlayerCommand{Action: structures.ActionSetPlayFeedback, Play: playing}
					}
				}

			}
		}
	}()

}
