package pump

type RelayMock struct {
	isOn      bool
	pinNumber uint
}

func CreateRelayMock(pinNumber uint) Relay {
	r := RelayMock{
		isOn:      false,
		pinNumber: pinNumber,
	}

	return &r
}

func (r *RelayMock) On() error {
	r.isOn = true

	return nil
}

func (r *RelayMock) Off() error {
	r.isOn = false

	return nil
}

func (r *RelayMock) IsOn() bool {
	return r.isOn
}

func (r *RelayMock) PinNumber() uint {
	return r.pinNumber
}

func (r *RelayMock) Cleanup() {
	//
}
