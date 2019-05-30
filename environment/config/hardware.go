package config

// HardwareConfig configuration for the hardware
type HardwareConfig struct {
	GPIO   GPIO
	Serial Serial
}

type GPIO struct {
	Pumps []Pump
}

// Pump config for a pump
type Pump struct {
	RelayPin        uint
	SensorPin       uint
	FlowLPM         float64
	PrimeTimeMillis int
}

type Serial struct {
	PortName           string
	BaudRate           uint
	RetryDelaySeconds  uint
	BufferSize         uint
	SensorUpdateMillis uint
}

func LoadHardwareConfig() (HardwareConfig, error) {
	cfg := HardwareConfig{}

	err := LoadTOMLFile("config.hardware.toml", &cfg)

	return cfg, err
}
