package main

import (
	"bytes"
	"fmt"
)

func main() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", contents, buffer1.Cap())
	unreadBytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	buffer1.WriteString("cdefghi")
	fmt.Printf("The capacity of the buffer: %d\n", buffer1.Cap())
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	unreadBytes[len(unreadBytes)-2] = byte('X')
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.String())
}
