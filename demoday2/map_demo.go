package demoday2

import "fmt"

func MapDemo() {

	m := map[string]string{
		"hello": "world",
		"a":     "a",
		"b":     "b",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
