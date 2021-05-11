# CapDec: Caption Decorator

CapDec provides a simple API and CLI for captioning pictures.

## Installation

1. `go get github.com/gusanmaz/capdec`

2. `cd ${GOPATH}/src/github.com/gusanmaz/capdec`

3. `go install cmd/capdec/capdec.go`

## API Usage

CapDec module currently offers only one function: `caption`
This function's simple usage could be checked from `main/main.go` or below

```go
dest := "bird_captioned.png"
captions := make([]string, 2)
captions[0] = "Blue-footed bird."
captions[1] = "What a fancy bird!"
codes := make([]string, 2)
codes[0] =  `document.querySelector('figure').style.borderWidth = "10px"`
codes[1] =  `document.querySelector('figure').style.borderColor = "blue" ;`

err := capdec.Caption("bird.jpeg", captions, dest, codes)
if err != nil{
    log.Fatal("Terminating because of an error. Error: " + err.Error())
}else{
    fmt.Printf("Captioned image successfully saved at %v\n",dest)
}
```

`Caption` function allows customizing default style of caption and picture border with utilization of `codes` array.

Every element of `codes` array should be valid Javascript statement that could change default styling of the template file of `tempplate/template.html`

### Content of template/template.html

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>CapDec Page</title>
        <style>
            figcaption{
                font-size: 20px;
                background-color: lightgoldenrodyellow;
                padding:10px;
                color:black;
            }
            figcaption:nth-child(2){
                color:mediumblue;
            }
            figure{
                width:100%;
                height: 100%;
                max-width: min-content;
                max-height: min-content;
                padding: 0;
                margin:0;
                border-color:red;
                border-width: 0px;
                border-style: solid;
            }
        </style>
    </head>
    <body>
    <figure id="figure">
        <img id="img" src="{{.ImgBase64}}" alt="Alt Text">
        {{range .Captions}}
            <figcaption>{{.}}</figcaption>
        {{end}}
    </figure>
    </body>
</html>
```

## CLI Usage

`capdec -s in.jpeg -d out.png "cap:quick fox jumps" "cap:over the lazy dog" "js:document.querySelector('figcaption').style.color = 'red'"`

* Arguments for caption text and codes should be typed after flags. 
* Prefix caption arguments with `cap`
* Prefix code argument with `js`

## Author

Güvenç Usanmaz

## License

MIT License