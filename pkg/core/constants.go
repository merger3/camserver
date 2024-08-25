package core

var (
	VideoWidth  = float64(1920)
	VideoHeight = float64(1080)
)

type Command struct {
	Channel string
	Command string `query:"command"`
}
