package main

import (
	"fmt"

	"mybtc/bcapi"
)

func main() {
	client := bcapi.NewSimpleClient()
	for i := uint32(0); i < 10000; i++ {
		rawBlock, err := client.GetRawBlockByHeight(i)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("block#%d: %s\n", i, rawBlock)
	}
}
