package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"time"
	"net/http"
	"strconv"
	"log"
)

var palette = [] color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		lissajous(w, r)
	})

	http.ListenAndServe("localhost:8000", nil)
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	delay, err := strconv.Atoi(r.FormValue("delay"))
	if err != nil {
		log.Print(err)
	}

	const (
		cycles 	= 5
		res 	= 0.001
		size 	= 100
		nframes = 128
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i< nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+2)
		img := image.NewPaletted(rect, palette)
		for t := 0.0;  t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}


