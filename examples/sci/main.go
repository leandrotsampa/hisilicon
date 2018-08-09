package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	hiapi "./hisilicon"
)

type ACTION_CARD int

const (
	ACTION_CARDIN  ACTION_CARD = iota /* Card Insert Flag */
	ACTION_CARDOUT ACTION_CARD = iota /* Card Pull Out Flag*/
	ACTION_NONE    ACTION_CARD = iota /* Card Immobile */
)

var (
	mu          sync.Mutex
	sPort       hiapi.HI_UNF_SCI_PORT_E = hiapi.HI_UNF_SCI_PORT0
	bCardStatus bool                    = false /* true = Indicate Card In; false = Indicate Card Out */
	sCardAction ACTION_CARD             = ACTION_NONE
)

/* Monitor Card Status */
func monitor_cstatus() {
	for {
		mu.Lock()

		bStatus := false
		status := hiapi.SCI_STATUS_S{Port: sPort}
		if _, err := hiapi.HI_UNF_SCI_GetCardStatus(&status); err == nil {
			if status.Status <= hiapi.HI_UNF_SCI_STATUS_NOCARD {
				bStatus = false /* No Card */
			} else {
				bStatus = true /* Have Card */
			}

			/* if bStatus is true indicated the card have been pull out or push in */
			if bCardStatus != bStatus {
				bCardStatus = bStatus
				if bStatus {
					sCardAction = ACTION_CARDIN /* Card In  */
				} else {
					sCardAction = ACTION_CARDOUT /* Card Out */
				}
			} else {
				sCardAction = ACTION_NONE /* No Operation */
			}
		}

		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

/* Receive or Send Data to Card */
func monitor_cdata() {
	for {
		mu.Lock()
		CardAction := sCardAction
		mu.Unlock()

		if CardAction == ACTION_CARDIN {
			fmt.Println("CARD IN")
			cardInProcess()
		} else if CardAction == ACTION_CARDOUT {
			fmt.Println("CARD OUT")
			//CardOutProcess();
		}

		time.Sleep(1 * time.Second)
	}
}

func cardInProcess() {
	stReset := hiapi.SCI_RESET_S{Port: sPort, WarmReset: hiapi.HI_FALSE}
	if _, err := hiapi.HI_UNF_SCI_ResetCard(stReset); err != nil {
		fmt.Printf("[RESET] Error: %s\n", err.Error())
		return
	}

	ResetTime := 0
	for {
		/* Will exit reset when reseting out of 10s */
		if ResetTime >= 10 {
			fmt.Println("[RESET] Failure.")
			return
		}

		/* Get SCI Card Status */
		status := hiapi.SCI_STATUS_S{Port: sPort}
		if _, err := hiapi.HI_UNF_SCI_GetCardStatus(&status); err != nil {
			fmt.Printf("[STATUS] Error: %s\n", err.Error())
			return
		}

		if status.Status >= hiapi.HI_UNF_SCI_STATUS_READY {
			/* Reset Success */
			fmt.Println("[RESET] Success.")
			break
		} else {
			//fmt.Printf("[RESET] Waiting ... (%d)\n", status.Status)
			time.Sleep(1 * time.Second)
			ResetTime++
		}
	}

	buf := make([]byte, 255)
	atr := hiapi.SCI_ATR_S{Port: sPort, AtrBuf: &buf[0], AtrBufSize: uint32(len(buf))}
	if _, err := hiapi.HI_UNF_SCI_GetATR(&atr); err != nil {
		fmt.Printf("[ATR] Error: %s\n", err.Error())
		return
	}

	fmt.Printf("[ATR] Size: %d\n", atr.DataLen)
	fmt.Printf("[ATR] Data: %s\n", strings.ToUpper(hex.EncodeToString(buf[:atr.DataLen])))
	rom := ""
	rev := ""
	rtype := ""
	if bytes.ContainsAny(buf[:atr.DataLen], "DNASP24") {
		rtype = "Nagra 3 Card"
	} else if bytes.ContainsAny(buf[:atr.DataLen], "DNASP") {
		rtype = "Nagra Card"
	} else if bytes.ContainsAny(buf[:atr.DataLen], "TIGER") || bytes.ContainsAny(buf[:atr.DataLen], "NCMED") {
		rtype = "Nagra Tiger Card"
	}

	if len(rtype) > 0 && atr.DataLen >= 20 {
		rom = string(buf[11:20])
		rev = string(buf[20:atr.DataLen])
	} else {
		rtype = "Unknown"
	}

	fmt.Printf("[ATR] Type: %s\n", rtype)
	fmt.Printf("[ATR] ROM : %s\n", rom)
	fmt.Printf("[ATR] REV : %s\n", rev)
}

func main() {
	fmt.Println("SCI")

	fmt.Println("    - Init")
	if _, err := hiapi.HI_UNF_SCI_Init(); err != nil {
		panic(err)
	}

	sciattrib := hiapi.SCI_OPEN_S{}
	sciattrib.Port = sPort
	sciattrib.Protocol = hiapi.HI_UNF_SCI_PROTOCOL_T1
	sciattrib.Frequency = 3570

	fmt.Println("    - Open")
	if _, err := hiapi.HI_UNF_SCI_Open(sciattrib); err != nil {
		panic(err)
	}

	fmt.Println("    - Config Clock Mode")
	clkmode := hiapi.SCI_IO_OUTPUTTYPE_S{}
	clkmode.Port = sPort
	clkmode.OutputType = hiapi.HI_UNF_SCI_MODE_CMOS
	if _, err := hiapi.HI_UNF_SCI_ConfigClkMode(clkmode); err != nil {
		panic(err)
	}

	fmt.Println("    - Config VCC Mode")
	vccmode := hiapi.SCI_IO_OUTPUTTYPE_S{}
	vccmode.Port = sPort
	vccmode.OutputType = hiapi.HI_UNF_SCI_MODE_CMOS
	if _, err := hiapi.HI_UNF_SCI_ConfigVccEnMode(vccmode); err != nil {
		panic(err)
	}

	fmt.Println("    - Config Reset Mode")
	resetmode := hiapi.SCI_IO_OUTPUTTYPE_S{}
	resetmode.Port = sPort
	resetmode.OutputType = hiapi.HI_UNF_SCI_MODE_CMOS
	if _, err := hiapi.HI_UNF_SCI_ConfigResetMode(resetmode); err != nil {
		panic(err)
	}

	fmt.Println("    - Config VCC Level")
	vcclevel := hiapi.SCI_LEVEL_S{}
	vcclevel.Port = sPort
	vcclevel.Level = hiapi.HI_UNF_SCI_LEVEL_LOW
	if _, err := hiapi.HI_UNF_SCI_ConfigVccEn(vcclevel); err != nil {
		panic(err)
	}

	fmt.Println("    - Config Detect Level")
	detectlevel := hiapi.SCI_LEVEL_S{}
	detectlevel.Port = sPort
	detectlevel.Level = hiapi.HI_UNF_SCI_LEVEL_HIGH
	if _, err := hiapi.HI_UNF_SCI_ConfigDetect(detectlevel); err != nil {
		panic(err)
	}

	fmt.Println("    - Starting Monitor for Card Status")
	go monitor_cstatus()

	fmt.Println("    - Starting Monitor for Receive/Send Data")
	go monitor_cdata()

	fmt.Println("Press any key to finish SCI demo.")
	fmt.Scanln()
	hiapi.HI_UNF_SCI_DeactiveCard(sPort)
	hiapi.HI_UNF_SCI_Close(sPort)
	hiapi.HI_UNF_SCI_DeInit()
}
