package encode

import (
	"testing"
)

func TestJPEGEncoder(t *testing.T) {

	err := testEncoder("jpeg://")

	if err != nil {
		t.Fatal(err)
	}
}
