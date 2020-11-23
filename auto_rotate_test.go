package bimg

import (
	"testing"
)

func TestAutoRotate(t *testing.T) {
	t.Run("For JPEG with orientation 6 (rotate 90), it autorotates the image", func(t *testing.T) {
		buf, _ := imageBuf("rotate_90.jpg")
		infoInput, _ := Metadata(buf)
		options := Options{
			NoAutoRotate: false,
			Flop: false,
		}
		outBuf, err := AutoRotate(buf, options)
		if err != nil {
			t.Errorf("Error in rotating JPEG images %v", err)
		}
		if len(outBuf) == 0 {
			t.Errorf("Buff cannot be empty")
		}
		infoOutput, _ := Metadata(outBuf)
		// For the given image, infoInput.Orientation is 6 which means the image should be rotated 90 deg.
		// Hence checking if the final output buf's orientation and input's buf orientation is different
		if infoOutput.Orientation == infoInput.Orientation {
			t.Errorf("Orientation is still the same! Error in image rotation %d", infoOutput.Orientation)
		}
	})
}
