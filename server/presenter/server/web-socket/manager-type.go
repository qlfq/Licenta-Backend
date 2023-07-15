package web_socket

import (
	"sync"
)

type Manager struct {
	Clients List
	sync.RWMutex

	Otps     RetentionMap
	Handlers map[string]EventHandler
}
