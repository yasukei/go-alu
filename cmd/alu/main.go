package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yasukei/go-alu"
)

func main() {
	var not = flag.Bool("n", false, "Not operation for input")
	var ope = flag.String("o", "", "And operation for input")
	//var or = flag.String("o", "", "Or operation for input")
	//var xor = flag.String("x", "", "Xor operation for input")

	flag.Parse()
	arg := flag.Arg(0)

	arg = strings.TrimPrefix(arg, "0x")
	if len(arg)%2 != 0 {
		fmt.Fprintln(os.Stderr, "input string has to be even length")
		return
	}
	bytes, err := hex.DecodeString(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if *not {
		alu.Not(bytes)
	}
	fmt.Print(bytes)

	switch *ope {
	case "and":
		f := alu.And
	case "or":
		f := alu.Or
	case "xor":
		f := alu.Xor
	default:
		fmt.Fprintln(os.Stderr, "invalid operator %s", *ope)
		return
	}

	//
	//reader := strings.NewReader(args)
	////reader := os.Stdin
	//writer := os.Stdout

	//scanner := bufio.NewScanner(reader)

	//for scanner.Scan() {
	//	input := scanner.Bytes()

	//	alu.Not(input)
	//	fmt.Fprint(writer, input)
	//}
}
