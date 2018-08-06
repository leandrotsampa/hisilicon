package hisilicon

import "unsafe"

/* Ioctl Definitions */
const HI_ID_SCI uintptr = 0x54

var (
	CMD_SCI_OPEN                = IoW(HI_ID_SCI, 0x1, unsafe.Sizeof(SCI_OPEN_S{}))
	CMD_SCI_CLOSE               = IoW(HI_ID_SCI, 0x2, unsafe.Sizeof(HI_UNF_SCI_PORT_E(0)))
	CMD_SCI_RESET               = IoW(HI_ID_SCI, 0x3, unsafe.Sizeof(SCI_RESET_S{}))
	CMD_SCI_DEACTIVE            = IoW(HI_ID_SCI, 0x4, unsafe.Sizeof(HI_UNF_SCI_PORT_E(0)))
	CMD_SCI_GET_ATR             = IoRW(HI_ID_SCI, 0x5, unsafe.Sizeof(SCI_ATR_S{}))
	CMD_SCI_COMPAT_GET_ATR      = IoRW(HI_ID_SCI, 0x5, unsafe.Sizeof(SCI_ATR_COMPAT_S{}))
	CMD_SCI_GET_STATUS          = IoRW(HI_ID_SCI, 0x6, unsafe.Sizeof(SCI_STATUS_S{}))
	CMD_SCI_CONF_VCC            = IoW(HI_ID_SCI, 0x7, unsafe.Sizeof(SCI_LEVEL_S{}))
	CMD_SCI_CONF_DETECT         = IoW(HI_ID_SCI, 0x8, unsafe.Sizeof(SCI_LEVEL_S{}))
	CMD_SCI_CONF_MODE           = IoW(HI_ID_SCI, 0x9, unsafe.Sizeof(SCI_IO_OUTPUTTYPE_S{}))
	CMD_SCI_SEND_DATA           = IoRW(HI_ID_SCI, 0xa, unsafe.Sizeof(SCI_DATA_S{}))
	CMD_SCI_COMPAT_SEND_DATA    = IoRW(HI_ID_SCI, 0xa, unsafe.Sizeof(SCI_DATA_COMPAT_S{}))
	CMD_SCI_RECEIVE_DATA        = IoRW(HI_ID_SCI, 0xb, unsafe.Sizeof(SCI_DATA_S{}))
	CMD_SCI_COMPAT_RECEIVE_DATA = IoRW(HI_ID_SCI, 0xb, unsafe.Sizeof(SCI_DATA_COMPAT_S{}))
	CMD_SCI_SWITCH              = IoW(HI_ID_SCI, 0xc, unsafe.Sizeof(SCI_OPEN_S{}))
	CMD_SCI_SET_BAUD            = IoW(HI_ID_SCI, 0xd, unsafe.Sizeof(SCI_EXT_BAUD_S{}))
	CMD_SCI_SET_CHGUARD         = IoW(HI_ID_SCI, 0xe, unsafe.Sizeof(SCI_ADD_GUARD_S{}))
	CMD_SCI_SEND_PPS_DATA       = IoW(HI_ID_SCI, 0xF, unsafe.Sizeof(SCI_PPS_S{}))
	CMD_SCI_GET_PPS_DATA        = IoRW(HI_ID_SCI, 0x10, unsafe.Sizeof(SCI_PPS_S{}))
	CMD_SCI_GET_PARAM           = IoRW(HI_ID_SCI, 0x11, unsafe.Sizeof(HI_UNF_SCI_PARAMS_S{}))
	CMD_SCI_SET_CHARTIMEOUT     = IoW(HI_ID_SCI, 0x12, unsafe.Sizeof(SCI_CHARTIMEOUT_S{}))
	CMD_SCI_SET_BLOCKTIMEOUT    = IoW(HI_ID_SCI, 0x13, unsafe.Sizeof(SCI_BLOCKTIMEOUT_S{}))
	CMD_SCI_SET_TXRETRY         = IoW(HI_ID_SCI, 0x14, unsafe.Sizeof(SCI_TXRETRY_S{}))
)

