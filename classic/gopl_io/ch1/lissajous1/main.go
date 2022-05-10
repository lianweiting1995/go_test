package lissajous1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.White,
	color.Black,
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的 x 轴转数
		res     = 0.001 // 旋转的角度
		size    = 100   // 图片封面增加/减少的尺寸
		nframes = 64    // 动画帧数
		delay   = 8     // 帧之间的间隔
	)
	freq := rand.Float64() * 3.0 // y轴旋转频率
	anim := gif.GIF{
		LoopCount: nframes,
	}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
