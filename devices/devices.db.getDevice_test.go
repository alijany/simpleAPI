package devices

import (
	"testing"
)

func TestGetDevice(t *testing.T) {
	device, err := getDevice(mockConfig, "id1")

	if err != nil || device == nil {
		t.Error("error in getDevice", err, device)
	}
}
