package subpkg

import "fmt"

type People struct {
}

func (h People) Say(s string) {
	fmt.Println("hello ", s)
}
