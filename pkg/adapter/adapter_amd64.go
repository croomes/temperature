// +build amd64

package adapter

import (
	"errors"
	"os"

	"gobot.io/x/gobot/platforms/firmata"
)

// ErrNoDevice is returned if no device was specified or if it doesn't exist.
var ErrNoDevice = errors.New("device not found, specify with --device FILENAME")

// New creates a new i2c adapter suitable for the arduino.  Device must be
// supplied, typically /dev/tty.usbmodem141421 or similar for Arduinos.
func New(device string) (*firmata.Adaptor, error) {

	if _, err := os.Stat(device); os.IsNotExist(err) {
		return nil, ErrNoDevice
	}

	return firmata.NewAdaptor(device), nil
}