package image

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func ToWebP(in string, out string) error {
	buffer, err := bimg.Read(in)
	if err != nil {
		fmt.Println(err)
		return err
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	bimg.Write(out, newImage)
	return nil
}
