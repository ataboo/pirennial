package reader

import (
	"testing"

	"github.com/ataboo/pirennial/hardware/sensors/sensor"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/sensors/connection"
)

func TestSensorReaderUnmarshaling(t *testing.T) {
	table := []struct {
		rawJson string
		valid   bool
		inputPins    []uint
		values  []int
	}{
		{
			rawJson: `{"0":1, "2":3}`,
			valid:   true,
			inputPins: []uint{0, 2},
			values: []int{1, 3},
		},
		// Not parsable
		{
			rawJson: `"0":1, "2":3}`,
			valid:   false,
		},
		// Pin mismatch
		{
			rawJson: `{"0":1, "5":3}`,
			valid:   false,
			inputPins: []uint{0, 2},
		},
	}

	mockConn := connection.CreateConnectionMock()
	serialCfg := config.Serial{
		BufferSize: 1000,
	}

	CreateSensorReaderSerial(serialCfg, mockConn)
	reader := SensorReaderSerial{
		cfg:        serialCfg,
		connection: mockConn,
		buffer:     make([]byte, serialCfg.BufferSize),
	}
	var sensors []sensor.Sensor

	for _, row := range table {
		mockConn.GetResponse = []byte(row.rawJson)
		sensors = make([]sensor.Sensor, len(row.inputPins))

		for i, inputPin := range row.inputPins {
			sensors[i] = sensor.CreateSoilSensorSerial(inputPin)
			sensors[i].Data().Value = -1
		}

		err := reader.Update(sensors)
		if !row.valid {
			if err == nil {
				t.Error("expected row to throw error")
			}

			continue
		}

		if err != nil {
			t.Error("no error should be thrown")
		}

		for i, sensor := range sensors {
			if sensor.Data().Value != row.values[i] {
				t.Errorf("mismatched value for pin %d: %d - %d", sensor.Data().InputPin, row.values[i], sensor.Data().Value)
			}
		}
	}
}
