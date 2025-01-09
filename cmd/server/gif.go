package server

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"time"

	"github.com/ericpauley/go-quantize/quantize"
	"github.com/go-rod/rod"
	"golang.org/x/image/draw"
)

func OpenHtmlPage(html string) *rod.Page {
	uri := fmt.Sprint("data:text/html;utf8,", html)
	page := rod.New().MustConnect().MustPage(uri).MustWaitLoad().MustWaitDOMStable()
	return page
}

func OpenUrlPage(url string) *rod.Page {
	fmt.Println(url)
	page := rod.New().MustConnect().MustPage(url).MustWaitLoad().MustWaitDOMStable()
	return page
}

var (
	gifWidth       int = 400
	gifHeigth      int = 400
	framePerSecond int = 2
	pixelPerSecond int = 200
)

func PageGif(page *rod.Page) *gif.GIF {
	pageFirstImgBytes, _ := page.Screenshot(true, nil)
	pageFirstImg, _ := png.Decode(bytes.NewReader(pageFirstImgBytes))
	pageWidth, pageHeight := pageFirstImg.Bounds().Dx(), pageFirstImg.Bounds().Dy()
	pageStep := (pixelPerSecond * pageWidth) / (framePerSecond * gifWidth)
	fmt.Println(pageWidth, pageHeight, (pageHeight-pageWidth)/pageStep)
	pageImageBytes := make([][]byte, ((pageHeight - pageWidth) / pageStep))
	for i := range pageImageBytes {
		fmt.Println(i)
		pageImageBytes[i], _ = page.Screenshot(true, nil)
		time.Sleep(time.Second / time.Duration(framePerSecond))
	}

	pageImages := make([]image.Image, len(pageImageBytes))
	for i, ib := range pageImageBytes {
		fmt.Println(i)
		pageImages[i], _ = png.Decode(bytes.NewReader(ib))
	}

	gifImages := make([]*image.Paletted, len(pageImages))
	gifDelays := make([]int, len(pageImages))
	gifRectangle := image.Rect(0, 0, gifWidth, gifHeigth)
	for i := range pageImages {
		fmt.Println(i)
		quantizer := quantize.MedianCutQuantizer{}
		palette := quantizer.Quantize(make([]color.Color, 0, 256), pageImages[i])

		gifImages[i] = image.NewPaletted(gifRectangle, palette)
		pageRectangle := image.Rect(0, i*pageStep, pageWidth, i*pageStep+pageWidth)
		draw.ApproxBiLinear.Scale(gifImages[i], gifRectangle, pageImages[i], pageRectangle, draw.Over, nil)
		gifDelays[i] = 100 / framePerSecond
	}

	return &gif.GIF{Image: gifImages, Delay: gifDelays}
}
