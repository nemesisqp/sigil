package main

import (
	"crypto/md5"
	"image/color"
	"image/png"
	"io/ioutil"
	"bytes"
	"github.com/cupcake/sigil/gen"
	"flag"
	"os"
)

var config = gen.Sigil{
	Rows: 5,
	Foreground: []color.NRGBA{
		rgb(45, 79, 255),
		rgb(254, 180, 44),
		rgb(226, 121, 234),
		rgb(30, 179, 253),
		rgb(232, 77, 65),
		rgb(49, 203, 115),
		rgb(141, 69, 170),
	},
	Background: rgb(224, 224, 224),
}

func rgb(r, g, b uint8) color.NRGBA { return color.NRGBA{r, g, b, 255} }



func md5hash(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func main() {
	var imgSize = flag.Int("size", 240, "image size")
	var input = flag.String("data", "", "image input seed")
	var data = md5hash(*input)
	var buf bytes.Buffer
	png.Encode(&buf, config.Make(*imgSize, false, data))
	
	
	err := ioutil.WriteFile("identicon.png", buf.Bytes(), os.ModePerm)
	if err != nil { panic(err) }
}
