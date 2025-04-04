package main

import (
	"math"
	"math/cmplx"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0.5, 1, 1)
	dc.Fill()
	dc.SavePNG("out.png")
	mandle_brot()
	// square()
	// test_dots()
	// println(base_fractal_func(0, 2, complex(+9.800000e-001, +0.800000e-001)))
}

func mandle_brot() {
	x_offset := -1.2
	max_lim := 1000.0
	lower_lim := -2 * max_lim
	upper_lim := 2 * max_lim
	granularity := 10.0
	scaling := granularity / 4 // 100 => 25.0
	color := 0.0
	points := 0

	dc := gg.NewContext(int(math.Abs(2*max_lim)), int(max_lim))
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	for x := lower_lim; x < upper_lim; x += granularity {
		for y := lower_lim; y < upper_lim; y += granularity {
			c := complex(x/max_lim+x_offset, y/max_lim)
			// println("before x,y=", x, y)
			draw_x, draw_y := transform_coords(x, y, max_lim, granularity)
			color = base_fractal_func(0, 2, c)
			dc.SetRGB(color, color, color)
			// println("after x,y=", draw_x, draw_y)
			dc.DrawCircle(draw_x, draw_y, scaling)
			dc.Fill()

			points++
			// println("color=", color)
			// println("points=", points)
		}
	}
	dc.Fill()
	dc.SavePNG("mandlebrot.png")
}

func base_fractal_func(z complex128, exp complex128, c complex128) float64 {
	// println("init")
	// println("z=", z)
	// println("exp=", exp)
	// println("c=", c)
	max_num := 1000
	for i := range max_num {
		z = cmplx.Pow(z, exp) + c
		// println("z=", z)
		if real(z) > math.MaxFloat64 || imag(z) > math.MaxFloat64 {
			// println("broke here z=", z)
			// println("i=", i)
			return float64(max_num-i+1) / float64(max_num)
		}
	}
	// println("out of max")
	return 0
}

func transform_coords(x float64, y float64, max_lim float64, granularity float64) (float64, float64) {
	// 100 => 25
	//
	offset := max_lim * 0.5
	// return 0.5*x + offset + scaling*2.5, -0.5*y + offset - scaling*2.5
	// new_scaling := 25 / scaling

	return 0.5*x + offset + granularity/4, -0.5*y + offset - granularity/4
	// 5  => 5
	// 10 => 2.5
	// 20 => 1.25
	// 40 => 0.625
	// new_scaling = 5*2**(-(scaling - 5)/5)
}

func test_dots() {
	max_lim := 1000.0
	lower_lim := -max_lim
	upper_lim := max_lim
	granularity := 100.0
	scaling := 25.0
	points := 0

	dc := gg.NewContext(int(max_lim), int(max_lim))
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	for x := lower_lim; x < upper_lim; x += granularity {
		for y := lower_lim; y < upper_lim; y += granularity {
			dc.SetRGB(0, 0, 0)
			println("before x,y=", x, y)
			draw_x, draw_y := transform_coords(x, y, max_lim, scaling)
			println("after x,y=", draw_x, draw_y)
			dc.DrawCircle(draw_x, draw_y, scaling)

			points++
			println("points=", points)
		}
	}
	dc.Fill()
	dc.SavePNG("mandlebrot.png")
}
