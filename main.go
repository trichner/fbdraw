package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/trichner/fbmon/framebuffer"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/gofont/gomono"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"os"
	"time"
)

const LOREM_IPSUM = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

func main() {

	dev := os.Args[1]

	fb, err := framebuffer.Open(dev)
	if err != nil {
		log.Panic(err)
	}

	//font := loadFont()
	font := basicfont.Face7x13

	tick := time.NewTicker(time.Millisecond * 1000)

	ctx := gg.NewContextForImage(fb)
	ctx.SetRGB(0, 0, 0)
	ctx.Clear()
	draw.Draw(fb, fb.Bounds(), ctx.Image(), image.ZP, draw.Src)

	// white background
	rect := fb.Bounds()

	center := image.Pt(rect.Dx()/2, rect.Dy()/2)

	var r byte = 0
	for {

		{
			ctx.Push()
			ctx.SetColor(color.RGBA{r, 0, 0, 0xff})
			ctx.DrawCircle(float64(center.X), float64(center.Y), float64(math.Min(float64(rect.Dy()), float64(rect.Dx())/3)))
			ctx.Fill()
			ctx.Pop()
		}

		{
			ctx.Push()
			ctx.SetFontFace(font)
			ctx.SetRGB(0, 1, 0)

			lines := ctx.WordWrap(LOREM_IPSUM, float64(rect.Dx()-2*20))

			y := rect.Min.Y + 20
			for i, l := range lines {
				ctx.DrawString(fmt.Sprintf("%2d: %s", i, l), 5, float64(y))
				ctx.Fill()
				y += 20
			}
			ctx.Pop()
		}

		draw.Draw(fb, fb.Bounds(), ctx.Image(), image.ZP, draw.Src)
		<-tick.C
		r += 0xff / 8
	}
}

func loadFont() font.Face {

	font, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}

	return truetype.NewFace(font, &truetype.Options{Size: 48})
}
