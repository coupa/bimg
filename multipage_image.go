package bimg

/*
#cgo pkg-config: vips
#include "vips/vips.h"
*/
import "C"

import (
	log "github.com/sirupsen/logrus"
)

func MultipageTIFFToPng(buf []byte, pages int) ([]byte, error) {
	defer C.vips_thread_shutdown()
	frames := make([]*C.VipsImage, pages)

	// For avoiding memory leak
	clear_frames := func() {
		for i := 0; i < pages; i++ {
			C.g_object_unref(C.gpointer(frames[i]))
		}
	}

	var err error
	for i := 0; i < pages; i++ {
		frames[i], err = vipsTIFFReadWithAlpha(buf, i)
		if err != nil {
			log.Errorf("[MultipageTIFFToPng] vipsTIFFReadWithAlpha err - %v", err)
			return nil, err
		}
	}

	outVipsImage, err := vipsArrayJoin(frames)
	if err != nil {
		clear_frames()
		log.Errorf("[MultipageTIFFToPng] vipsArrayJoin err - %v", err)
		return nil, err
	}

	clear_frames()
	defer C.g_object_unref(C.gpointer(outVipsImage))

	outBuf, err := getImageBufferV2(outVipsImage, PNG)
	if err != nil {
		log.Errorf("[MultipageTIFFToPng] getImageBuffer err - %v", err)
		return nil, err
	}
	return outBuf, nil
}
