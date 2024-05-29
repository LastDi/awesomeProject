package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var X rune = 88
var Y rune = 89
var Z rune = 90

type Node struct {
	data int
	next *Node
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	var sb strings.Builder
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		sb.WriteString(find(&s))
		sb.WriteString("\n")
		//res := find(&s)
		//fmt.Println(res)
	}
	//fmt.Println(sb.String())
	//fmt.Println()

	d1 := []byte(sb.String())
	err := os.WriteFile("text.txt", d1, 0644)
	fmt.Println(err)
}

func find(str *string) string {
	arr := []rune(*str)
	var x *Node
	var bx *Node
	var y *Node
	var by *Node
	xc := 0
	yc := 0
	zc := 0
	for i, ch := range *str {
		switch ch {
		case X:
			if bx == nil {
				x = &Node{data: i, next: nil}
				bx = x
			} else {
				tmp := Node{data: i, next: nil}
				x.next = &tmp
				x = &tmp
			}
			xc++
		case Y:
			if by == nil {
				y = &Node{data: i, next: nil}
				by = y
			} else {
				tmp := Node{data: i, next: nil}
				y.next = &tmp
				y = &tmp
			}
			yc++
		case Z:
			zc++
		}
	}
	for i, ch := range *str {
		switch ch {
		case Z:
			if xc == yc+zc {
				// todo cycle for
				for bx != nil && by != nil && by.data > bx.data {
					bx = bx.next
					by = by.next
				}
				if bx != nil && by != nil && bx.next != nil && bx.next.data > by.data && bx.next.data < i {
					//fmt.Println("--------")
					arr[bx.data] = -2
					arr[by.data] = -2
					bx = bx.next
					by = by.next
					xc--
					yc--
				}
				//fmt.Println("+++++++++")
				arr[i] = -1
				arr[bx.data] = -1
				bx = bx.next
				xc--
			} else if yc > 0 && by.data < i {
				arr[i] = -1
				arr[by.data] = -1
				by = by.next
				yc--
			} else if xc > 0 && bx.data < i {
				arr[i] = -1
				arr[bx.data] = -1
				bx = bx.next
				xc--
			} else {
				return "No"
			}
			zc--
			//fmt.Println(arr)
		}
	}

	for bx != nil && by != nil {
		if bx.data >= by.data {
			return "No"
		}
		bx = bx.next
		by = by.next
	}
	if bx != nil || by != nil {
		return "No"
	}
	return "Yes"
}
