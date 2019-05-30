package relay

type Relay interface {
	IsOn() bool
	PinNumber() uint
	On() error
	Off() error
	Cleanup()
}
