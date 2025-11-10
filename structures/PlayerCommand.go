package structures

const (
	ActionSetPlay         = "set_play"
	ActionSetTrack        = "set_track"
	ActionSetLevel        = "set_level"
	ActionSetPlayFeedback = "set_play_feedback"
	ActionSetMode         = "set_mode"
)

type PlayerCommand struct {
	Action string

	Track *AudioFile
	Level bool
	Play  bool
	Mode  int
}
