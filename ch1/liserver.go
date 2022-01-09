package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles := int(5)
		if r.URL.Query().Has("cycles") {
			c, err := strconv.Atoi(r.URL.Query().Get("cycles"))
			if err == nil {
				cycles = c
			}
		}
		lissajous(w, float64(cycles))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{
	color.White,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 200
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*math.Pi*cycles; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := (uint8)(rand.Intn(3) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
