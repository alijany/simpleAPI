package devices

import (
	"testing"
)

func TestCreateModelTable(t *testing.T) {
	channel := make(chan error)
	go createModelTable(mockConfig, channel)
	err := <-channel
	if err != nil {
		t.Error("error in getDevice", err)
	}
}
