package main

import (
	"image/png"
	"log"
	"os"
	"github.com/mhowell/goray"
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
