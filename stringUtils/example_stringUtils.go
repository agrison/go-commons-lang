package stringUtils

import (
	"fmt"
)

func ExampleRemoveEnd() {
	s := "www.domain.com"
	fmt.Println(RemoveEnd(s, ".com"))
	// Output: www.domain
}