/* Structs Definitions */
/** Output configuration of the smart card interface clock (SCICLK) pin **/
type HI_UNF_SCI_MODE_E int32

const (
	HI_UNF_SCI_MODE_CMOS HI_UNF_SCI_MODE_E = iota /* Complementary metal-oxide semiconductor (CMOS) output */
	HI_UNF_SCI_MODE_OD   HI_UNF_SCI_MODE_E = iota /* Open drain (OD) output */
	HI_UNF_SCI_MODE_BUTT HI_UNF_SCI_MODE_E = iota
)

/** SCI Port **/
type HI_UNF_SCI_PORT_E int32

const (
	HI_UNF_SCI_PORT0     HI_UNF_SCI_PORT_E = iota /* SCI Port 0 */
	HI_UNF_SCI_PORT1     HI_UNF_SCI_PORT_E = iota /* SCI Port 1 */
	HI_UNF_SCI_PORT_BUTT HI_UNF_SCI_PORT_E = iota
)

/** Status of the SCI Card **/
type HI_UNF_SCI_STATUS_E int32

const (
	HI_UNF_SCI_STATUS_UNINIT       HI_UNF_SCI_STATUS_E = iota /* The SCI Card is not initialized. (Reserved status) */
	HI_UNF_SCI_STATUS_FIRSTINIT    HI_UNF_SCI_STATUS_E = iota /* The SCI Card is being initialized.(Reserved status) */
	HI_UNF_SCI_STATUS_NOCARD       HI_UNF_SCI_STATUS_E = iota /* There is no SCI Card. */
	HI_UNF_SCI_STATUS_INACTIVECARD HI_UNF_SCI_STATUS_E = iota /* The SCI Card is not activated (unavailable). */
	//HI_UNF_SCI_STATUS_CARDFAULT    HI_UNF_SCI_STATUS_E = iota /* The SCI Card is faulty.*/
	HI_UNF_SCI_STATUS_WAITATR HI_UNF_SCI_STATUS_E = iota /* The SCI Card is waiting for the ATR data. */
	HI_UNF_SCI_STATUS_READATR HI_UNF_SCI_STATUS_E = iota /* The SCI Card is receiving the ATR data. */
	HI_UNF_SCI_STATUS_READY   HI_UNF_SCI_STATUS_E = iota /* The SCI Card is available (activated). */
	HI_UNF_SCI_STATUS_RX      HI_UNF_SCI_STATUS_E = iota /* The SCI Card is busy receiving data. */
	HI_UNF_SCI_STATUS_TX      HI_UNF_SCI_STATUS_E = iota /* The SCI Card is busy transmitting data. */
)

/** SCI Protocol **/
type HI_UNF_SCI_PROTOCOL_E int32

const (
	HI_UNF_SCI_PROTOCOL_T0   HI_UNF_SCI_PROTOCOL_E = iota /* 7816 T0 Protocol */
	HI_UNF_SCI_PROTOCOL_T1   HI_UNF_SCI_PROTOCOL_E = iota /* 7816 T1 Protocol */
	HI_UNF_SCI_PROTOCOL_T14  HI_UNF_SCI_PROTOCOL_E = iota /* 7816 T14 Protocol */
	HI_UNF_SCI_PROTOCOL_BUTT HI_UNF_SCI_PROTOCOL_E = iota
)

/** SCI Active Level **/
type HI_UNF_SCI_LEVEL_E int32

const (
	HI_UNF_SCI_LEVEL_LOW  HI_UNF_SCI_LEVEL_E = iota /* Active Low */
	HI_UNF_SCI_LEVEL_HIGH HI_UNF_SCI_LEVEL_E = iota /* Active High */
	HI_UNF_SCI_LEVEL_BUTT HI_UNF_SCI_LEVEL_E = iota
)

