package capdec

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetImgSrcAsBase64(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(bytes)
	base64Encoding := ""
	if mimeType == "image/jpeg" || mimeType == "image/png" {
		base64Encoding = fmt.Sprintf("data:%v;base64,", mimeType)
	} else {
		errMsg := "File mime type is suitable for image src.\n"
		errMsg += fmt.Sprintf("Filepath: %v, Mime Type: %v", filepath, mimeType)
		return "", errors.New(errMsg)
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding, nil
}
