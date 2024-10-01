package core

type AuthHeaders struct {
	User  string `header:"X-User-Name"`
	Token string `header:"X-Twitch-Token"`
}

type Command struct {
	Channel string
	User    string `header:"X-User-Name"`
	Command string `query:"command"`
}

type ClickedCam struct {
	Found    bool   `json:"found"`
	Name     string `json:"cam"`
	Position int    `json:"position"`
}

type CamRequest struct {
	AuthHeaders
	Cam string `json:"camera"`
}
