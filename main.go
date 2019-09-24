package main

import (
	"fmt"
	"lhx-github/go_web_demo/dao"
)

func main() {


	dao.InitDb()
	if err := InstanceRoutine().Run(":8080"); err != nil {
		fmt.Println("xxx")
	}

}
