package hisilicon

import (
	"errors"
	"os"
)

var module = HiDevice{fd: nil}

/** Internal Function for Module IOCTL Calls **/
func moduleCall(op uintptr, arg interface{}) (bool, error) {
	module.mu.Lock()
	defer module.mu.Unlock()

	if module.fd == nil {
		return false, errors.New("Module Device not initialized.")
	}

	if err := Ioctl(module.fd.Fd(), op, arg); err != nil {
		return false, err
	}

	return true, nil
}

func HI_MODULE_Init() (bool, error) {
	module.mu.Lock()
	defer module.mu.Unlock()

	if module.fd != nil {
		module.InUse++
		return true, nil
	}

	var err error
	if module.fd, err = os.OpenFile("/dev/hi_module", os.O_RDWR, 0); err != nil {
		return false, err
	}

	return true, nil
}

func HI_MODULE_DeInit() (bool, error) {
	module.mu.Lock()
	defer module.mu.Unlock()

	if module.fd == nil {
		return true, nil
	} else if module.InUse > 0 {
		module.InUse--
		return true, nil
	}

	if err := module.fd.Close(); err != nil {
		return false, err
	}

	return true, nil
}

func HI_MODULE_Register(u32ModuleID HI_U32, pszModuleName string) (bool, error) {
	if len(pszModuleName) == 0 {
		return false, errors.New("Module name is invalid")
	}

	stModule := MODULE_INFO_S{}
	copy(stModule.u8ModuleName[:], []byte(pszModuleName))
	stModule.u32ModuleID = u32ModuleID

	if _, err := moduleCall(CMD_ADD_MODULE_INFO, &stModule); err != nil {
		return false, err
	}

	return true, nil
}

func HI_MODULE_RegisterByName(pszModuleName string, pu32ModuleID *HI_U32) (bool, error) {
	if len(pszModuleName) == 0 || len(pszModuleName) > MAX_MODULE_NAME-1 {
		return false, errors.New("Module name is invalid")
	} else if pu32ModuleID == nil {
		return false, errors.New("ModuleID is invalid")
	}

	stModule := MODULE_ALLOC_S{}
	copy(stModule.u8ModuleName[:], []byte(pszModuleName))

	if _, err := moduleCall(CMD_ALLOC_MODULE_ID, &stModule); err == nil {
		if stModule.s32Status == 1 {
			*pu32ModuleID = stModule.u32ModuleID
			return true, nil
		} else if stModule.s32Status == 0 {
			if _, err := HI_MODULE_Register(stModule.u32ModuleID, pszModuleName); err == nil {
				*pu32ModuleID = stModule.u32ModuleID
				return true, nil
			}
		}
	}

	return false, errors.New("Can't register module.")
}

func HI_MODULE_UnRegister(u32ModuleID HI_U32) (bool, error) {
	stModule := MODULE_INFO_S{}
	stModule.u32ModuleID = u32ModuleID

	if _, err := moduleCall(CMD_GET_MODULE_INFO, &stModule); err != nil {
		return false, err
	}

	if _, err := moduleCall(CMD_DEL_MODULE_INFO, &stModule); err != nil {
		return false, err
	}

	return true, nil
}

func HI_MODULE_GetModuleID(pu8ModuleName string) (HI_U32, error) {
	stModule := MODULE_INFO_S{}
	copy(stModule.u8ModuleName[:], []byte(pu8ModuleName))

	if _, err := moduleCall(CMD_GET_MODULE_INFO, &stModule); err != nil {
		return 0, err
	}

	return stModule.u32ModuleID, nil
}

func HI_MODULE_GetModuleName(u32ModuleID HI_U32) (string, error) {
	stModule := MODULE_INFO_S{}
	stModule.u32ModuleID = u32ModuleID

	if _, err := moduleCall(CMD_GET_MODULE_INFO, &stModule); err != nil {
		return "", err
	}

	return string(stModule.u8ModuleName[:]), nil
}
