// intsToString 和 fmt.Sprint(values)类似但是插入了,
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3, 4})) // [1, 2, 3, 4]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for k, v := range values {
		if k > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}