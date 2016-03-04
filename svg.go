//SVG go image package driver.
package svg

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg onload="loaded()" xmlns="http://www.w3.org/2000/svg`

func init() {
	image.RegisterFormat("svg", svgHeader, Decode, DecodeConfig)
}

//Encode SVG to PNG as image.Image
func decode(input []byte) (image.Image, error) {
	cmd := exec.Command("rsvg-convert")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stdin = strings.NewReader(string(input))

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	img, err := png.Decode(bytes.NewReader(out.Bytes()))
	if err != nil {
		return nil, err
	}
	return img, nil
}

//Decodes the first frame of an SVG file into an image
func Decode(r io.Reader) (image.Image, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return decode(b)
}

//Returns metadata
func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, errors.New("Not implemented")
}
