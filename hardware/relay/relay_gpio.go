package relay

import (
	"fmt"

	"github.com/brian-armstrong/gpio"
)

// Relay representation of a relay driving a pump
type RelayGPIO struct {
	pin       *gpio.Pin
	isOn      bool
	pinNumber uint
}

func CreateRelayGPIO(pinNumber uint) Relay {
	pin := gpio.NewOutput(pinNumber, false)

	r := RelayGPIO{
		pin:       &pin,
		isOn:      false,
		pinNumber: pinNumber,
	}

	return &r
}

func (r *RelayGPIO) IsOn() bool {
	return r.isOn
}

func (r *RelayGPIO) PinNumber() uint {
	return r.pinNumber
}

// On turn the relay on
func (r *RelayGPIO) On() error {
	return r.setOn(true)
}

// Off turn the relay off
func (r *RelayGPIO) Off() error {
	return r.setOn(false)
}

func (r *RelayGPIO) setOn(on bool) error {
	r.isOn = on

	if r.pin == nil {
		return fmt.Errorf("pin on relay %d not innitialized")
	}

	if on {
		return r.pin.High()
	}

	return r.pin.Low()
}

// Cleanup release the pin
func (r *RelayGPIO) Cleanup() {
	if r.pin == nil {
		return
	}

	r.pin.Low()
	r.pin.Close()
}
