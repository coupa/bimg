package bimg

/*
#cgo pkg-config: vips
#include "vips/vips.h"
*/
import "C"

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	//"strings"
)

const PNGBufferError = "vips2png: unable to write to buffer"
const PNGTargetError = "vips2png: unable to write to target"

func ImageJoin(imgArr []*Image, imageType ImageType) ([]byte, error) {
	defer C.vips_thread_shutdown()

	image, err := vipsArrayJoin(imgArr)
	defer C.g_object_unref(C.gpointer(image))

	if err != nil {
		log.Errorf("[ImageJoin] vipsArrayJoin err - %v", err)
		return nil, err
	}

	return getImageBufferFromFile(image, imageType)
	// imageCopy := image
	// //imageCopy := C.vips_image_copy_memory(image)
	// buf, err := getImageBuffer(image, imageType)
	// if err != nil {
	// 	log.Errorf("Inside getImageBuffer error - %v", err)
	// 	log.Errorf("err.Error() %s", err.Error())

	// 		`getImageBuffer` internally calls `vips_pngsave_bridge` which throws "vips2png: unable to write to buffer" when we try to save a PNG that has more (100) pages.
	// 		So when we get vips2png error, we write the vips image to a temp file and read the data from temp file. Directly writing to a temp file without calling `getImageBuffer` has
	// 		performance impacts and hence writing to temp file when PNGError is thrown from `getImageBuffer`.

	// 	if strings.Contains(err.Error(), PNGBufferError) || strings.Contains(err.Error(), PNGTargetError) {
	// 		return getImageBufferFromFile(imageCopy, imageType)
	// 	}
	// 	return nil, err
	// }

	// return buf, nil
}

// Write C.VipsImage to temp file and read the buffer from temp file.
// This is called only when `getImageBuffer` throws `vips2png: unable to write to buffer` error.
func getImageBufferFromFile(image *C.VipsImage, imageType ImageType) ([]byte, error) {
	fileExt := getExtension(imageType)
	if fileExt == "unknown" {
		return nil, errors.New(fmt.Sprintf("[ImageJoin] Unknown extension for the imagetype %s", ImageTypeName(imageType)))
	}

	tmpFile, err := ioutil.TempFile("", "joined-img-*"+fileExt)
	if err != nil {
		log.Errorf("[ImageJoin] Error in initializing tmp file - %v", err)
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	err = vipsWriteToFile(image, tmpFile.Name())
	if err != nil {
		log.Errorf("[ImageJoin] vipsWriteToFile error - %v", err)
		return nil, err
	}

	return Read(tmpFile.Name())
}

func getExtension(imageType ImageType) string {
	if imageType == PNG {
		return ".png"
	}

	// Right now, `ImageJoin` function joins only PNG images.
	return "unknown"
}
