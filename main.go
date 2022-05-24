package main

import (
	"hello-world/subpkg"
)

func main() {
	h := human{}
	h.Say("baby!")
}

type human struct {
	subpkg.People
}
