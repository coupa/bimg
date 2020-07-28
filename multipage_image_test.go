package bimg

import (
	"fmt"
	"os"
	"testing"
)

func TestMultipageTIFFToPng(t *testing.T) {
	t.Run("For converting multi page tiff to png", func(t *testing.T) {
		buf, _ := imageBuf("cust_unable_target_issue.tif")
		outBuf, err := MultipageTIFFToPng(buf, 61)
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
