package variables

import "fmt"

func Demo1() {
	var metin string = "hello world!"
	fmt.Print(metin)
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)

	var kdv2 float32 = 0.5
	fmt.Println(kdv2)

	kdv3 := 25.4 //
	fmt.Printf("veri tipi= %T\n", kdv3)

	var state bool

	var metin1, metin2 string = "ahmet", "veysel"

	state = metin1 == metin2
	fmt.Println(state)
}
