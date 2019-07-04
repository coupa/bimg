package bimg

import (
	"testing"
)

func TestImageJoin(t *testing.T) {
	t.Run("For joining png images of different width and height", func(t *testing.T) {
		var imgArray []*Image
		imgArray = append(imgArray, initImage("test.png"))
		imgArray = append(imgArray, initImage("test2.png"))
		_, err := ImageJoin(imgArray, PNG)
		if err != nil {
			t.Errorf("Issue in joining PNG images %v", err)
		}
	})
}
