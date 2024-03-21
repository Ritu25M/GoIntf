package main

import (
	"fmt"
	"goIntf/species"
)

/*
main invokes the method 'GenerateTag' from the species package
*/
func main() {
	const lion = "lion"
	const cow = "cow"
	fmt.Println(species.GenerateTag(lion)) // Output: U+'timestamp'
	fmt.Println(species.GenerateTag(cow))  // Output: HR+'timestamp'
}
