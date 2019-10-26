// +build !arm,!amd64

package adapter

import (
	"errors"

	"gobot.io/x/gobot"
)

// New will return an error on unsupported platforms.
func New(device string) (gobot.Adaptor, error) {
	return nil, errors.New("unsupported platform")
}