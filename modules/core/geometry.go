package core

type Geom struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	FrameWidth  float64 `json:"frameWidth"`
	FrameHeight float64 `json:"frameHeight"`
}

func (r Geom) IsPoint() bool {
	return (r.Width == 0 && r.Height == 0)
}

func (r Geom) IsLine() bool {
	return (r.Width == 0 || r.Height == 0)
}

func (r Geom) GetTopLeft() (float64, float64) {
	return r.X, r.Y
}

func (r Geom) GetMidpoint() (float64, float64) {
	if r.IsPoint() {
		return r.GetTopLeft()
	}

	xMid := r.X + (r.Width / 2)
	yMid := r.Y + (r.Height / 2)

	return xMid, yMid
}

func (r Geom) GetMeasurements() (float64, float64) {
	return r.Width, r.Height
}

func (r Geom) GetScaledCoordinates(xIn, yIn float64) (float64, float64) {
	scaleX := r.FrameWidth / VideoWidth
	scaleY := r.FrameHeight / VideoHeight

	x := xIn / scaleX
	y := yIn / scaleY

	return x, y
}

func (r Geom) GetScaledMeasurments(wIn, hIn float64) (float64, float64) {
	scaleX := r.FrameWidth / VideoWidth
	scaleY := r.FrameHeight / VideoHeight

	w := wIn / scaleX
	h := hIn / scaleY

	return w, h
}
