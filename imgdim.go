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
	cfg, _, err := image.DecodeConfig(f)
	return cfg.Width, cfg.Height, err
}
