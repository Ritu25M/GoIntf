package main

import (
	"fmt"
	"goIntf/species"
)

func main() {
	fmt.Println(species.GenerateTag("lion")) // Output: Unknown
	fmt.Println(species.GenerateTag("cow"))  // Output: HR
}
