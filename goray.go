package goray



func createScene() (Scene) {

}

func createCamera() (Camera) {

}

func main() {

	scene := createScene()
	camera := createCamera()

	renderer := Renderer{&scene, &camera}

	renderer.renderScene()
}
