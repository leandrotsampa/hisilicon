package hisilicon

import (
	"errors"
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
