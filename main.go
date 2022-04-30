package main

import (
	"fmt"
	setting "webDesign/pkg"
	"webDesign/routers"
)

func main() {
	router := routers.InitRouter()

	err := router.Run(fmt.Sprintf(":%d", setting.HTTPPort))
	if err != nil {
		return
	}
}
