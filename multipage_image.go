package bimg

/*
#cgo pkg-config: vips
#include "vips/vips.h"
*/
import "C"

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type VIPSImage struct {
	image *C.VipsImage
	err   error
}

func MultipageTIFFToPng(buf []byte, pages int) ([]byte, error) {
	defer C.vips_thread_shutdown()
	var wg sync.WaitGroup

	tiffImages := make([]VIPSImage, pages)
	// Load each page in tiff file via go-routines
	for i := 0; i < pages; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			out, err := vipsTIFFReadWithAlpha(buf, i)
			if err != nil {
				tiffImages[i] = VIPSImage{image: out, err: err}
			}
			tiffImages[i] = VIPSImage{image: out, err: nil}
		}(i, &wg)
	}

	// Wait for all the go-routines to finish
	wg.Wait()

	frames := make([]*C.VipsImage, pages)
	defer func(frames []*C.VipsImage) {
		for i := 0; i < len(frames); i++ {
			if frames[i] != nil {
				C.g_object_unref(C.gpointer(frames[i]))
			}
		}
	}(frames)

	for i, tiffImage := range tiffImages {
		// Check if there are any errors in the individual tiff page load and return if there are any
		if tiffImage.err != nil {
			log.Errorf("[vipsTIFFReadWithAlpha] Page %d -> Error is - %v", i, tiffImage.err)
			return nil, tiffImage.err
		}
		frames[i] = tiffImage.image
	}

	outVipsImage, err := vipsArrayJoin(frames)
	if err != nil {
		log.Errorf("[MultipageTIFFToPng] vipsArrayJoin err - %v", err)
		return nil, err
	}

	defer C.g_object_unref(C.gpointer(outVipsImage))

	outBuf, err := getImageBufferViaSlice(outVipsImage, PNG)
	if err != nil {
		log.Errorf("[MultipageTIFFToPng] getImageBuffer err - %v", err)
		return nil, err
	}
	return outBuf, nil
}
