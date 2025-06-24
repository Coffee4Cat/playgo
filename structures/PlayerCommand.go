package structures


const (
	ActionSetPlay = "set_play"
	ActionSetTrack = "set_track"
	ActionSetLevel = "set_level"
)

type PlayerCommand struct {
	Action string

	Track *AudioFile
	Level bool
	Play bool
}






