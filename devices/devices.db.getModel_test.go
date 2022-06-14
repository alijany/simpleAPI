package devices

import (
	"testing"
)

func TestGetModel(t *testing.T) {
	device, err := getModel(mockConfig, "id1")

	if err != nil || device == nil {
		t.Error("error in getModel", err, device)
	}
}
