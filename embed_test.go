package golang_embed

import (
	"fmt"
	"testing"
	_ "embed"
	"io/fs"
	"io/ioutil"
)
//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed image.png
var image []byte

func TestEmbedFile(t *testing.T) {
	err := ioutil.WriteFile("image_new.png", image, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}