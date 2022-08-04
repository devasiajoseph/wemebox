package image

import "testing"

func TestWebP(t *testing.T) {
	err := ToWebP("in.jpg", "out.webp")
	if err != nil {
		t.Errorf("Error converting to webp")
	}
}
