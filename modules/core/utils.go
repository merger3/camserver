package core

type Command struct {
	User    string
	Channel string
	Command string `query:"command"`
}

type ClickedCam struct {
	Found    bool   `json:"found"`
	Name     string `json:"cam"`
	Position int    `json:"position"`
}
