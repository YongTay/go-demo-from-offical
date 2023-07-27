package main

import (
	"fmt"

	"framework/modules/utils"
)

func main() {
	res := utils.Add(1, 2, 3)
	fmt.Println(res)
}
