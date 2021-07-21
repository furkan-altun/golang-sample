package main

import "fmt"

func main() {

	s := make([]string, 3, 77)
	fmt.Println("emp:", s)
	fmt.Println("length:", len(s))
	fmt.Println("capasity:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	f := []string{"f","u","r","k"}
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	s = append(s, f...)
	// region deleting slice element
	fmt.Println("after f:",s) //  [a b c d e f f u r k] -- deleting f ( 5 and 6th index)
	fmt.Println("temp 1", s[:5]) // [a b c d e] 5.th index not included
	fmt.Println("temp 2", s[7:]) // [u r k] 7.tl index included
	s = append(s[:5], s[7:]...)
	// endregion deleting slice element
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}