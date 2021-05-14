package capdec

import (
	"embed"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"text/template"
)

var (
	MaxBrowserWidth  = 10000
	MaxBrowserHeight = 10000
)

//go:embed template/template.html
var embedFS embed.FS
var browser *rod.Browser

func init() {
	browser = rod.New().MustConnect()
}

func GetImageDimensions(path string) (int, int, error) {
	reader, err := os.Open(path)
	if err != nil {
		return -1, -1, err
	}
	defer reader.Close()

	im, _, err := image.DecodeConfig(reader)
	if err != nil {
		return -1, -1, err
	}
	return im.Width, im.Height, nil
}

func Caption(srcImgPath string, captions []string, destImgPath string, codes []string) error {
	imgBase64, err := GetImgSrcAsBase64(srcImgPath)
	if err != nil {
		return err
	}

	w, h, err := GetImageDimensions(srcImgPath)
	if err != nil {
		return err
	}


	MaxBrowserWidth = int(float64(w) * 1.2)
	// Assuming captions are not so long.
	// TODO User could overwrite default browser dimensions assumptions
	MaxBrowserHeight = h + ((1000 * 1000) / MaxBrowserWidth)

	tempDir := os.TempDir()
	htmlFile, err := ioutil.TempFile(tempDir, "capdec_*.html")
	if err != nil {
		return err
	}

	templ, err := template.ParseFS(embedFS, "template/template.html")
	if err != nil {
		return err
	}

	type Params struct {
		ImgBase64 string
		Captions  []string
	}
	params := Params{ImgBase64: imgBase64, Captions: captions}

	err = templ.Execute(htmlFile, params)
	if err != nil {
		return err
	}

	page := browser.MustPage("file://" + htmlFile.Name())
	page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{Width: MaxBrowserWidth, Height: MaxBrowserHeight, Scale: 1})
	for _, code := range codes {
		page.Eval(code)
	}
	element := page.MustElement("#figure")
	element.MustScreenshot(destImgPath)
	page.Close()

	err = htmlFile.Close()
	if err != nil {
		fmt.Printf("%v cannot be closed: Error: %v", htmlFile.Name(), err)
		return err
	}
	err = os.Remove(htmlFile.Name())
	if err != nil {
		fmt.Printf("%v cannot be removed: Error: %v", htmlFile.Name(), err)
		return err
	}
	return nil
}
