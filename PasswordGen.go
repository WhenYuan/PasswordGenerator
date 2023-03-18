package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	A := [89]string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "D", "F", "G", "H", "J", "K", "L", "Z", "X", "C", "V", "B", "N", "M", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "`", "-", "=", "[]", "]", ";", "'", ",", ".", "/", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "+", "{", "}", ":", "<", ">", "?", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}

	var b int

	rand.NewSource(time.Now().UnixNano())

	fmt.Print("Enter the lenth of the password that you want to generate:")
	fmt.Scanln(&b)
	if b == 0 {
		main()
	}
	var B int = b + 1
	var d string
	for i := 0; i != B; i++ {
		var c int64 = rand.Int63n(88) + 1
		d += A[c]
		if i == b-1 {
			var Z string
			var I int
			for i := 1; i != b; i++ {
				Z = string(d[i])
				for i := 1; i != b; i++ {
					if Z == string(d[i]) {
						I++
						if I > len(d)/3 {
							main()
							break
						}
					}
				}
			}
			fmt.Println(d)
			break
		}
	}

	main()
}
