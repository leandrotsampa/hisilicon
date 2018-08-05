package hisilicon

import (
	"encoding/binary"
	"errors"
	"os"
	"sync"
)

var (
	mu     sync.Mutex
	fd     *os.File = nil
	in_use int      = 0
)

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
	mu.Lock()
	defer mu.Unlock()

	if fd != nil {
		in_use++
		return true, nil
	}

	var err error
	if fd, err = os.OpenFile("/dev/hi_ir", os.O_RDWR, 0); err != nil {
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
	mu.Lock()
	defer mu.Unlock()

	if fd == nil {
		return true, nil
	} else if in_use > 0 {
		in_use--
		return true, nil
	}

	if err := fd.Close(); err != nil {
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
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	var status int32
	if bEnable {
		status = 1
	}

	if err := Ioctl(fd.Fd(), CMD_IR_SET_ENABLE, &status); err != nil {
		return false, err
	}

	return true, nil
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

	if fd == nil {
		return IrKey, errors.New("IR Device not initialized.")
	}

	if err := Ioctl(fd.Fd(), CMD_IR_SET_BLOCKTIME, &u32TimeoutMs); err != nil {
		return IrKey, err
	}

	if err := binary.Read(fd, binary.LittleEndian, &IrKey); err != nil {
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
Input:          mode: true-> key mode. false-> raw symbol mode.
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_IR_SetFetchMode(s32Mode HI_BOOL) (bool, error) {
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	var mode int32
	if s32Mode {
		mode = 1
	}

	if err := Ioctl(fd.Fd(), CMD_IR_SET_FETCH_METHOD, &mode); err != nil {
		return false, err
	}

	return true, nil
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
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	var status int32
	if bEnable {
		status = 1
	}

	if err := Ioctl(fd.Fd(), CMD_IR_ENABLE_KEYUP, &status); err != nil {
		return false, err
	}

	return true, nil
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
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	var status int32
	if bEnable {
		status = 1
	}

	if err := Ioctl(fd.Fd(), CMD_IR_ENABLE_REPKEY, &status); err != nil {
		return false, err
	}

	return true, nil
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
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	if u32TimeoutMs > 65536 {
		return false, errors.New("The max timeout supported is 65536ms.")
	}

	if err := Ioctl(fd.Fd(), CMD_IR_SET_REPKEY_TIMEOUT, &u32TimeoutMs); err != nil {
		return false, err
	}

	return true, nil
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
	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	if err := Ioctl(fd.Fd(), CMD_IR_RESET, nil); err != nil {
		return false, err
	}

	return true, nil
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

	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	bProtocol := []byte(pszProtocolName)
	if err := Ioctl(fd.Fd(), CMD_IR_SET_PROT_ENABLE, &bProtocol[0]); err != nil {
		return false, err
	}

	return true, nil
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

	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	bProtocol := []byte(pszProtocolName)
	if err := Ioctl(fd.Fd(), CMD_IR_SET_PROT_DISABLE, &bProtocol[0]); err != nil {
		return false, err
	}

	return true, nil
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

	if fd == nil {
		return false, errors.New("IR Device not initialized.")
	}

	bProtocol := []byte(pszProtocolName)
	if err := Ioctl(fd.Fd(), CMD_IR_GET_PROT_ENABLED, &bProtocol[0]); err != nil {
		return false, err
	}

	if bProtocol[0] == 0 {
		return true, nil
	} else if bProtocol[0] == 1 {
		return false, nil
	}

	return false, errors.New("Data return is invalid, need check this method!")
}
