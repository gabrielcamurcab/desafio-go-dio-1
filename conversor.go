package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	var tempk = ebulicaoK
	var tempc = tempk - 273.0

	fmt.Printf("A temperatura de ebulição da água em K = %g, corresponde a %gºC", tempk, tempc)

}
