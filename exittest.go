//go:build ignore
// +build ignore

package main

import "github.com/mackerelio/checkers"

func main() {
	c := checkers.Warning("test")
	c.Name = "exittest"
	c.Exit()
}
