package devices

import (
	"testing"
)

func TestGetItem(t *testing.T) {
	// Test normal action
	device, err := getItemById[Device](mockConfig, deviceTableName, "/devices/id1")
	if err != nil || device == nil {
		t.Error("error in getItem", err, device)
	}
	model, err := getItemById[Device](mockConfig, modelTableName, "/devicemodels/id2")
	if err != nil || device == nil {
		t.Error("error in getItem", err, model)
	}
	// Test empty result
	empty, err := getItemById[struct{}](mockConfig, "Empty", "/devices/id1")
	if err != nil || empty != nil {
		t.Error("error in getItem", err, empty)
	}
	// test table not found
	notFound, err := getItemById[struct{}](mockConfig, "NotFound", "/devices/id1")
	if err == nil || notFound != nil {
		t.Error("error in getItem", err, notFound)
	}
	// test invalid generic
	gen, err := getItemById[string](mockConfig, modelTableName, "/devicemodels/id1")
	if err == nil || gen != nil {
		t.Error("error in getItem", err, notFound)
	}
}
