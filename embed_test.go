package golang_embed

import (
	"fmt"
	"testing"
	_ "embed"
)
//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}