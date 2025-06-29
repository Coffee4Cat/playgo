package audio_player

import (
	"github.com/hajimehoshi/oto"
	"fmt"
	"io"
	"playgo/structures"
	"time"
)


var totalBytesRead int64 = 0


func InitializePlayer(command chan structures.PlayerCommand, feedback chan structures.PlayerCommand) {
	var audiofile *structures.AudioFile
	var ctx *oto.Context
	var player *oto.Player
	var playing bool
	var volume float64 = 0.5


	go func() {
		buf := make([]byte, 4096)
		ctx, _ = oto.NewContext(44100, 2, 2, 5000) // Well... Cannot think bout better solution right now
		for {
			select {
			case cmd := <-command:
				switch cmd.Action {
				case structures.ActionSetPlay:
					playing = !playing
				case structures.ActionSetTrack:
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
				        totalBytesRead += int64(n)
						for i := 0; i < n; i += 2 {
							sample := int16(buf[i]) | int16(buf[i+1])<<8
							adjusted := int16(float64(sample) * volume)
							buf[i] = byte(adjusted & 0xFF)
							buf[i+1] = byte((adjusted >> 8) & 0xFF)
						}
						player.Write(buf[:n])
						audiofile.CurrentTime += int(float64(n) * 1000 / (4.0 * 44100.0))
						if audiofile.CurrentTime > audiofile.Duration {
							audiofile.CurrentTime = audiofile.Duration
						}
					}
					if err == io.EOF {
						playing = false
						audiofile.CurrentTime = 0
						feedback <- structures.PlayerCommand{Action: structures.ActionSetPlayFeedback, Play: playing}
						continue
					}
					if err != nil {
						fmt.Println("Decoder error:", err)
						playing = false
						feedback <- structures.PlayerCommand{Action: structures.ActionSetPlayFeedback, Play: playing}
					}
				} else {
					time.Sleep(10 * time.Millisecond)
					playing = false
					// feedback <- structures.PlayerCommand{Action: structures.ActionSetPlayFeedback, Play: playing}
					totalBytesRead = 0
				}
			}
		}
	}()

}