/** SCI System Parameters **/
type HI_UNF_SCI_PARAMS_S struct {
	Port          HI_UNF_SCI_PORT_E     /* SCI Port ID */
	Protocol      HI_UNF_SCI_PROTOCOL_E /* Used Protocol Type */
	ActalClkRate  HI_U32                /* Actual clock rate conversion factor F */
	ActalBitRate  HI_U32                /* Actual bit rate conversion factor D */
	Fi            HI_U32                /* Clock factor returned by the answer to reset (ATR) */
	Di            HI_U32                /* Bit rate factor returned by the ATR */
	GuardDelay    HI_U32                /* Extra Guard Time N */
	CharTimeouts  HI_U32                /* Character timeout of T0 or T1 */
	BlockTimeouts HI_U32                /* Block Timeout of T1 */
	TxRetries     HI_U32                /* Number of transmission retries */
}

type SCI_OPEN_S struct {
	Port      HI_UNF_SCI_PORT_E
	Protocol  HI_UNF_SCI_PROTOCOL_E
	Frequency HI_U32
}

type SCI_RESET_S struct {
	Port      HI_UNF_SCI_PORT_E
	WarmReset HI_BOOL
}

type SCI_ATR_S struct {
	Port       HI_UNF_SCI_PORT_E
	AtrBuf     *HI_U8
	AtrBufSize HI_U32
	DataLen    HI_U8
}

type SCI_ATR_COMPAT_S struct {
	Port       HI_UNF_SCI_PORT_E
	AtrBuf     HI_U32
	AtrBufSize HI_U32
	DataLen    HI_U8
}

type SCI_STATUS_S struct {
	Port   HI_UNF_SCI_PORT_E
	Status HI_UNF_SCI_STATUS_E
}

type SCI_DATA_S struct {
	Port      HI_UNF_SCI_PORT_E
	DataBuf   *HI_U8
	BufSize   HI_U32
	DataLen   HI_U32
	TimeoutMs HI_U32
}

type SCI_DATA_COMPAT_S struct {
	Port      HI_UNF_SCI_PORT_E
	DataBuf   HI_U32
	BufSize   HI_U32
	DataLen   HI_U32
	TimeoutMs HI_U32
}

type SCI_LEVEL_S struct {
	Port  HI_UNF_SCI_PORT_E
	Level HI_UNF_SCI_LEVEL_E
}

type SCI_IO_E int32

const (
	SCI_IO_CLK    SCI_IO_E = iota
	SCI_IO_RESET  SCI_IO_E = iota
	SCI_IO_VCC_EN SCI_IO_E = iota
	SCI_IO_BUTT   SCI_IO_E = iota
)

type SCI_IO_OUTPUTTYPE_S struct {
	Port       HI_UNF_SCI_PORT_E
	IO         SCI_IO_E
	OutputType HI_UNF_SCI_MODE_E
}

type SCI_DEV_STATE_S struct {
	bSci [HI_UNF_SCI_PORT_BUTT]HI_BOOL
}

type SCI_EXT_BAUD_S struct {
	Port    HI_UNF_SCI_PORT_E
	ClkRate HI_U32
	BitRate HI_U32
}

type SCI_ADD_GUARD_S struct {
	Port         HI_UNF_SCI_PORT_E
	AddCharGuard HI_U32
}

type SCI_PPS_S struct {
	Port        HI_UNF_SCI_PORT_E
	Send        [6]HI_U8
	Receive     [6]HI_U8
	SendLen     HI_U32
	ReceiveLen  HI_U32
	RecTimeouts HI_U32
}

type SCI_CHARTIMEOUT_S struct {
	Port         HI_UNF_SCI_PORT_E
	Protocol     HI_UNF_SCI_PROTOCOL_E
	CharTimeouts HI_U32
}

type SCI_BLOCKTIMEOUT_S struct {
	Port          HI_UNF_SCI_PORT_E
	BlockTimeouts HI_U32
}

type SCI_TXRETRY_S struct {
	Port         HI_UNF_SCI_PORT_E
	TxRetryTimes HI_U32
}
