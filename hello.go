package demo_pro

import "fmt"

func Hello(name string) error {
	fmt.Println("v2版本操作")
	fmt.Println("v2不兼容" ,"Hello world !", name)
	return nil
}
