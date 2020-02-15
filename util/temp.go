package util

import (
	"context"
	"fmt"
	"github.com/aaronland/go-image-encode"
	"image"
	"io/ioutil"
	"os"
)

func AsTempFile(ctx context.Context, enc encode.Encoder, im image.Image) (string, error) {

	wr, err := ioutil.TempFile("", "image-")

	if err != nil {
		return "", err
	}

	defer wr.Close()

	err = enc.Encode(ctx, im, wr)

	if err != nil {
		return "", err
	}

	// see what's going on here - this (appending the format
	// extension) is necessary because without it fpdf.GetImageInfo
	// gets confused and FREAKS out triggering fatal errors
	// along the way... oh well (20171125/thisisaaronland)

	fname := wr.Name()
	wr.Close()

	ext := enc.Extension()

	fq_fname := fmt.Sprintf("%s%s", fname, ext)

	err = os.Rename(fname, fq_fname)

	if err != nil {
		return "", err
	}

	return fq_fname, nil
}
