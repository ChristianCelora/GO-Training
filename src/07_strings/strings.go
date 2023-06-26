package main

import (
	"fmt";
	"unicode/utf8"
)

func printStringInfo(str string) {
	fmt.Println("str:", str, "len", len(str))
	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range str {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
}

func printStringRunesInfo(str string) {
	fmt.Println("|", str, "|")
	fmt.Println("rune count", utf8.RuneCountInString(str))
	runeValue, runeWidth := utf8.DecodeRuneInString(str)
	fmt.Printf("runes to string <value, width>: <%#U, %d> ", runeValue, runeWidth)
	fmt.Println()
}

/**
* In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
*/
func main() {
	hello_eng := "hello"
	printStringInfo(hello_eng)
	hello_thai := "สวัสดี"
	printStringInfo(hello_thai)

	// utf8 package has some function to help with runes
	printStringRunesInfo(hello_eng)
	printStringRunesInfo(hello_thai)
}