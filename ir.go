package hisilicon

import (
	"encoding/binary"
	"errors"
	"os"
	"strings"
	"unsafe"
)

var (
	ir = HiDevice{name: "hi_ir", fd: nil, InUse: 0}

	/** Ioctl Definitions **/
	/* 1:check keyup */
	CMD_IR_ENABLE_KEYUP uintptr

	/* 1:check repkey, 0:hardware behave */
	CMD_IR_ENABLE_REPKEY      uintptr
	CMD_IR_SET_REPKEY_TIMEOUT uintptr

	/* 1:enable ir, 0:disable ir */
	CMD_IR_SET_ENABLE    uintptr
	CMD_IR_RESET         uintptr
	CMD_IR_SET_BLOCKTIME uintptr
	CMD_IR_SET_FORMAT    uintptr
	CMD_IR_SET_BUF       uintptr

	/* raw symbol fetch(1) or key fetch(0) */
	CMD_IR_SET_FETCH_METHOD uintptr

	/* enable or disalbe a protocol */
	CMD_IR_SET_PROT_ENABLE  uintptr
	CMD_IR_SET_PROT_DISABLE uintptr
	CMD_IR_GET_PROT_ENABLED uintptr
)

/** Internal Function for IR IOCTL Calls **/
func irLoadIoctl() error {
	var id uint32
	var err error

	if ir.fd != nil {
		return nil
	} else if _, err = HI_MODULE_Init(); err != nil {
		return err
	} else if id, err = HI_MODULE_GetModuleID(strings.ToUpper(ir.name)); err != nil {
		return err
	}

	CMD_IR_ENABLE_KEYUP = IoW(uintptr(id), 0x1, unsafe.Sizeof(int32(0)))
	CMD_IR_ENABLE_REPKEY = IoW(uintptr(id), 0x2, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_REPKEY_TIMEOUT = IoW(uintptr(id), 0x3, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_ENABLE = IoW(uintptr(id), 0x4, unsafe.Sizeof(int32(0)))
	CMD_IR_RESET = Io(uintptr(id), 0x5)
	CMD_IR_SET_BLOCKTIME = IoW(uintptr(id), 0x6, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_FORMAT = IoW(uintptr(id), 0x7, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_BUF = IoW(uintptr(id), 0x8, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_FETCH_METHOD = IoW(uintptr(id), 0x9, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_PROT_ENABLE = IoW(uintptr(id), 0xa, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_PROT_DISABLE = IoW(uintptr(id), 0xb, unsafe.Sizeof(int32(0)))
	CMD_IR_GET_PROT_ENABLED = IoRW(uintptr(id), 0xc, unsafe.Sizeof(int32(0)))

	HI_MODULE_DeInit()
	return nil
}

func irCall(op uintptr, arg interface{}) (bool, error) {
	if ir.fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	if err := Ioctl(ir.fd.Fd(), op, arg); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_IR_Init
Description:    open ir device,and do the basical initialization
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_Init() (bool, error) {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	if ir.fd != nil {
		ir.InUse++
		return true, nil
	}

	var err error
	if err = irLoadIoctl(); err != nil {
		return false, err
	} else if ir.fd, err = os.OpenFile("/dev/"+ir.name, os.O_RDWR, 0); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_IR_DeInit
Description:    close ir device
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_DeInit() (bool, error) {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	if ir.fd == nil {
		return true, nil
	} else if ir.InUse > 0 {
		ir.InUse--
		return true, nil
	}

	if err := ir.fd.Close(); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_IR_Enable
Description:    Enable ir device
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_Enable(bEnable HI_BOOL) (bool, error) {
	if bEnable != HI_FALSE && bEnable != HI_TRUE {
		return false, errors.New("The parameter is invalid.")
	}

	return irCall(CMD_IR_SET_ENABLE, &bEnable)
}

/*************************************************************
Function:       HI_UNF_IR_GetProtocol
Description:    get the type of ir protocol
Calls:
Data Accessed:
Data Updated:   NA
Input:          NA
Output:         penProtocol:    ir protocol of the key
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_GetProtocol(penProtocol *HI_UNF_IR_PROTOCOL_E) (bool, error) {
	return false, errors.New("Not implemented.")
}

/*************************************************************
Function:       HI_UNF_IR_GetProtocolName
Description:    reserved
Calls:
Data Accessed:
Data Updated:   NA
Input:          pProtocolName, s32BufLen
Output:         pProtocolName
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_GetProtocolName(pProtocolName *HI_CHAR, s32BufLen HI_S32) (bool, error) {
	return false, errors.New("Not implemented.")
}

/*************************************************************
Function:       HI_UNF_IR_GetValueWithProtocol
Description:    get the value and status of key
Calls:
Data Accessed:
Data Updated:   NA
Input:          u32TimeoutMs: overtime value with unit of ms : 0 means no block while 0xFFFFFFFF means block forever
Output:
Return:         KeyAttr
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_GetValueWithProtocol(u32TimeoutMs HI_U32) (KeyAttr, error) {
	IrKey := KeyAttr{}

	if _, err := irCall(CMD_IR_SET_BLOCKTIME, &u32TimeoutMs); err != nil {
		return IrKey, err
	}

	if err := binary.Read(ir.fd, binary.LittleEndian, &IrKey); err != nil {
		return IrKey, err
	}

	return IrKey, nil
}

/*************************************************************
Function:       HI_UNF_IR_SetFetchMode
Description:    set key fetch mode or symbol mode.
Calls:
Data Accessed:
Data Updated:   NA
Input:          bMode: true-> key mode. false-> raw symbol mode.
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_SetFetchMode(bMode HI_BOOL) (bool, error) {
	if bMode != HI_FALSE && bMode != HI_TRUE {
		return false, errors.New("The parameter is invalid.")
	}

	return irCall(CMD_IR_SET_FETCH_METHOD, &bMode)
}

/*************************************************************
Function:       HI_UNF_IR_GetSymbol
Description:    get one raw symbols from ir module.
Calls:
Data Accessed:
Data Updated:   NA
Input:          u32TimeoutMs: read timeout in ms.
Output:         pu64lower, pu64upper.
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_GetSymbol(pu64First *HI_U64, pu64Second *HI_U64, u32TimeoutMs HI_U32) (bool, error) {
	if pu64First == nil || pu64Second == nil {
		return false, errors.New("Invalid parameters!")
	}

	key, err := HI_UNF_IR_GetValueWithProtocol(u32TimeoutMs)
	if err != nil {
		return false, err
	}

	*pu64First = key.Lower
	*pu64Second = key.Upper

	return true, nil
}

/*************************************************************
Function:       HI_UNF_IR_EnableKeyUp
Description:    config whether report the state of key release
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_EnableKeyUp(bEnable HI_BOOL) (bool, error) {
	if bEnable != HI_FALSE && bEnable != HI_TRUE {
		return false, errors.New("The parameter is invalid.")
	}

	return irCall(CMD_IR_ENABLE_KEYUP, &bEnable)
}

/*************************************************************
Function:       HI_UNF_IR_EnableRepKey
Description:    config whether report repeat key
Calls:
Data Accessed:
Data Updated:   NA
Input:          bEnable
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_EnableRepKey(bEnable HI_BOOL) (bool, error) {
	if bEnable != HI_FALSE && bEnable != HI_TRUE {
		return false, errors.New("The parameter is invalid.")
	}

	return irCall(CMD_IR_ENABLE_REPKEY, &bEnable)
}

/*************************************************************
Function:       HI_UNF_IR_RepKeyTimeoutVal
Description:    Set the reporting interval when you keep pressing button.
Calls:
Data Accessed:
Data Updated:   NA
Input:          u32TimeoutMs  The minimum interval to report repeat key
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_SetRepKeyTimeoutAttr(u32TimeoutMs HI_U32) (bool, error) {
	if u32TimeoutMs > 65536 {
		return false, errors.New("The max timeout supported is 65536ms.")
	}

	return irCall(CMD_IR_SET_REPKEY_TIMEOUT, &u32TimeoutMs)
}

/*************************************************************
Function:       HI_UNF_IR_SetCodeType
Description:    reserved interface.
Calls:
Data Accessed:
Data Updated:   NA
Input:          enIRCode
Output:
Return:         bool

Others:         NA
*************************************************************/
func HI_UNF_IR_SetCodeType(enIRCode HI_UNF_IR_CODE_E) bool {
	return true
}

/*************************************************************
Function:       HI_UNF_IR_Reset
Description:    Reset ir device
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_Reset() (bool, error) {
	return irCall(CMD_IR_RESET, nil)
}

/*************************************************************
Function:       HI_UNF_IR_EnableProtocol
Description:    enable an infrared code specified by @prot_name
Calls:
Data Accessed:
Data Updated:   NA
Input:          pszProtocolName: infrared code name.
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_EnableProtocol(pszProtocolName string) (bool, error) {
	if len(pszProtocolName) <= 0 {
		return false, errors.New("Invalid parameter.")
	}

	bProtocol := []byte(pszProtocolName)
	return irCall(CMD_IR_SET_PROT_ENABLE, &bProtocol[0])
}

/*************************************************************
Function:       HI_UNF_IR_DisableProtocol
Description:    disable a infrared code specified by @prot_name
Calls:
Data Accessed:
Data Updated:   NA
Input:          pszProtocolName: infrared code name.
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_DisableProtocol(pszProtocolName string) (bool, error) {
	if len(pszProtocolName) <= 0 {
		return false, errors.New("Invalid parameter.")
	}

	bProtocol := []byte(pszProtocolName)
	return irCall(CMD_IR_SET_PROT_DISABLE, &bProtocol[0])
}

/*************************************************************
Function:       HI_UNF_IR_GetProtocolEnabled
Description:    get the enable status of an infrared code specified by @prot_name
Calls:
Data Accessed:
Data Updated:   NA
Input:          pszProtocolName: infrared code name.
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_GetProtocolEnabled(pszProtocolName string) (bool, error) {
	if len(pszProtocolName) <= 0 {
		return false, errors.New("Invalid parameter.")
	}

	bProtocol := []byte(pszProtocolName)
	if _, err := irCall(CMD_IR_GET_PROT_ENABLED, &bProtocol[0]); err != nil {
		return false, err
	}

	if bProtocol[0] == 0 {
		return true, nil
	} else if bProtocol[0] == 1 {
		return false, nil
	}

	return false, errors.New("Data return is invalid, need check this method!")
}
