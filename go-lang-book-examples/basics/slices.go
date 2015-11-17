package main

import "fmt"

func main() {
	// 	arr := [5]int{10, 20, 30, 40, 50}
	// 	x := arr[3:5]

	// slice1 := []int{1, 2, 3}
	// slice2 := make([]int, 2)
	// copy(slice2, slice1)
	// fmt.Println(slice1, slice2)

	// -> map[string]int -> um Mapa com a chave sendo string e o value int -> { 'a' => 1 } em Ruby.
	// var x map[string]int 
	// x["number"] = 1
	// fmt.Println(x) // -> Gera erro, pois map precisa ser inicializado antes de usar...

	/////////////////////

	// x := make(map[string]int)
	// x["key"] = 10
	// x["other"] = 20
	// fmt.Println(x["key"], x["other"], len(x))

	// delete(x, "key")
	// fmt.Println(x["key"], x["other"], len(x))

	/////////////////////

	// elements := make(map[string]string)
	// elements["H"] = "Hidrogênio"
	// elements["He"] = "Hélio"
	// elements["Li"] = "Lítio"
	// elements["C"] = "Carbono"

	// name, ok := elements["xC"] 
	// fmt.Println(name, ok) // => " false" elements["xC"] == ""

	people := map[string]map[string]string{
		"Luiz": map[string]string{
			"full_name":"Luiz Cezer",
			"age":"25",
		},
		"Cezer": map[string]string{
			"full_name":"Cezer Luiz",
			"age": "20",
		},
		"Cezinha":map[string]string{
			"full_name":"Cezinha Cezer",
			"age":"18",
		},
	}

	fmt.Println(people)

	if person, ok := people["Luiz"]; ok {
		fmt.Println(person["full_name"], person["age"])
	}

}