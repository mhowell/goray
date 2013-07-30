package main

import (
	"github.com/mhowell/goray"
	"image/png"
	"log"
	"os"
)

func main() {

	renderer := new(goray.Renderer)

	renderer.CreateScene()
	renderer.CreateCamera()

	img := renderer.RenderScene()

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

}
