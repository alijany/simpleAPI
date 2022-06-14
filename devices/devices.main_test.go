package devices

import (
	"testing"
)

func TestHandler(t *testing.T) {
	_, err := Handler(mockConfig)

	if err != nil {
		t.Error("error in putDevice", err)
	}
}
