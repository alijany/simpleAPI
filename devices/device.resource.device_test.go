package devices

import (
	"testing"
)

func TestCreateDeviceTable(t *testing.T) {
	channel := make(chan error)
	go createDeviceTable(mockConfig, channel)
	err := <-channel
	if err != nil {
		t.Error("error in create device table : ", err)
	}
}
