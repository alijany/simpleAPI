package devices

import (
	"testing"
)

func TestGetItem(t *testing.T) {
	// Test normal action
	device, err := getItemById[Device](mockConfig, deviceTableName, "id1")
	if err != nil || device == nil {
		t.Error("error in getItem", err, device)
	}
	// Test empty result
	empty, err := getItemById[struct{}](mockConfig, "Empty", "id1")
	if err != nil || empty != nil {
		t.Error("error in getItem", err, empty)
	}
	// test table not found
	notFound, err := getItemById[struct{}](mockConfig, "NotFound", "id1")
	if err == nil || notFound != nil {
		t.Error("error in getItem", err, notFound)
	}
	// test invalid generic
	gen, err := getItemById[string](mockConfig, deviceTableName, "id1")
	if err == nil || gen != nil {
		t.Error("error in getItem", err, notFound)
	}
}
