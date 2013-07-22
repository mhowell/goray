package goray



func createScene() (Scene) {
	
	s1 := new(Sphere)
	s1.center = Point{-4, 0, 10 }
	s1.radius = 2

	s2 := new(Sphere)
	s2.center = Point{ 4, 0, 10}
	s2.radius = 3

	spheres := make([]Shapes, 2)

	spheres[0] = s1
	spheres[1] = s2

	l1 := Light{Point{0.0, 5, 5}}

	lights := make([]Light, 1)
	lights[0] = l1

	return Scene{ lights, spheres}

}

func createCamera() (Camera) {
	return Camera{Point{0,0,-10}, 100, 100}
}

func main() {

	scene := createScene()
	camera := createCamera()

	renderer := Renderer{&scene, &camera}

	renderer.renderScene()
}
