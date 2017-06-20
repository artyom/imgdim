// Command imgdim returns image dimensions in format "width,height" on stdout.
// It's intended as a lightweight replacement for ImageMagick's `identify
// -format "%w,%h"` call which is too heavy on big images.
package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename.{jpeg,png,gif,bmp,tiff}\n", os.Args[0])
		os.Exit(1)
	}
	w, h, err := dimensions(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d,%d\n", w, h)
}

func dimensions(name string) (width, height int, err error) {
	f, err := os.Open(name)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	cfg, format, err := image.DecodeConfig(f)
	if format != "jpeg" {
		return cfg.Width, cfg.Height, err
	}
	if _, err := f.Seek(0, os.SEEK_SET); err != nil {
		return 0, 0, err
	}
	e, err := exif.Decode(f)
	if err != nil {
		return cfg.Width, cfg.Height, nil
	}
	o, err := e.Get(exif.Orientation)
	if err != nil || o == nil || len(o.Val) != 2 {
		return cfg.Width, cfg.Height, nil
	}
	for _, x := range o.Val {
		switch x {
		case 6, 8: // 90ºCCW, 90ºCW
			return cfg.Height, cfg.Width, nil
		}
	}
	return cfg.Width, cfg.Height, nil
}
