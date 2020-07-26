package bimg

import (
	"fmt"
	"os"
	"testing"
)

func TestImageJoin(t *testing.T) {
	t.Run("For joining png images of different width and height", func(t *testing.T) {
		buf, _ := imageBuf("2020141C26A074000.tif")
		outBuf, err := ImageJoinNew(buf, 60, Options{})
		if err != nil {
			t.Errorf("Issue in joining PNG images %v", err)
		}

		if len(outBuf) == 0 {
			t.Errorf("Buff cannot be empty")
		}

		f, err := os.Create("output.png")
		_, err1 := f.Write(outBuf)
		if err1 != nil {
			fmt.Println(err1)
			f.Close()
		}
	})
}
