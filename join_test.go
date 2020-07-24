package bimg

import (
	"testing"
)

func TestImageJoin(t *testing.T) {
	t.Run("For joining png images of different width and height", func(t *testing.T) {
		buf, _ := imageBuf("2020141C26A074000.tif")
		_, err := ImageJoinRev(buf, 10, Options{})
		if err != nil {
			t.Errorf("Issue in joining PNG images %v", err)
		}
	})
}
