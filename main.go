package main

import (
	"fmt"
	setting "webDesign/pkg"
	"webDesign/pkg/crypto"
	"webDesign/routers"
)

func open() {
	router := routers.InitRouter()

	err := router.Run(fmt.Sprintf(":%d", setting.HTTPPort))
	if err != nil {
		return
	}
}

func test() {
	s := crypto.Encrypt("testtesttesttest")
	fmt.Println(s)
}
func main() {
	open()
	//test()
}
