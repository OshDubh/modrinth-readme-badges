/*
Filename: main.go
Description: routes api + requests badges
Created by: main
        at: 12:12 on Thursday, the 29th of January, 2026.
Last edited 13:21 on Thursday, the 29th of January, 2026
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(GenerateBadge(os.Args[1] == "true", os.Args[2], os.Args[3]))
}
