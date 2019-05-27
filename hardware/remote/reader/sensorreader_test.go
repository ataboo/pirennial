package reader

import (
	"testing"

	"github.com/ataboo/pirennial/hardware/remote/sensor"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/connection"
)

func TestSensorReaderUnmarshaling(t *testing.T) {
	table := []struct {
		rawJson string
		valid   bool
		cfgs    []config.SoilSensor
		values  []int
	}{
		{
			rawJson: `{"0":1, "2":3}`,
			valid:   true,
			cfgs: []config.SoilSensor{
				{InputPin: 0},
				{InputPin: 2},
			},
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
			cfgs: []config.SoilSensor{
				{InputPin: 0},
				{InputPin: 2},
			},
		},
	}

	mockConn := connection.CreateConnectionMock()
	cfg := config.Serial{
		BufferSize: 1000,
	}

	CreateSensorReaderSerial(cfg, mockConn)
	reader := SensorReaderSerial{
		cfg:        cfg,
		connection: mockConn,
		buffer:     make([]byte, cfg.BufferSize),
	}
	var sensors []sensor.Sensor

	for _, row := range table {
		mockConn.GetResponse = []byte(row.rawJson)
		cfg.SoilSensors = row.cfgs
		sensors = make([]sensor.Sensor, len(cfg.SoilSensors))

		for i, sensorCfg := range cfg.SoilSensors {
			sensors[i] = sensor.CreateSoilSensorSerial(sensorCfg)
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
