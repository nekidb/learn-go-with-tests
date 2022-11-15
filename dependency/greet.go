package main

import (
		"fmt"
		"io"
		"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "\d", 3)
}

func main() {
	Greet(os.Stdout, "Nikita")
}