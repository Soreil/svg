package svg

import (
	"image/png"
	"os"
	"testing"
)

const dataDir = "testData/"
const input = dataDir + "The United States of America.svg"
const output = input + ".png"

func TestSVG(t *testing.T) {
	if _, err := os.Stat(input); err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	img, err := Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	out, err := os.Create(output)
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(out, img); err != nil {
		t.Fatal(err)
	}
}
