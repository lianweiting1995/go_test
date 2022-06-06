package ch3

import (
	"image"
	"image/color"
	"image/jpeg"
	"math/cmplx"
	"net/http"
)

func MandelbrotMain() *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -4, -4, +4, +4
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img

	// outFile, err := os.Create("../storage/MandelbrotMain.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outFile.Close()
	// b := bufio.NewWriter(outFile)
	// err = png.Encode(b, img)
	// if err != nil {
	// 	panic(err)
	// }
	// err = b.Flush()
	// if err != nil {
	// 	panic(err)
	// }
}

func HttpServer() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(rep http.ResponseWriter, req *http.Request) {
	rep.Header().Add("content-type", "image/png")
	img := MandelbrotMain()
	jpeg.Encode(rep, img, nil)
	// fmt.Fprintf(rep, "hello world")
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.RGBA{
		uint8(100),
		uint8(29),
		uint8(100),
		uint8(255),
	}
}
