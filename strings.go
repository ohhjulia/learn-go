package main

import (
	"fmt"
	"io"
	"strings"
)

func main9() {
	offset := int64(17)
	var reader1 strings.Reader
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset
	readingIndex, _ := reader1.Seek(offset, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek) \n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me) \n", expectedIndex)
	reader2 := strings.NewReader("hello go 爱好者!")
	fmt.Printf("reader2 Len is %d \n", reader2.Len())
	for i := 0; reader2.Len() > 0; i++ {
		ch, size, _ := reader2.ReadRune()
		fmt.Printf("%d is %q, size is %d\n", i, ch, size)
	}
}
