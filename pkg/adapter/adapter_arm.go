// +build arm

package adapter

import "gobot.io/x/gobot/platforms/raspi"

// New creates a new i2c adapter suitable for the raspberry pi.  The device
// should be set to an empty string as it's not needed.
func New(device string) (*raspi.Adaptor, error) {
	return raspi.NewAdaptor(), nil
}