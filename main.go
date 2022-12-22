package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	txt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	newtxt := strings.Split(txt, " ")
	fmt.Println(len(newtxt))
}
