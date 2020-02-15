package encode

import (
	"testing"
)

func TestGIFEncoder(t *testing.T) {

	err := testEncoder("gif://")

	if err != nil {
		t.Fatal(err)
	}
}
