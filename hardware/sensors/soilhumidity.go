package sensors

import "time"

type SerialReturn struct {
	Sensors  []SoilHumidity
	LastRead time.Time
}

type SoilHumidity struct {
	Pin   uint `json:"pin"`
	Value uint `json:"value"`
}
