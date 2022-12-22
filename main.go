package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

github.com/Alex920679/wordcount
package main

import (
"bufio"
"fmt"
"os"
"strings"
)

func main() {
	txt, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		err := fmt.Errorf("Something goes wrong")
		fmt.Println(err)
	}
	newtxt := strings.Split(txt, " ")
	fmt.Println(len(newtxt))
}
