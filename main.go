package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func win() {
	fmt.Println("Congratulations, you win!")
	os.Exit(0)
}

func lose() {
	fmt.Println("Oops, you lose.")
	os.Exit(0)
}

var (
	bullet     map[uint8]bool
	l, b       uint8
	curl, curb uint8
)

func handleCommand(cmd string) bool {
	if len(cmd) == 1 {
		switch cmd {
		case "w":
			win()
		case "l":
			lose()
		case "v":
			view()
		case "n":
			return true
		default:
			fmt.Println("No such command:", cmd)
		}
		return false
	} else if len(cmd) == 2 {
		var n uint8
		var c rune
		if unicode.IsDigit(rune(cmd[0])) {
			n = cmd[0] - '0'
			c = rune(cmd[1])
		} else {
			c = rune(cmd[0])
			n = cmd[1] - '0'
		}
		switch c {
		case 'b':
			if n > 0 && n <= b+l {
				bullet[n] = false
				fmt.Printf("b: Position %d is BLANK\n", n)
			} else {
				fmt.Printf("-b: The data [%d] is out of range\n", n)
			}

		case 'l':
			if n > 0 && n <= b+l {
				bullet[n] = true
				fmt.Printf("l: Position %d is LIVE\n", n)
			} else {
				fmt.Printf("-l: The data [%d] is out of range\n", n)
			}

		default:
			fmt.Println("No such command:", cmd)
			return false
		}
	}
	return false
}

func view() {
	fmt.Printf("Now we have %d LIVE shell, %d BLANK shell.\n", curl, curb)
}

func loop(n uint8) {
	for i := 0; i < int(n); i++ {
		fmt.Printf("Now we have %d LIVE shell, %d BLANK shell.\n", curl, curb)
		fmt.Printf("Current shell is %s at posi %d.\n", checkShell(uint8(i+1)), i+1)
		for {
			fmt.Print("> ")
			var cmd string
			fmt.Scanln(&cmd)
			if strings.TrimSpace(cmd) != "" {
				ok := handleCommand(cmd)
				if ok {
					break
				}
			}
		}
		if status := checkShell(uint8(i + 1)); status == "LIVE SHELL" {
			curl--
		} else if status == "BLANK SHELL" {
			curb--
		}
	}
}

func checkShell(n uint8) string {
	v, ok := bullet[n]
	if ok {
		if v {
			return "LIVE SHELL"
		} else {
			return "BLANK SHELL"
		}
	}
	return "UNKNOWN"
}

func main() {
	fmt.Print("Enter the number of shells: ")
	var t uint8
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if t < 11 || t > 71 {
		log.Fatalln("The data you enter is invalid.")
	}
	l, b = t/10, t%10
	curl, curb = t/10, t%10
	//if l+b > 8 || l < 1 || l > 7 || b < 1 || b > 7 {
	//	log.Fatal("The data you enter is invalid.")
	//}
	bullet = make(map[uint8]bool, 8)
	loop(l + b)
}
