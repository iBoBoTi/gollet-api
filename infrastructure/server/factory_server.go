package server

import (
	"errors"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"time"
)

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin int = iota
)

func NewWebServerFactory(instance int, port int, ctxTimeout time.Duration) (ports.Server, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
