// Mendelbrot create a png for mandelbrot set.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	maxReal, minReal = 2, -2      //the range of real number
	maxImag, minImag = 2, -2      //the rang of imaginary
	width, height    = 1024, 1024 //the size of canvas
)

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for cy := 0; cy < height; cy++ {
		y := 2 - 4*float64(cy)/height
		for cx := 0; cx < width; cx++ {
			x := 4*float64(cx)/width - 2
			z := complex(x, y)
			img.Set(cx, cy, mendelbrot(z))
		}
	}

	// this part is from book.
	// for cy := 0; cy < height; cy++ {
	// 	y := float64(cy)/height*(maxReal-minReal) + minReal
	// 	for cx := 0; cx < width; cx++ {
	// 		x := float64(cx)/width*(maxImag-minImag) + minImag
	// 		z := complex(x, y)
	// 		img.Set(cx, cy, mendelbrot(z))
	// 	}
	// }

	// creat a file of png
	pngFile, err := os.Create("mendelbrot.png")
	err = png.Encode(pngFile, img)
	if err != nil {
		fmt.Println(err)
	}
}

// mandelbrot returns color of z point
func mendelbrot(z complex128) color.Color {
	var v complex128
	const iterations, constrast = 200, 15
	for i := uint8(0); i < iterations; i++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*i}
		}
	}
	return color.Black
}
