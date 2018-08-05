package hisilicon

import "unsafe"

/* Ioctl Definitions */
const HI_ID_SCI uintptr = 0x54

var (
	CMD_SCI_OPEN = IoW(HI_ID_SCI, 0x1, unsafe.Sizeof(SCI_OPEN_S{}))
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

type SCI_OPEN_S struct {
	Port      HI_UNF_SCI_PORT_E
	Protocol  HI_UNF_SCI_PROTOCOL_E
	Frequency HI_U32
}
