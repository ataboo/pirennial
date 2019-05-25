package sensors

import (
	"encoding/json"
	"testing"
)

func TestSerialParsing(t *testing.T) {
	raw := []byte(`[{"pin":0,"value":10},{"pin":1,"value":11},{"pin":2,"value":12},{"pin":3,"value":13}]`)

	var out []SoilHumidity

	err := json.Unmarshal(raw, &out)
	if err != nil {
		t.Error(err)
	}

	t.Logf("out: %+v", out)
}
