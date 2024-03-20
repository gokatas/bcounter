// Bytecounter implements a (concrete) type whose method counts bytes before
// discarding them. Since it satisfies the io.Writer interfaces (an abstract type),
// we can pass it to fmt.Fprint.
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c := ByteCounter(len(p)) // must explicitly convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello")) // c == 5
	c = 0                    // reset the counter
	fmt.Fprint(&c, "world")  // c == 5
}
