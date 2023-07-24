package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "a name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'foo' or 'bar' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("sub command 'foo'")
		fmt.Println("enable", *fooEnable)
		fmt.Println("name", *fooName)
		fmt.Println("tail", fooCmd.Args()) // 额外参数
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("level", *barLevel)
		fmt.Println("tail", barCmd.Args()) // 额外参数
	default:
		fmt.Println("Expected 'foo' or 'bar' subcommand")
		os.Exit(1)
	}
}
