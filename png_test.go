package encode

import (
	"testing"
)

func TestPNGEncoder(t *testing.T) {

	err := testEncoder("png://")

	if err != nil {
		t.Fatal(err)
	}
}
