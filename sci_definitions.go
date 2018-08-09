package hisilicon

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
