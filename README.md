# Hisilicon API for Golang
This code is based on source for SoC Hi3798Cv200.

**_NOTE:_** The compatibility with old and new SoC's is possible.

## Example

```go
package main

import (
	"fmt"

	hiapi "github.com/leandrotsampa/hisilicon"
)

func main() {
	if _, err := hiapi.HI_UNF_IR_Init(); err != nil {
		panic(err)
	}
	if _, err := hiapi.HI_UNF_IR_EnableRepKey(true); err != nil {
		panic(err)
	}

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
		}
	}

	hiapi.HI_UNF_IR_DeInit()
}
```

## Installing

### Using *go get*

    $ go get github.com/leandrotsampa/hisilicon

After this command *hisilicon* is ready to use. Its source will be in:

    $GOPATH/src/pkg/github.com/leandrotsampa/hisilicon

You can use `go get -u` to update the package.

### Hisilicon API Porting Status

- [ ] ADEC
- [ ] ADVCA
- [ ] AENC
- [ ] AI
- [ ] AO
- [ ] AVPLAY
- [ ] CI
- [ ] CIPHER
- [ ] DEMUX
- [ ] FRONTEND
- [ ] GFX2D
- [ ] GPIO
- [ ] GPU
- [ ] HDCP
- [ ] HDMI
- [ ] HDMIRX
- [ ] HIGO
- [ ] I2C
- [X] IR
- [ ] JPEG
- [ ] JPGE
- [ ] KEYLED
- [ ] MCE
- [ ] OMX
- [ ] OTP
- [ ] PDM
- [ ] PM
- [ ] PNG
- [ ] PQ
- [ ] PVR
- [ ] PWM
- [ ] SCI
- [ ] SPI
- [ ] SYNC
- [ ] TDE
- [ ] VDEC
- [ ] VENC
- [ ] VI
- [ ] VO
- [ ] WDG