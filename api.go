package capdec

import (
	"embed"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
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

func init(){
	browser = rod.New().MustConnect()
}

func ChangeMaxBrowserDimensions(width, height int){
	MaxBrowserWidth = width
	MaxBrowserHeight = height
}

func Caption(srcImgPath string, captions []string, destImgPath string, codes []string) error {
	imgBase64, err := GetImgSrcAsBase64(srcImgPath)
	if err != nil {
		return err
	}

	tempDir := os.TempDir()
	htmlFile, err := ioutil.TempFile(tempDir, "capdec_*.html")
	if err != nil {
		return err
	}
	//defer os.Remove(htmlFile.Name())

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

	//page := rod.New().MustConnect().MustPage("file://" + htmlFile.Name())
	page := browser.MustPage("file://" + htmlFile.Name())
	page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{Width: MaxBrowserWidth, Height: MaxBrowserHeight, Scale: 1})
	for _, code := range codes {
		page.Eval(code)
	}
	element := page.MustElement("#figure")
	element.MustScreenshot(destImgPath)
	page.Close()


	err = htmlFile.Close()
	if err != nil{
		fmt.Println("%v cannot be closed: Error: %v", htmlFile.Name(), err)
		return err
	}
	err = os.Remove(htmlFile.Name())
	if err != nil{
		fmt.Println("%v cannot be removed: Error: %v", htmlFile.Name(), err)
		return err
	}
	return nil
}
