package bimg

/*
#cgo pkg-config: vips
#include "vips/vips.h"
*/
import "C"

func ImageJoin(imgArr []*Image, imageType ImageType) ([]byte, error) {
	defer C.vips_thread_shutdown()

	image, err := vipsArrayJoin(imgArr)
	if err != nil {
		return nil, err
	}

	buf, err := getImageBuffer(image, imageType)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
