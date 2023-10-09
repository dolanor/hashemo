// hashemo util transform hexadecimal sum outputs from traditional checksum tools
// into a emoji hash.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dolanor/hashemo"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		hash, fileName, ok := strings.Cut(line, " ")
		if !ok {
			panic("non traditional checksum tool output:\n\t<checksum in hex> <file name>")
		}
		hashemo := hashemo.FromHexString(hash)
		fmt.Println(hashemo, fileName)
	}
}
