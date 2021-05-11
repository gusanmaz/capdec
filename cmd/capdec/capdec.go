package main

import (
	"flag"
	"fmt"
	"github.com/gusanmaz/capdec"
	"os"
	"strings"
)

var (
	srcFlag  string
	destFlag string
	codeFlag string
)

const (
	srcDefault  = "in.png"
	destDefault = "out.png"
	//codeDefault = ""
	srcUsage  = "Filepath for source image"
	destUsage = "Filepath for destination image"
	//codeUsage = "JS code to change default CSS values of the HTML template."
)

func main() {
	flag.StringVar(&srcFlag, "src", srcDefault, srcUsage)
	flag.StringVar(&srcFlag, "s", srcDefault, srcUsage)
	flag.StringVar(&destFlag, "dest", destDefault, destUsage)
	flag.StringVar(&destFlag, "d", destDefault, destUsage)
	//flag.StringVar(&codeFlag, "code", codeDefault, codeUsage)
	//flag.StringVar(&codeFlag, "c", codeDefault, codeUsage)
	flag.Parse()

	captions := make([]string, 0)
	codes := make([]string, 0)
	for _, arg := range flag.Args() {
		if strings.HasPrefix(arg, "js:") {
			codes = append(codes, strings.TrimPrefix(arg, "js:"))
		}
		if strings.HasPrefix(arg, "cap:") {
			captions = append(captions, strings.TrimPrefix(arg, "cap:"))
		}
	}

	fmt.Printf("Caption generation for %v is in progres...\n", srcFlag)
	err := capdec.Caption(srcFlag, captions, destFlag, codes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error has occured!\n")
		fmt.Fprintf(os.Stderr, "Error: "+err.Error()+"\n")
		fmt.Fprintf(os.Stderr, "Terminating program...\n")
	} else {
		fmt.Printf("Captioned image successfully saved at %v\n", destFlag)
	}
}
