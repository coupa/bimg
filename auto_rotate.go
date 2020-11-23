package bimg

/*
#cgo pkg-config: vips
#include "vips/vips.h"
*/
import "C"

import (
	"runtime"
	log "github.com/sirupsen/logrus"

)

// AutoRotate is used to rotate a given image based on the image's orientation found via exiftool - // https://linux.die.net/man/1/exiftool
// Note: This method does not rotate if the orientation is 1 (i.e. horizontal)
//
// Refer the following documents to know more about orientation and exiftool:
// https://sirv.com/help/articles/rotate-photos-to-be-upright/
// http://sylvana.net/jpegcrop/exif_orientation.html
func AutoRotate(buf []byte, o Options) ([]byte, error) {
	// Required in order to prevent premature garbage collection. See:
	// https://github.com/h2non/bimg/pull/162
	defer runtime.KeepAlive(buf)
	defer C.vips_thread_shutdown()

	// Override the NoAutoRotate option to set to false so that it autorotates.
	o.NoAutoRotate = false

	image, imageType, err := loadImageWithOptions(buf, o)
	if err != nil {
		log.Errorf("[AutoRotate] loadImageWithOptions err - %v", err)
		return nil, err
	}

	outVipsImage, _, err := rotateAndFlipImage(image, o)
	if err != nil {
		log.Errorf("[AutoRotate] rotateAndFlipImage err - %v", err)
		return nil, err
	}

	defer C.g_object_unref(C.gpointer(outVipsImage))

	outBuf, err := getImageBufferViaSlice(outVipsImage, imageType)
	if err != nil {
		log.Errorf("[AutoRotate] getImageBufferViaSlice err - %v", err)
		return nil, err
	}
	return outBuf, nil
}
