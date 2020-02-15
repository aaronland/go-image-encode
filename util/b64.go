package util

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aaronland/go-image-encode"
	"image"
)

func AsBase64String(ctx context.Context, enc encode.Encoder, im image.Image) (string, error) {

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err := enc.Encode(ctx, im, wr)

	if err != nil {
		return "", err
	}

	wr.Flush()

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
