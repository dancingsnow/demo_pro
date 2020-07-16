package demo_pro

import "fmt"

func Hello(name string) {
	fmt.Println("这是V1版本！")
	fmt.Println("v2不兼容" ,"Hello world !", name)
}
