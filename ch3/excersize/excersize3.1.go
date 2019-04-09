// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	width, height = 600, 450            //height and width of the canvas
	cells         = 100                 //size of grid
	xyrange       = 100.0               // range of axes
	xyscale       = width / 2 / xyrange //The width of the unit coordinates on the canvas
	zscale        = 0.4 * height        //The height of the unit coordinates on the canvas
	scale         = 30                  //custom zoom factor
	angle         = math.Pi / 6         //30°
)

// corner returns x and y of canvas
func corner(i, j int) (float64, float64, error) {
	x := (float64(i)/cells - 0.5) * xyrange
	y := (float64(j)/cells - 0.5) * xyrange
	z, err := heightZ(x, y)

	sx := (x-y)*math.Cos(angle)*xyscale + width/2
	sy := (x+y)*math.Sin(angle)*xyscale - zscale*z + height/2
	return sx, sy, err
}

// corner returns x and y of canvas. ideas such as resove.md
func corner1(i, j int) (float64, float64, error) {
	x := float64(i) - cells/2
	y := float64(j) - cells/2
	z, err := heightZ(x, y)

	sx := (x-y)*math.Cos(angle) + cells/2
	sy := (x+y)*math.Sin(angle) - scale*z + cells/2

	dx := sx / cells * width
	dy := sy / cells * height
	return dx, dy, err
}

// heightZ return z of axes
func heightZ(x, y float64) (float64, error) {
	r := math.Hypot(x, y)
	z := math.Sin(r) / r
	if z >= math.MaxFloat64 {
		return z, errors.New("Data out of bounds")
	}

	return z, nil
}

func main() {
	svgStr := "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='" + strconv.Itoa(width) + "px' height='" + strconv.Itoa(height) + "px' >\n"
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner1(i+1, j)
			bx, by, err := corner1(i, j)
			cx, cy, err := corner1(i, j+1)
			dx, dy, err := corner1(i+1, j+1)
			if err != nil {
				continue
			}
			svgStr += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svgStr += "</svg>"
	err := ioutil.WriteFile("surface.svg", []byte(svgStr), 0777)
	fmt.Sprintln(err)
	// os.Create("surface.svg")
}
