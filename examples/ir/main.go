package main

import (
	"fmt"

	hiapi "./hisilicon"
)

func main() {
	repkey_timeout := uint32(108)
	fmt.Println("IR")

	if _, err := hiapi.HI_UNF_IR_Init(); err != nil {
		panic(err)
	}
	if _, err := hiapi.HI_UNF_IR_EnableRepKey(hiapi.HI_TRUE); err != nil {
		panic(err)
	}

	hiapi.HI_UNF_IR_SetRepKeyTimeoutAttr(repkey_timeout)

	if status, err := hiapi.HI_UNF_IR_GetProtocolEnabled("tc9012"); err != nil {
		panic(err)
	} else {
		if status {
			fmt.Println("Protocol 'tc9012' is enabled.")
		} else {
			fmt.Println("Protocol 'tc9012' is disabled.")
		}
	}

	if _, err := hiapi.HI_UNF_IR_DisableProtocol("tc9012"); err != nil {
		panic(err)
	}

	if status, err := hiapi.HI_UNF_IR_GetProtocolEnabled("tc9012"); err != nil {
		panic(err)
	} else {
		if status {
			fmt.Println("Protocol 'tc9012' is enabled.")
		} else {
			fmt.Println("Protocol 'tc9012' is disabled.")
		}
	}

	fmt.Println("Listen keys ...")
	for {
		if key, err := hiapi.HI_UNF_IR_GetValueWithProtocol(200); err == nil {
			var status string
			switch key.StatusKey {
			case hiapi.HI_UNF_KEY_STATUS_DOWN:
				status = "DOWN"
			case hiapi.HI_UNF_KEY_STATUS_HOLD:
				status = "HOLD"
			case hiapi.HI_UNF_KEY_STATUS_UP:
				status = "UP"
			default:
				status = "Unknown"
			}

			fmt.Printf("Received Key: %#x, %s,\tProtocol: %s.\n", key.Lower, status, key.ProtocolName)
			if key.Lower == 0x5fa0377d && key.StatusKey == hiapi.HI_UNF_KEY_STATUS_HOLD {
				break
			}

			if key.Lower == 0x3bc4377d && key.StatusKey == hiapi.HI_UNF_KEY_STATUS_DOWN {
				repkey_timeout += 10
				if _, err := hiapi.HI_UNF_IR_SetRepKeyTimeoutAttr(repkey_timeout); err != nil {
					repkey_timeout -= 10
				} else {
					fmt.Printf("Up Repeat Key TimeOut From: %d To: %d.\n", repkey_timeout-10, repkey_timeout)
				}
			}
		}
	}
	if _, err := hiapi.HI_UNF_IR_Reset(); err != nil {
		panic(err)
	}
	if _, err := hiapi.HI_UNF_IR_EnableProtocol("tc9012"); err != nil {
		panic(err)
	}
	if status, err := hiapi.HI_UNF_IR_GetProtocolEnabled("tc9012"); err != nil {
		panic(err)
	} else {
		if status {
			fmt.Println("Protocol 'tc9012' is enabled.")
		} else {
			fmt.Println("Protocol 'tc9012' is disabled.")
		}
	}
	hiapi.HI_UNF_IR_DeInit()
	fmt.Println("Exit application.")
}
