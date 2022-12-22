package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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