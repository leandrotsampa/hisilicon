package main

import (
	"fmt"

	hiapi "./hisilicon"
)

func main() {
	fmt.Println("MODULE")

	fmt.Println("    - Init")
	if _, err := hiapi.HI_MODULE_Init(); err != nil {
		panic(err)
	}

	var id uint32
	var err error
	var name string

	fmt.Println("    - Get Module ID")
	if id, err = hiapi.HI_MODULE_GetModuleID("HI_SCI"); err != nil {
		panic(err)
	} else {
		fmt.Printf("    - Module ID = %#x\n", id)
	}

	fmt.Println("    - Get Module Name")
	if name, err = hiapi.HI_MODULE_GetModuleName(id); err != nil {
		panic(err)
	} else {
		fmt.Printf("    - Module Name = %s\n", name)
	}

	hiapi.HI_MODULE_DeInit()
}
