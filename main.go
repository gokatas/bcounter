// Bcounter implements a (concrete) type whose method counts bytes before
// discarding them. Since it satisfies the io.Writer interfaces (an abstract type),
// we can pass it to fmt.Fprint.
package main

import "fmt"

type bCounter int

func (b *bCounter) Write(p []byte) (int, error) {
	*b += bCounter(len(p)) // must explicitly convert int to bCounter
	return len(p), nil
}

func main() {
	var b bCounter
	b.Write([]byte("hello")) // b == 5
	b = 0                    // reset the counter
	fmt.Fprint(&b, "world")  // b == 5
}
