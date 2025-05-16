package socket

type SocketManager struct {
	Aliases     map[string]string
	CommonNames map[string]string
}

func NewSocketManager() *SocketManager {
	return &SocketManager{}
}
