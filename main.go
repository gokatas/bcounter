// Bcounter implements a (concrete) type whose method counts bytes before
// discarding them. Since it satisfies the io.Writer interface (an abstract type),
// we can pass it to fmt.Fprint.
package main

import "fmt"

type bcounter int

func (b *bcounter) Write(p []byte) (int, error) {
	n := len(p)
	*b += bcounter(n) // must explicitly convert int to bcounter
	return n, nil
}

func main() {
	var b bcounter // declare variable of type bcounter

	b.Write([]byte("hello"))
	fmt.Printf("%v is %v bytes: %v\n", "world", b, []byte("world"))

	b = 0 // reset the byte counter

	katas := "形"
	fmt.Fprint(&b, katas)
	fmt.Printf("%v is %v bytes: %v\n", katas, b, []byte(katas))
}
