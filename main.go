// Bcounter implements a (concrete) type whose method counts bytes before
// discarding them. Since it satisfies the io.Writer interface (an abstract
// type), we can pass it to fmt.Fprint. Adapted from
// https://github.com/adonovan/gopl.io/tree/master/ch7/bytecounter.
package main

import "fmt"

type bcounter int

func (b *bcounter) Write(p []byte) (int, error) {
	*b += bcounter(len(p)) // must explicitly convert int to bcounter
	return len(p), nil
}

func main() {
	var b bcounter           // declare variable of type bcounter
	b.Write([]byte("hello")) // b == 5
	b = 0                    // reset the byte counter
	data := "形"
	fmt.Fprint(&b, data)
	fmt.Printf("%s is %d bytes: %v\n", data, b, []byte(data))
}
