package main

import "fmt"

func main() {
	i1 := map[string] int{
		"Furkan" : 27,
		"Ricardo" : 34,
		"Pepe" : 35,
	}
	fmt.Println(i1)

	fmt.Println(i1["Furkan"])

	fmt.Println(i1["Ali"])

	// i1 map inde olmayan bir atama yapıldığında value değeri tipinin zero value değeri olarak atanır.
	// örneğin i1 mapinde Ali diye bir key yok. Bu sebeple bunun value si map in response değeri olan int in zero valuesi 0 dır.
	// i1 map inde olmayan bir key değeriyle declaration yapıldığında hem value si hem de böyle bir değer olup olmadığını anlayabiliriz.
	// ilk partta (newValue) o key e ait varsa değeri yoksa tipinin zero value si tutulur.
	// ikinci partta (state) o key değeri var sa true yoksa false değeri döner.
	newValue, state := i1["Ali"]
	fmt.Println(newValue)
	fmt.Println(state)

	// map içerisinde olup olmadığının kontrolü için aşağıdaki şekilde kod bloğu kullanılabilir.
	if val, flag := i1["Rosier"]; flag == true {
		fmt.Println(val, " değeri mevcut")
	}else {
		fmt.Println(val, " değeri mevcut değil")
	}

	// map e ekleme yapma
	i1["Altun"] = 99
	i1["Bjk"] = 1903
	for key, val := range i1{
		fmt.Println(key, val)
	}

	i2 := []int{5,3,6,7,8,22,11}

	for _, val := range i2{
		fmt.Println(val)
	}

	// delete from map
	i3 := map[string] int{
		"a1" : 1,
		"a2" : 2,
	}
	fmt.Println(i3)
	delete(i3, "a2")
	fmt.Println(i3)
	delete(i3, "a3")
	fmt.Println(i3)

	if val, ok := i3["a3"]; !ok{
		fmt.Println(val, "not exist")
	}
	delete(i3, "a1")
	fmt.Println(i3)

}
