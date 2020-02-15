package util

import (
	"context"
	"fmt"
	"github.com/aaronland/go-image-encode"
	"image"
)

// this works but not in templates for some reason - the b64 data is
// always truncated... (20181024/straup)

func AsDataURL(ctx context.Context, enc encode.Encoder, im image.Image) string {

	b64, err := AsBase64String(ctx, enc, im)

	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	format := enc.MimeType()

	return fmt.Sprintf("data:image/%s;base64,%s", format, b64)
}
