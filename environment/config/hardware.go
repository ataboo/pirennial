package config

// HardwareConfig configuration for the hardware
type HardwareConfig struct {
	Pumps  []Pump
	Serial Serial
}

// Pump config for a pump
type Pump struct {
	Pin             uint
	FlowLPM         float64
	PrimeTimeMillis int
}

type Serial struct {
	PortName          string
	BaudRate          uint
	RetryDelaySeconds uint
	BufferSize        uint
}

func LoadHardwareConfig() (HardwareConfig, error) {
	cfg := HardwareConfig{}

	err := LoadTOMLFile("config.hardware.toml", &cfg)

	return cfg, err
}
