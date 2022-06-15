package devices

import (
	"testing"
)

func TestPutItem(t *testing.T) {
	err := putItem(mockConfig, &Device{
		Id:          "/devices/id1",
		DeviceModel: "/devicemodels/id1",
		Name:        "Sensor",
		Note:        "Testing a sensor.",
		Serial:      "A020000102",
	})
	if err != nil {
		t.Error("error in putItem", err)
	}
	// check for invalid DeviceModel
	err = putItem(mockConfig, &Device{
		Id:          "/devices/id1",
		DeviceModel: "/devicemodels/id2",
		Name:        "Sensor",
		Note:        "Testing a sensor.",
		Serial:      "A020000102",
	})
	if err == nil {
		t.Error("error in putItem", err)
	}
}
