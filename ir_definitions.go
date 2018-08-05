package hisilicon

import (
	"unsafe"
)

/* Ioctl Definitions */
const HI_ID_IR uintptr = 0x51
const PROTOCOL_NAME_SZ int32 = 32

var (
	/* 1:check keyup */
	CMD_IR_ENABLE_KEYUP = IoW(HI_ID_IR, 0x1, unsafe.Sizeof(int32(0)))

	/* 1:check repkey, 0:hardware behave */
	CMD_IR_ENABLE_REPKEY      = IoW(HI_ID_IR, 0x2, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_REPKEY_TIMEOUT = IoW(HI_ID_IR, 0x3, unsafe.Sizeof(int32(0)))

	/* 1:enable ir, 0:disable ir */
	CMD_IR_SET_ENABLE    = IoW(HI_ID_IR, 0x4, unsafe.Sizeof(int32(0)))
	CMD_IR_RESET         = Io(HI_ID_IR, 0x5)
	CMD_IR_SET_BLOCKTIME = IoW(HI_ID_IR, 0x6, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_FORMAT    = IoW(HI_ID_IR, 0x7, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_BUF       = IoW(HI_ID_IR, 0x8, unsafe.Sizeof(int32(0)))

	/* raw symbol fetch(1) or key fetch(0) */
	CMD_IR_SET_FETCH_METHOD = IoW(HI_ID_IR, 0x9, unsafe.Sizeof(int32(0)))

	/* enable or disalbe a protocol */
	CMD_IR_SET_PROT_ENABLE  = IoW(HI_ID_IR, 0xa, unsafe.Sizeof(int32(0)))
	CMD_IR_SET_PROT_DISABLE = IoW(HI_ID_IR, 0xb, unsafe.Sizeof(int32(0)))
	CMD_IR_GET_PROT_ENABLED = IoRW(HI_ID_IR, 0xc, unsafe.Sizeof(int32(0)))
)

/* Structs Definitions */
type HI_UNF_KEY_STATUS_E int32

const (
	HI_UNF_KEY_STATUS_DOWN HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_HOLD HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_UP   HI_UNF_KEY_STATUS_E = iota
	HI_UNF_KEY_STATUS_BUTT HI_UNF_KEY_STATUS_E = iota
)

type HI_UNF_IR_CODE_E int32

const (
	HI_UNF_IR_CODE_NEC_SIMPLE HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_TC9012     HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_NEC_FULL   HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_SONY_12BIT HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_RAW        HI_UNF_IR_CODE_E = iota
	HI_UNF_IR_CODE_BUTT       HI_UNF_IR_CODE_E = iota
)

type HI_UNF_IR_PROTOCOL_E int32

const (
	HI_UNF_IR_NEC  HI_UNF_IR_PROTOCOL_E = 0
	HI_UNF_IR_RC6A HI_UNF_IR_PROTOCOL_E = iota + 10
	HI_UNF_IR_RC5
	HI_UNF_IR_LOW_LATENCY_PROTOCOL
	HI_UNF_IR_RC6_MODE0
	HI_UNF_IR_RCMM
	HI_UNF_IR_RUWIDO
	HI_UNF_IR_RCRF8
	HI_UNF_IR_MULTIPLE
	HI_UNF_IR_RMAP
	HI_UNF_IR_RSTEP
	HI_UNF_IR_RMAP_DOUBLEBIT
	HI_UNF_IR_LOW_LATENCY_PRO_PROTOCOL
	HI_UNF_IR_XMP
	HI_UNF_IR_USER_DEFINED
	HI_UNF_IR_PROTOCOL_BUTT
)

type KeyAttr struct {
	/* upper 16bit data under key mode */
	Upper uint64
	/* lower 32bit data under key mode
	 * or symbol value under symbol mode
	 */
	Lower        uint64
	ProtocolName [PROTOCOL_NAME_SZ]byte
	/* indentify key status. */
	StatusKey HI_UNF_KEY_STATUS_E
}
