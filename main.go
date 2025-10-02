package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите что-то: ")
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		fmt.Println(reader.Text())
	}
}
