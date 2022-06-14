package devices

import (
	"testing"
)

func TestPutDevice(t *testing.T) {
	err := putDevice(mockConfig, &Device{
		Id:          "/devices/id1",
		DeviceModel: "/devicemodels/id1",
		Name:        "Sensor",
		Note:        "Testing a sensor.",
		Serial:      "A020000102",
	})

	if err != nil {
		t.Error("error in putDevice", err)
	}
}
