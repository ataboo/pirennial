package sensors

import "time"

type SerialReturn struct {
	Sensors  []SoilHumidity
	LastRead time.Time
}

type SoilHumidity struct {
	Raw uint16
}
