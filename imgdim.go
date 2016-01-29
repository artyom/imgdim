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
	"log"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s filename.{jpeg,png,gif,bmp,tiff}", os.Args[0])
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d,%d\n", cfg.Width, cfg.Height)
}

func init() {
	log.SetFlags(0)
}
