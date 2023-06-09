package golang_embed

import (
	"fmt"
	"testing"
	"embed"
	_ "embed"
	"io/fs"
	"io/ioutil"
)
// embed string
//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

// embed byte
//go:embed image.png
var image []byte

func TestEmbedFile(t *testing.T) {
	err := ioutil.WriteFile("image_new.png", image, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}

// embed multiple file
//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
}