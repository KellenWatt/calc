package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/KellenWatt/calc"
)

const lineHeader = "> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := calc.New()

	var currentOp func(*calc.Calculator)error

	fmt.Print(lineHeader)
	EvalLoop:
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		for _, t := range tokens {
			if f,ok := mappings[t]; ok {
				if currentOp != nil {
					printError("operation already pending")
					continue
				}
				err := f(c)
				if err == calc.StackUnderflowError {
					currentOp = f	
					continue
				} else if err != nil {
					printError(err)
				}
				last, _ := c.Last()
				fmt.Printf("\033[1;36m%g\033[0m\n", last)
			} else {
				switch t {
				case "print", "all", "stack":
					fmt.Println(c)
				case "quit", "exit":
					break EvalLoop
				case "clean":
					c.Sanitize()
				case "help", "?":
					fmt.Println(helptext())
				case "clear", "reset":
					c.Clear()
				case "cancel":
					currentOp = nil
				default:
					n, err := strconv.ParseFloat(t, 64)
					if err != nil {
						printError(err)
						continue
					}
					c.Push(n)
					if currentOp != nil {
						err := currentOp(c)
						if err != nil {
							printError(err)
						} else {
							last, _ := c.Last()
							fmt.Printf("\033[1;36m%g\033[0m\n", last)
						}
						currentOp = nil

					}
				}
			}
		}
		fmt.Print(lineHeader)
	}
}

func helptext() string {
	return "This will eventually be the helptext."
}

func printError(err interface{}) {
	switch err.(type) {
	case string:
		fmt.Fprintln(os.Stderr, "Error:", err.(string))
	case error:
		fmt.Fprintln(os.Stderr, "Error:", err.(error).Error())
	default:
		panic("Invalid type provided to printError")
	}
}
