package hisilicon

import (
	"errors"
	"fmt"
	"os"
)

var sci = HiDevice{fd: nil}

/** Internal Function for SCI IOCTL Calls **/
func sciCall(op uintptr, arg interface{}) (bool, error) {
	if sci.fd == nil {
		return false, errors.New("SCI Device not initialized.")
	}

	if err := Ioctl(sci.fd.Fd(), op, arg); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_SCI_Init
Description:    open sci device,and do the basical initialization
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_Init() (bool, error) {
	sci.mu.Lock()
	defer sci.mu.Unlock()

	if sci.fd != nil {
		sci.InUse++
		return true, nil
	}

	var err error
	if sci.fd, err = os.OpenFile("/dev/hi_sci", os.O_RDWR, 0); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_SCI_DeInit
Description:    close sci device
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_DeInit() (bool, error) {
	sci.mu.Lock()
	defer sci.mu.Unlock()

	if sci.fd == nil {
		return true, nil
	} else if sci.InUse > 0 {
		sci.InUse--
		return true, nil
	}

	if err := sci.fd.Close(); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_SCI_Open
Description:    open SCI device
Calls:			HI_SCI_Open
Data Accessed:	NA
Data Updated:   NA
Input:			config
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_Open(config SCI_OPEN_S) (bool, error) {
	if config.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if config.Protocol >= HI_UNF_SCI_PROTOCOL_BUTT {
		return false, errors.New("SCI Protocol is invalid.")
	}

	if config.Protocol == HI_UNF_SCI_PROTOCOL_T14 {
		if config.Frequency < 1000 || config.Frequency > 6000 {
			return false, errors.New("SCI Frequency is invalid.")
		}
	} else if config.Frequency < 1000 || config.Frequency > 5000 {
		return false, errors.New("SCI Frequency is invalid.")
	}

	return sciCall(CMD_SCI_OPEN, &config)
}

/*************************************************************
Function:       HI_UNF_SCI_Close
Description:    close SCI device
Calls:			HI_SCI_Close
Data Accessed:	NA
Data Updated:   NA
Input:			port
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_Close(port HI_UNF_SCI_PORT_E) (bool, error) {
	if port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	return sciCall(CMD_SCI_CLOSE, &port)
}

/*************************************************************
Function:       HI_UNF_SCI_ResetCard
Description:    Reset Card
Calls:			HI_UNF_SCI_ResetCard
Data Accessed:	NA
Data Updated:   NA
Input:			reset
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ResetCard(reset SCI_RESET_S) (bool, error) {
	if reset.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if reset.WarmReset != HI_FALSE && reset.WarmReset != HI_TRUE {
		return false, errors.New("SCI WarmReset is invalid.")
	}

	return sciCall(CMD_SCI_RESET, &reset)
}

/*************************************************************
Function:       HI_UNF_SCI_DeactiveCard
Description:    Deactive Card
Calls:			HI_SCI_DeactiveCard
Data Accessed:	NA
Data Updated:   NA
Input:			reset
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_DeactiveCard(port HI_UNF_SCI_PORT_E) (bool, error) {
	if port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	return sciCall(CMD_SCI_DEACTIVE, &port)
}

/*************************************************************
Function:       HI_UNF_SCI_GetATR
Description:    Get ATR Data
Calls:          HI_SCI_GetATR
Data Accessed:  NA
Data Updated:   NA
Input:          atr
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_GetATR(atr *SCI_ATR_S) (bool, error) {
	if atr.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if atr.AtrBuf == nil {
		return false, errors.New("The buffer is null.")
	}

	if atr.AtrBufSize <= 0 {
		return false, errors.New("The buffer size is invalid.")
	}

	return sciCall(CMD_SCI_GET_ATR, atr)
}

/*************************************************************
Function:       HI_UNF_SCI_GetCardStatus
Description:    Get the Status of Card
Calls:          HI_UNF_SCI_GetCardStatus
Data Accessed:  NA
Data Updated:   NA
Input:          status
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_GetCardStatus(status *SCI_STATUS_S) (bool, error) {
	if status.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	return sciCall(CMD_SCI_GET_STATUS, status)
}

/*************************************************************
Function:       HI_UNF_SCI_Send
Description:    Send Data to Card
Calls:          HI_SCI_Send
Data Accessed:  NA
Data Updated:   NA
Input:          data
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_Send(data *SCI_DATA_S) (bool, error) {
	if data.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if data.DataBuf == nil {
		return false, errors.New("The buffer is null.")
	}

	if data.BufSize <= 0 || data.BufSize > 512 {
		return false, errors.New("The buffer size is invalid.")
	}

	status := SCI_STATUS_S{Port: data.Port}
	if _, err := HI_UNF_SCI_GetCardStatus(&status); err != nil {
		return false, err
	} else if status.Status < HI_UNF_SCI_STATUS_READY {
		return false, errors.New("The current state can't execute send operation.")
	}

	if _, err := sciCall(CMD_SCI_SEND_DATA, data); err != nil {
		return false, err
	}

	if data.DataLen < data.BufSize {
		return false, errors.New(fmt.Sprintf("Not all data is writed, requested to write size %d and writed %d.", data.BufSize, data.DataLen))
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_SCI_Receive
Description:    Receive Data from Card
Calls:          HI_SCI_Receive
Data Accessed:  NA
Data Updated:   NA
Input:          data
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_Receive(data *SCI_DATA_S) (bool, error) {
	if data.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if data.DataBuf == nil {
		return false, errors.New("The buffer is null.")
	}

	if data.BufSize <= 0 || data.BufSize > 256 {
		return false, errors.New("The buffer size is invalid.")
	}

	status := SCI_STATUS_S{Port: data.Port}
	if _, err := HI_UNF_SCI_GetCardStatus(&status); err != nil {
		return false, err
	} else if status.Status < HI_UNF_SCI_STATUS_READATR {
		return false, errors.New("The current state can't execute receive operation.")
	}

	if _, err := sciCall(CMD_SCI_RECEIVE_DATA, data); err != nil {
		return false, err
	} else if data.DataLen < data.BufSize {
		return false, errors.New(fmt.Sprintf("The received size is wrong, received %d expected %d.", data.DataLen, data.BufSize))
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_SCI_ConfigVccEn
Description:    Set the valid level of VCC
Calls:          HI_SCI_ConfigVccEn
Data Accessed:  NA
Data Updated:   NA
Input:          level
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ConfigVccEn(level SCI_LEVEL_S) (bool, error) {
	if level.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if level.Level >= HI_UNF_SCI_LEVEL_BUTT {
		return false, errors.New("The VCC Level is invalid.")
	}

	return sciCall(CMD_SCI_CONF_VCC, &level)
}

/*************************************************************
Function:       HI_UNF_SCI_ConfigDetect
Description:    Set the valid level of Detect
Calls:          HI_SCI_ConfigDetect
Data Accessed:  NA
Data Updated:   NA
Input:          level
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ConfigDetect(level SCI_LEVEL_S) (bool, error) {
	if level.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if level.Level >= HI_UNF_SCI_LEVEL_BUTT {
		return false, errors.New("The Detect Level is invalid.")
	}

	return sciCall(CMD_SCI_CONF_DETECT, &level)
}

/*************************************************************
Function:       HI_UNF_SCI_ConfigClkMode
Description:    Config CLK Work Mode(OD or CMOS)
Calls:          HI_SCI_ConfigClkMode
Data Accessed:  NA
Data Updated:   NA
Input:          mode
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ConfigClkMode(mode SCI_IO_OUTPUTTYPE_S) (bool, error) {
	if mode.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if mode.OutputType >= HI_UNF_SCI_MODE_BUTT {
		return false, errors.New("The CLK Mode is invalid.")
	}

	mode.IO = SCI_IO_CLK
	return sciCall(CMD_SCI_CONF_MODE, &mode)
}

/*************************************************************
Function:       HI_UNF_SCI_ConfigResetMode
Description:    Config Reset Work Mode (OD or CMOS)
Calls:          HI_UNF_SCI_ConfigResetMode
Data Accessed:  NA
Data Updated:   NA
Input:          mode
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ConfigResetMode(mode SCI_IO_OUTPUTTYPE_S) (bool, error) {
	if mode.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if mode.OutputType >= HI_UNF_SCI_MODE_BUTT {
		return false, errors.New("The Reset Mode is invalid.")
	}

	mode.IO = SCI_IO_RESET
	return sciCall(CMD_SCI_CONF_MODE, &mode)
}

/*************************************************************
Function:       HI_UNF_SCI_ConfigVccEnMode
Description:    Config VCC Work Mode (OD or CMOS)
Calls:          HI_UNF_SCI_ConfigVccEnMode
Data Accessed:  NA
Data Updated:   NA
Input:          mode
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_ConfigVccEnMode(mode SCI_IO_OUTPUTTYPE_S) (bool, error) {
	if mode.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if mode.OutputType >= HI_UNF_SCI_MODE_BUTT {
		return false, errors.New("The VCC Mode is invalid.")
	}

	mode.IO = SCI_IO_VCC_EN
	return sciCall(CMD_SCI_CONF_MODE, &mode)
}

/*************************************************************
Function:       HI_UNF_SCI_SwitchCard
Description:    Change Card
Calls:          HI_UNF_SCI_SwitchCard
Data Accessed:  NA
Data Updated:   NA
Input:          config
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SwitchCard(config SCI_OPEN_S) (bool, error) {
	if config.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if config.Protocol >= HI_UNF_SCI_PROTOCOL_BUTT {
		return false, errors.New("SCI Protocol is invalid.")
	}

	if config.Protocol == HI_UNF_SCI_PROTOCOL_T14 {
		if config.Frequency < 1000 || config.Frequency > 6000 {
			return false, errors.New("SCI Frequency is invalid.")
		}
	} else if config.Frequency < 1000 || config.Frequency > 5000 {
		return false, errors.New("SCI Frequency is invalid.")
	}

	return sciCall(CMD_SCI_SWITCH, &config)
}

/*************************************************************
Function:       HI_UNF_SCI_SetEtuFactor
Description:    Set Work BaudRate
Calls:          HI_UNF_SCI_SetEtuFactor
Data Accessed:  NA
Data Updated:   NA
Input:          baudrate
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SetEtuFactor(baudrate SCI_EXT_BAUD_S) (bool, error) {
	if baudrate.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if baudrate.ClkRate < 372 || baudrate.ClkRate > 2048 {
		return false, errors.New("The ClockRate is invalid.")
	}

	if baudrate.BitRate < 1 || baudrate.BitRate > 32 || ((baudrate.BitRate != 1) && (baudrate.BitRate%2) != 0) {
		return false, errors.New("The BitRate is invalid.")
	}

	return sciCall(CMD_SCI_SET_BAUD, &baudrate)
}

/*************************************************************
Function:       HI_UNF_SCI_SetGuardTime
Description:    Set Guard Delay Time
Calls:          HI_UNF_SCI_SetGuardTime
Data Accessed:  NA
Data Updated:   NA
Input:          guard
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SetGuardTime(guard SCI_ADD_GUARD_S) (bool, error) {
	if guard.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if guard.AddCharGuard > 254 {
		return false, errors.New("The Guard Time is invalid.")
	}

	return sciCall(CMD_SCI_SET_CHGUARD, &guard)
}

/*************************************************************
Function:       HI_UNF_SCI_NegotiatePPS
Description:    Request PPS Negotiation
Calls:          HI_UNF_SCI_NegotiatePPS
Data Accessed:  NA
Data Updated:   NA
Input:          pps
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_NegotiatePPS(pps SCI_PPS_S) (bool, error) {
	if pps.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if pps.SendLen < 5 {
		return false, errors.New("The SendLen is invalid.")
	}

	if pps.RecTimeouts < 1 || pps.RecTimeouts > 10000 {
		return false, errors.New("The RecTimeouts is invalid.")
	}

	status := SCI_STATUS_S{Port: pps.Port}
	if _, err := HI_UNF_SCI_GetCardStatus(&status); err != nil {
		return false, err
	} else if status.Status < HI_UNF_SCI_STATUS_READY {
		return false, errors.New("The current state can't execute send operation.")
	}

	return sciCall(CMD_SCI_SEND_PPS_DATA, &pps)
}

/*************************************************************
Function:       HI_UNF_SCI_GetPPSResponData
Description:    Get PPS Negotiation Respond Data
Calls:          HI_UNF_SCI_GetPPSResponData
Data Accessed:  NA
Data Updated:   NA
Input:          pps
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_GetPPSResponData(pps *SCI_PPS_S) (bool, error) {
	if pps.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	return sciCall(CMD_SCI_GET_PPS_DATA, pps)
}

/*************************************************************
Function:       HI_UNF_SCI_GetParams
Description:    Get SCI Parameter
Calls:          HI_UNF_SCI_GetParams
Data Accessed:  NA
Data Updated:   NA
Input:          params
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_GetParams(params *HI_UNF_SCI_PARAMS_S) (bool, error) {
	if params.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	return sciCall(CMD_SCI_GET_PARAM, params)
}

/*************************************************************
Function:       HI_UNF_SCI_SetCharTimeout
Description:    Set T0 or T1 char timeout
Calls:          HI_UNF_SCI_SetCharTimeout
Data Accessed:  NA
Data Updated:   NA
Input:          timeout
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SetCharTimeout(timeout SCI_CHARTIMEOUT_S) (bool, error) {
	if timeout.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if timeout.Protocol == HI_UNF_SCI_PROTOCOL_T1 {
		if timeout.CharTimeouts < 12 || timeout.CharTimeouts > 32779 {
			return false, errors.New("The CharTimeouts for T1 is invalid.")
		}
	} else {
		if timeout.CharTimeouts < 960 || timeout.CharTimeouts > 244800 {
			return false, errors.New("The CharTimeouts is invalid.")
		}
	}

	return sciCall(CMD_SCI_SET_CHARTIMEOUT, &timeout)
}

/*************************************************************
Function:       HI_UNF_SCI_SetBlockTimeout
Description:    Set T1 block timeout
Calls:          HI_UNF_SCI_SetBlockTimeout
Data Accessed:  NA
Data Updated:   NA
Input:          timeout
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SetBlockTimeout(timeout SCI_BLOCKTIMEOUT_S) (bool, error) {
	if timeout.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if timeout.BlockTimeouts < 971 || timeout.BlockTimeouts > 491531 {
		return false, errors.New("The BlockTimeouts is invalid.")
	}

	return sciCall(CMD_SCI_SET_BLOCKTIMEOUT, &timeout)
}

/*************************************************************
Function:       HI_UNF_SCI_SetTxRetries
Description:    Set TX Retry Times
Calls:          HI_UNF_SCI_SetTxRetries
Data Accessed:  NA
Data Updated:   NA
Input:          tx
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_SCI_SetTxRetries(tx SCI_TXRETRY_S) (bool, error) {
	if tx.Port >= HI_UNF_SCI_PORT_BUTT {
		return false, errors.New("SCI Port is invalid.")
	}

	if tx.TxRetryTimes > 7 {
		return false, errors.New("The TxRetryTimes is invalid.")
	}

	return sciCall(CMD_SCI_SET_TXRETRY, &tx)
}
