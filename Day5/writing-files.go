package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checking(e error)  {
	if e != nil {
		panic(e)
	}
}

func main()  {
	d1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	checking(err)

	f, err := os.Create("/tmp/dat2")
	checking(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checking(err)
	fmt.Printf("write %d byte\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrotes %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("write %d bytes\n", n4)

	w.Flush()
}
