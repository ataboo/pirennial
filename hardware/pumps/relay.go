package pumps

import (
	"runtime"

	"github.com/brian-armstrong/gpio"
	"github.com/op/go-logging"
)

var logger *logging.Logger
var gpioActive bool

func init() {
	logger = logging.MustGetLogger("pirennial")

	if !gpioActive {
		logger.Infof("GPIO functions simulated since architecture is %s", runtime.GOARCH)
	}
}

// Relay representation of a relay driving a pump
type Relay struct {
	pin       *gpio.Pin
	isOn      bool
	pinNumber uint
}

// NewRelay create a new relay
func NewRelay(pinNumber uint) *Relay {
	var pin *gpio.Pin
	if gpioActive {
		p := gpio.NewOutput(pinNumber, false)
		pin = &p
	}

	relay := Relay{
		pin:       pin,
		isOn:      false,
		pinNumber: pinNumber,
	}

	return &relay
}

func (r *Relay) PinNumber() uint {
	return r.pinNumber
}

// On turn the relay on
func (r *Relay) On() {
	r.isOn = true

	if r.pin == nil {
		return
	}

	err := r.pin.High()
	if err != nil {
		logger.Error(err.Error())
	}
}

// Off turn the relay off
func (r *Relay) Off() {
	r.isOn = false
	if r.pin == nil {
		return
	}

	err := r.pin.Low()
	if err != nil {
		logger.Error(err.Error())
	}
}

// Set the relay on or off
func (r *Relay) Set(on bool) {
	if on {
		r.On()
	} else {
		r.Off()
	}
}

// Toggle the relay on or off
func (r *Relay) Toggle() bool {
	r.Set(!r.isOn)

	return r.isOn
}

// IsOn if the relay is on
func (r *Relay) IsOn() bool {
	return r.isOn
}

// Cleanup release the pin
func (r *Relay) Cleanup() {
	if r.pin == nil {
		return
	}

	r.pin.Low()
	r.pin.Close()
}
