package hisilicon

/* Structs Definitions */
/* Defines the capability of the DEMUX module */
type HI_UNF_DMX_CAPABILITY_S struct {
	u32IFPortNum    HI_U32 /* Number of IF ports. */
	u32TSIPortNum   HI_U32 /* Number of TS In ports. */
	u32TSOPortNum   HI_U32 /* Number of TS Out ports. */
	u32RamPortNum   HI_U32 /* Number of Ram ports. */
	u32DmxNum       HI_U32 /* Number of DEMUX devices */
	u32ChannelNum   HI_U32 /* Number of channels, containing the audio and video channels */
	u32AVChannelNum HI_U32 /* Number of av channels */
	u32FilterNum    HI_U32 /* Number of filters */
	u32KeyNum       HI_U32 /* Number of keys */
	u32RecChnNum    HI_U32 /* Number of record channels */
	u32TagPortNum   HI_U32 /* Number of Tag ports. */
}

/* TSO port ID */
type HI_UNF_DMX_TSO_PORT_E int32

const (
	HI_UNF_DMX_PORT_TSO_0 HI_UNF_DMX_TSO_PORT_E = iota /* The first TS OUT port. */
	HI_UNF_DMX_PORT_TSO_1
	HI_UNF_DMX_PORT_TSO_2
	HI_UNF_DMX_PORT_TSO_3
)

/* TS port ID */
type HI_UNF_DMX_PORT_E int32

const (
	HI_UNF_DMX_PORT_IF_0 HI_UNF_DMX_PORT_E = iota /* The first IF port (port with QAM inside chipset). */
	HI_UNF_DMX_PORT_IF_1
	HI_UNF_DMX_PORT_IF_2
	HI_UNF_DMX_PORT_IF_3
	HI_UNF_DMX_PORT_IF_4
	HI_UNF_DMX_PORT_IF_5
	HI_UNF_DMX_PORT_IF_6
	HI_UNF_DMX_PORT_IF_7
	HI_UNF_DMX_PORT_IF_8
	HI_UNF_DMX_PORT_IF_9
	HI_UNF_DMX_PORT_IF_10
	HI_UNF_DMX_PORT_IF_11
	HI_UNF_DMX_PORT_IF_12
	HI_UNF_DMX_PORT_IF_13
	HI_UNF_DMX_PORT_IF_14
	HI_UNF_DMX_PORT_IF_15
)

const (
	HI_UNF_DMX_PORT_TSI_0 HI_UNF_DMX_PORT_E = iota + 0x20 /* The first TS IN port. */
	HI_UNF_DMX_PORT_TSI_1
	HI_UNF_DMX_PORT_TSI_2
	HI_UNF_DMX_PORT_TSI_3
	HI_UNF_DMX_PORT_TSI_4
	HI_UNF_DMX_PORT_TSI_5
	HI_UNF_DMX_PORT_TSI_6
	HI_UNF_DMX_PORT_TSI_7
	HI_UNF_DMX_PORT_TSI_8
	HI_UNF_DMX_PORT_TSI_9
	HI_UNF_DMX_PORT_TSI_10
	HI_UNF_DMX_PORT_TSI_11
	HI_UNF_DMX_PORT_TSI_12
	HI_UNF_DMX_PORT_TSI_13
	HI_UNF_DMX_PORT_TSI_14
	HI_UNF_DMX_PORT_TSI_15
)

const (
	HI_UNF_DMX_PORT_RAM_0 HI_UNF_DMX_PORT_E = iota + 0x80 /* The first RAM port. */
	HI_UNF_DMX_PORT_RAM_1
	HI_UNF_DMX_PORT_RAM_2
	HI_UNF_DMX_PORT_RAM_3
	HI_UNF_DMX_PORT_RAM_4
	HI_UNF_DMX_PORT_RAM_5
	HI_UNF_DMX_PORT_RAM_6
	HI_UNF_DMX_PORT_RAM_7
	HI_UNF_DMX_PORT_RAM_8
	HI_UNF_DMX_PORT_RAM_9
	HI_UNF_DMX_PORT_RAM_10
	HI_UNF_DMX_PORT_RAM_11
	HI_UNF_DMX_PORT_RAM_12
	HI_UNF_DMX_PORT_RAM_13
	HI_UNF_DMX_PORT_RAM_14
	HI_UNF_DMX_PORT_RAM_15

	HI_UNF_DMX_PORT_BUTT
)

/* TS port mode */
type HI_UNF_DMX_PORT_MODE_E int32

const (
	HI_UNF_DMX_PORT_MODE_EXTERNAL HI_UNF_DMX_PORT_MODE_E = iota /* External TS input mode */
	HI_UNF_DMX_PORT_MODE_INTERNAL                               /* Internal TS input mode */
	HI_UNF_DMX_PORT_MODE_RAM                                    /* Memory input mode */
	HI_UNF_DMX_PORT_MODE_BUTT
)

/* TS port type */
type HI_UNF_DMX_PORT_TYPE_E int32

const (
	HI_UNF_DMX_PORT_TYPE_PARALLEL_BURST          HI_UNF_DMX_PORT_TYPE_E = iota /* Parallel burst mode */
	HI_UNF_DMX_PORT_TYPE_PARALLEL_VALID                                        /* Parallel valid mode */
	HI_UNF_DMX_PORT_TYPE_PARALLEL_NOSYNC_188                                   /* Self-sync 188 mode */
	HI_UNF_DMX_PORT_TYPE_PARALLEL_NOSYNC_204                                   /* Self-sync 204 mode */
	HI_UNF_DMX_PORT_TYPE_PARALLEL_NOSYNC_188_204                               /* self-sync 188/204 auto-identification mode */
	HI_UNF_DMX_PORT_TYPE_SERIAL                                                /* Serial sync mode, 1bit */
	HI_UNF_DMX_PORT_TYPE_USER_DEFINED                                          /* User defined mode */
	HI_UNF_DMX_PORT_TYPE_SERIAL2BIT                                            /* Serial sync mode, 2bit */
	HI_UNF_DMX_PORT_TYPE_SERIAL_NOSYNC                                         /* Serial nosync mode, 1bit */
	HI_UNF_DMX_PORT_TYPE_SERIAL2BIT_NOSYNC                                     /* Serial nosync mode, 2bit */
	HI_UNF_DMX_PORT_TYPE_AUTO                                                  /* Auto mode */
	HI_UNF_DMX_PORT_TYPE_BUTT
)

/* TSO clock mode */
type HI_UNF_DMX_TSO_CLK_MODE_E int32

const (
	HI_UNF_DMX_TSO_CLK_MODE_NORMAL HI_UNF_DMX_TSO_CLK_MODE_E = iota /* Normal clock: clock always active */
	HI_UNF_DMX_TSO_CLK_MODE_JITTER                                  /* Jittered clock: clock active only when outputing data */
	HI_UNF_DMX_TSO_CLK_MODE_BUTT
)

/* TSO valid mode */
type HI_UNF_DMX_TSO_VALID_MODE_E int32

const (
	HI_UNF_DMX_TSO_VALID_ACTIVE_OUTPUT HI_UNF_DMX_TSO_VALID_MODE_E = iota /* Valid signal high when outputing datas */
	HI_UNF_DMX_TSO_VALID_ACTIVE_HIGH                                      /* Valid signal always high */
	HI_UNF_DMX_TSO_VALID_ACTIVE_BUTT
)

/* TSO port signal line selector */
type HI_UNF_DMX_TSO_SERIAL_BIT_E int32

const (
	HI_UNF_DMX_TSO_SERIAL_BIT_0    HI_UNF_DMX_TSO_SERIAL_BIT_E = 0x0 /* Serial output data using data[0] as signal line */
	HI_UNF_DMX_TSO_SERIAL_BIT_7                                = 0x7 /* Serial output data using data[7] as signal line */
	HI_UNF_DMX_TSO_SERIAL_BIT_BUTT                             = 0x8
)

/* TS out mode clock frequency */
type HI_UNF_DMX_TSO_CLK_E int32

const (
	HI_UNF_DMX_TSO_CLK_100M  HI_UNF_DMX_TSO_CLK_E = iota /* TS out mode clock frequency 100M */
	HI_UNF_DMX_TSO_CLK_150M                              /* TS out mode clock frequency 150M */
	HI_UNF_DMX_TSO_CLK_1200M                             /* TS out mode clock frequency 1200M */
	HI_UNF_DMX_TSO_CLK_1500M                             /* TS out mode clock frequency 1500M */
	HI_UNF_DMX_TSO_CLK_BUTT
)

/* Tag sync mode */
type HI_UNF_DMX_TAG_SYNC_MODE_E int32

const (
	HI_UNF_DMX_TAG_HEAD_SYNC    HI_UNF_DMX_TAG_SYNC_MODE_E = 0x0 /* tag sync signal at tag head */
	HI_UNF_DMX_NORMAL_HEAD_SYNC                            = 0x1 /* tag sync signal at 47 header */
)

/* TS Tag attributes */
const MAX_TAG_LENGTH = 12

type HI_UNF_DMX_TAG_ATTR_S struct {
	au8Tag    [MAX_TAG_LENGTH]HI_U8      /* [IN]tag index value */
	u32TagLen HI_U32                     /* [IN & OUT]Valid tag length(1 ~ 12bytes), which is an input para when setTagAttr, otherwise out para when getTagAttr */
	bEnabled  HI_BOOL                    /* [IN & OUT]Port state(default disabled), which is an input para when setTagAttr, otherwise out para when getTagAttr */
	enSyncMod HI_UNF_DMX_TAG_SYNC_MODE_E /* [IN & OUT]Sync mode(default HI_UNF_DMX_TAG_HEAD_SYNC), which is an input para when setTagAttr, otherwise out para when getTagAttr */
}

/* TS out port attributes */
type HI_UNF_DMX_TSO_PORT_ATTR_S struct {
	bEnable       HI_BOOL                     /* Port enable, default value HI_TRUE means enable */
	bClkReverse   HI_BOOL                     /* Clock phase reverse, default value HI_FALSE means do not reverse the phase of clock */
	enTSSource    HI_UNF_DMX_PORT_E           /* Source of this TS Out port ,can choose from HI_UNF_DMX_PORT_IF_0 to HI_UNF_DMX_PORT_TSI_9 */
	enClkMode     HI_UNF_DMX_TSO_CLK_MODE_E   /* Clock mode: HI_UNF_DMX_TSO_CLK_MODE_NORMAL is the default value */
	enValidMode   HI_UNF_DMX_TSO_VALID_MODE_E /* Wether valid signal always enable : HI_UNF_DMX_TSO_VALID_ACTIVE_OUTPUT is the default value */
	bBitSync      HI_BOOL                     /* The sync signal duration : HI_TRUE: only valid when output the first bit(default). HI_FALSE: keep when outputing the whole byte */
	bSerial       HI_BOOL                     /* Wether out put mode is serial: HI_FALSE: parallel mode. HI_TRUE: serial mode (default) */
	enBitSelector HI_UNF_DMX_TSO_SERIAL_BIT_E /* Port line sequence select In serial mode.only valid when using serial out put mode,HI_UNF_DMX_TSO_SERIAL_BIT_7 is the default value. */
	bLSB          HI_BOOL                     /* Out put byte endian .only valid when using serial out put mode: HI_FALSE: first output MSB (default). HI_TRUE:  first output LSB */
	enClk         HI_UNF_DMX_TSO_CLK_E        /* TS out mode clock frequency,default is HI_UNF_DMX_TSO_CLK_150M */
	u32ClkDiv     HI_U32                      /* TS out mode clock frequency divider,must be times of 2 ,and must meet (2 <= u32ClkDiv <= 32) .default is 2 */
}

/* TS port attributes */
type HI_UNF_DMX_PORT_ATTR_S struct {
	enPortMod            HI_UNF_DMX_PORT_MODE_E /* Port mode.Readonly */
	enPortType           HI_UNF_DMX_PORT_TYPE_E /* Port type */
	u32SyncLostTh        HI_U32                 /* Sync loss threshold.The default value is recommended. */
	u32SyncLockTh        HI_U32                 /* Sync lock threshold.The default value is recommended. */
	u32TunerInClk        HI_U32                 /* Whether to reverse the phase of the clock input from the tuner */
	u32SerialBitSelector HI_U32                 /* Port line sequence select In parallel mode: 0: cdata[7] is the most significant bit (MSB) (default). 1: cdata[0] is the MSB.
	   In serial mode: 1: cdata[0] is the data line (default). 0: cdata[7] is the data line. */
	u32TunerErrMod HI_U32 /* Level mode of the cerr_n line from the tuner to a DEMUX.
	   0: A data error occurs when the cerr_n line is high.
	   1: A data error occurs when the cerr_n line is low (default). */
	u32UserDefLen1 HI_U32 /* User defined length1,valid when enPortType is HI_UNF_DMX_PORT_TYPE_USER_DEFINED,188~255 */
	u32UserDefLen2 HI_U32 /* User defined length2,valid when enPortType is HI_UNF_DMX_PORT_TYPE_USER_DEFINED,188~255 */
}

/* Status of the TS port */
type HI_UNF_DMX_PORT_PACKETNUM_S struct {
	u32TsPackCnt    HI_U32 /* Number of TS packets received from the TS port */
	u32ErrTsPackCnt HI_U32 /* Number of error TS packets received from the TS port */
}

/* Status of a TS buffer of a DEMUX */
type HI_UNF_DMX_TSBUF_STATUS_S struct {
	u32BufSize  HI_U32 /* Buffer size */
	u32UsedSize HI_U32 /* Used buffer size */
}

/* Channel type */
type HI_UNF_DMX_CHAN_TYPE_E int32

const (
	HI_UNF_DMX_CHAN_TYPE_SEC     HI_UNF_DMX_CHAN_TYPE_E = iota /* Channel that receives sections data such as program specific information (PSI) or service information (SI) data */
	HI_UNF_DMX_CHAN_TYPE_PES                                   /* Channel that receives packetized elementary stream (PES) data */
	HI_UNF_DMX_CHAN_TYPE_AUD                                   /* Channel that receives audio data */
	HI_UNF_DMX_CHAN_TYPE_VID                                   /* Channel that receives video data */
	HI_UNF_DMX_CHAN_TYPE_POST                                  /* Entire-packet posting channel that receives an entire TS packet with a specific packet identifier (PID). */
	HI_UNF_DMX_CHAN_TYPE_ECM_EMM                               /* Channel that receives entitlement control message (ECM) or entitlement management message (EMM) data */
	HI_UNF_DMX_CHAN_TYPE_BUTT
)

/* Cyclic redundancy check (CRC) mode of a channel */
type HI_UNF_DMX_CHAN_CRC_MODE_E int32

const (
	HI_UNF_DMX_CHAN_CRC_MODE_FORBID                HI_UNF_DMX_CHAN_CRC_MODE_E = iota /* The CRC check is disabled */
	HI_UNF_DMX_CHAN_CRC_MODE_FORCE_AND_DISCARD                                       /* The CRC check is enabled, and the error Section data is discarded */
	HI_UNF_DMX_CHAN_CRC_MODE_FORCE_AND_SEND                                          /* The CRC check is enabled, and the error Section data is received */
	HI_UNF_DMX_CHAN_CRC_MODE_BY_SYNTAX_AND_DISCARD                                   /* Whether the CRC check is performed depends on the syntax, and the error Section data is discarded */
	HI_UNF_DMX_CHAN_CRC_MODE_BY_SYNTAX_AND_SEND                                      /* Whether the CRC check is performed depends on the syntax, and the error Section data is received */

	HI_UNF_DMX_CHAN_CRC_MODE_BUTT
)

/* Output mode of a channel */
type HI_UNF_DMX_CHAN_OUTPUT_MODE_E int32

const (
	HI_UNF_DMX_CHAN_OUTPUT_MODE_PLAY     HI_UNF_DMX_CHAN_OUTPUT_MODE_E = iota + 0x1 /* Mode of playing audios/videos or receiving data */
	HI_UNF_DMX_CHAN_OUTPUT_MODE_REC                                                 /* Recording mode */
	HI_UNF_DMX_CHAN_OUTPUT_MODE_PLAY_REC                                            /* Mode of recording and playing data or receiving data */
	HI_UNF_DMX_CHAN_OUTPUT_MODE_BUTT
)

/* Secure mode type */
type HI_UNF_DMX_SECURE_MODE_E int32

const (
	HI_UNF_DMX_SECURE_MODE_NONE HI_UNF_DMX_SECURE_MODE_E = iota /* no security protection */
	HI_UNF_DMX_SECURE_MODE_TEE                                  /* trustedzone security protection */
	HI_UNF_DMX_SECURE_MODE_BUTT
)

/* Channel attribute */
type HI_UNF_DMX_CHAN_ATTR_S struct {
	u32BufSize    HI_U32                        /* Buffer size used by channels */
	enChannelType HI_UNF_DMX_CHAN_TYPE_E        /* Channel type */
	enCRCMode     HI_UNF_DMX_CHAN_CRC_MODE_E    /* CRC mode.It is valid for the DEMUX_CHAN_SEC channel. */
	enOutputMode  HI_UNF_DMX_CHAN_OUTPUT_MODE_E /* Output mode of the channel data */
	enSecureMode  HI_UNF_DMX_SECURE_MODE_E      /* Secure channel indication */
}

/* Scrambled flag of the channel data */
type HI_UNF_DMX_SCRAMBLED_FLAG_E int32

const (
	HI_UNF_DMX_SCRAMBLED_FLAG_TS  HI_UNF_DMX_SCRAMBLED_FLAG_E = iota /* TS data is scrambled */
	HI_UNF_DMX_SCRAMBLED_FLAG_PES                                    /* PES data is scrambled */
	HI_UNF_DMX_SCRAMBLED_FLAG_NO                                     /* Data is not scrambled */
	HI_UNF_DMX_SCRAMBLED_FLAG_BUTT
)

/* Channel status */
type HI_UNF_DMX_CHAN_STATUS_E int32

const (
	HI_UNF_DMX_CHAN_CLOSE       HI_UNF_DMX_CHAN_STATUS_E = iota + 0x0 /* The channel is stopped. */
	HI_UNF_DMX_CHAN_PLAY_EN                                           /* The channel is playing audios/videos or receiving data. */
	HI_UNF_DMX_CHAN_REC_EN                                            /* The channel is recording data. */
	HI_UNF_DMX_CHAN_PLAY_REC_EN                                       /* The channel is recording and receiving data. */
)

/* Defines the channel status */
type HI_UNF_DMX_CHAN_STATUS_S struct {
	enChanStatus HI_UNF_DMX_CHAN_STATUS_E /* Channel status */
}

/* Filter attribute */
const DMX_FILTER_MAX_DEPTH = 16

type HI_UNF_DMX_FILTER_ATTR_S struct {
	u32FilterDepth HI_U32                      /* Depth of a filter. */
	au8Match       [DMX_FILTER_MAX_DEPTH]HI_U8 /* Matched bytes of a filter.The data is compared by bit. */
	au8Mask        [DMX_FILTER_MAX_DEPTH]HI_U8 /* Masked bytes of a filter. The conditions are set by bit. 0: no mask. Comparison is required. 1: mask. Comparison is not required. */
	au8Negate      [DMX_FILTER_MAX_DEPTH]HI_U8 /* Negated bytes of a filter. 0: not negated; 1: negated */
}

/* Type of the DEMUX data packet. */
type HI_UNF_DMX_DATA_TYPE_E int32

const (
	HI_UNF_DMX_DATA_TYPE_WHOLE HI_UNF_DMX_DATA_TYPE_E = iota /* The data segment contains a complete data packet */
	HI_UNF_DMX_DATA_TYPE_HEAD                                /* The data segment contains the head of a data packet, but the data packet may not be complete */
	HI_UNF_DMX_DATA_TYPE_BODY                                /* This type is valid only for the PES data.The data segment contains the body of a data packet. */
	HI_UNF_DMX_DATA_TYPE_TAIL                                /* This type is valid only for the PES data.The data segment contains the tail of a data packet, and is used to identify the end of a data packet. */
	HI_UNF_DMX_DATA_TYPE_BUTT
)

/* DEMUX data packet */
type HI_UNF_DMX_DATA_S struct {
	pu8Data    *HI_U8                 /* Data pointer */
	u32Size    HI_U32                 /* Data length */
	enDataType HI_UNF_DMX_DATA_TYPE_E /* Data packet type */
}

/* type of record */
type HI_UNF_DMX_REC_TYPE_E int32

const (
	HI_UNF_DMX_REC_TYPE_SELECT_PID HI_UNF_DMX_REC_TYPE_E = iota
	HI_UNF_DMX_REC_TYPE_ALL_PID
	HI_UNF_DMX_REC_TYPE_BUTT
)

/**type of index*/
type HI_UNF_DMX_REC_INDEX_TYPE_E int32

const (
	HI_UNF_DMX_REC_INDEX_TYPE_NONE  HI_UNF_DMX_REC_INDEX_TYPE_E = iota /* No index is created */
	HI_UNF_DMX_REC_INDEX_TYPE_VIDEO                                    /* Video index */
	HI_UNF_DMX_REC_INDEX_TYPE_AUDIO                                    /* Audio index */
	HI_UNF_DMX_REC_INDEX_TYPE_BUTT
)

/* record attribute */
type HI_UNF_DMX_REC_ATTR_S struct {
	u32DmxId       HI_U32
	u32RecBufSize  HI_U32                      /* Buffer size used by record  */
	enRecType      HI_UNF_DMX_REC_TYPE_E       /* Record type  */
	bDescramed     HI_BOOL                     /* HI_TRUE is the descrambled TS. HI_FALSE is the original TS. */
	enIndexType    HI_UNF_DMX_REC_INDEX_TYPE_E /* Index type */
	u32IndexSrcPid HI_U32                      /* The index information is formed according to the PID. when indexing video, it has to be set to the video of PID. when indexing audio, it has to be set to the audio of PID. */
	enVCodecType   HI_UNF_VCODEC_TYPE_E        /* Video encoding protocol. The protocol needs to be set only when the index type is HI_UNF_DMX_REC_INDEX_TYPE_VIDEO. */
	enSecureMode   HI_UNF_DMX_SECURE_MODE_E    /* Secure record indication */
}

/* record data */
type HI_UNF_DMX_REC_DATA_S struct {
	pDataAddr      *HI_U8 /* Data address */
	u32DataPhyAddr HI_U32 /* Data physical address */
	u32Len         HI_U32 /* Data length */
}

/* index data */
type HI_UNF_DMX_REC_INDEX_S struct {
	enFrameType     HI_UNF_VIDEO_FRAME_TYPE_E /* it is meaningless when indexing audio. */
	u32PtsMs        HI_U32
	u64GlobalOffset HI_U64
	u32FrameSize    HI_U32 /* it is meaningless when indexing audio. */
	u32DataTimeMs   HI_U32
}

/* index and record data */
const DMX_MAX_IDX_ACQUIRED_EACH_TIME = 256

type HI_UNF_DMX_REC_DATA_INDEX_S struct {
	u32IdxNum     HI_U32 /* Number of index */
	u32RecDataCnt HI_U32 /* Number of record data block */
	stIndex       [DMX_MAX_IDX_ACQUIRED_EACH_TIME]HI_UNF_DMX_REC_INDEX_S
	stRecData     [2]HI_UNF_DMX_REC_DATA_S
}

/* record buffer status */
type HI_UNF_DMX_RECBUF_STATUS_S struct {
	u32BufSize  HI_U32 /* Buffer size */
	u32UsedSize HI_U32 /* Used buffer */
}

/* Repeat CC mode of channel */
type HI_UNF_DMX_CHAN_CC_REPEAT_MODE_E int32

const (
	HI_UNF_DMX_CHAN_CC_REPEAT_MODE_RSV  HI_UNF_DMX_CHAN_CC_REPEAT_MODE_E = 0x0 /* Receive CC repeat ts packet */
	HI_UNF_DMX_CHAN_CC_REPEAT_MODE_DROP                                  = 0x1 /* Drop CC repeat ts packet */
	HI_UNF_DMX_CHAN_CC_REPEAT_MODE_BUTT                                  = -1
)

type HI_UNF_DMX_CHAN_CC_REPEAT_SET_S struct {
	hChannel       HI_HANDLE                        /* The channel handle */
	enCCRepeatMode HI_UNF_DMX_CHAN_CC_REPEAT_MODE_E /* Repeat CC mode of channel */
}

/* PUSI (Payload Unit Start Index) config structure */
type HI_UNF_DMX_PUSI_SET_S struct {
	bPusi HI_BOOL /* Value of Pusi , Default is HI_FALSE means receive ts packet without checking PUSI */
}

/* TEI (Transport Error Index) config structure */
type HI_UNF_DMX_TEI_SET_S struct {
	u32DemuxID HI_U32  /* The Subdiviece ID */
	bTei       HI_BOOL /* Value of bTei, Default is HI_FALSE means receive ts packet even TEI equal 1 */
}

/* Define of how TSI and TSO to be attached */
type HI_UNF_DMX_TSI_ATTACH_TSO_S struct {
	enTSI HI_UNF_DMX_PORT_E     /* The TSI ID */
	enTSO HI_UNF_DMX_TSO_PORT_E /* The TSO ID */
}

type HI_UNF_DMX_INVOKE_TYPE_E int32

const (
	HI_UNF_DMX_INVOKE_TYPE_CHAN_CC_REPEAT_SET HI_UNF_DMX_INVOKE_TYPE_E = iota /* dmx set channel extra attr,param:HI_UNF_DMX_CHAN_CC_REPEAT_SET_S */
	HI_UNF_DMX_INVOKE_TYPE_PUSI_SET                                           /* dmx set PUSI flag,param:HI_UNF_DMX_PUSI_SET_S */
	HI_UNF_DMX_INVOKE_TYPE_TEI_SET                                            /* dmx set TEI flag,param:HI_UNF_DMX_TEI_SET_S */
	HI_UNF_DMX_INVOKE_TYPE_TSI_ATTACH_TSO                                     /* Attach TSI with TSO ,param:HI_UNF_DMX_TSI_ATTACH_TSO_S */
	HI_UNF_DMX_INVOKE_TYPE_BUTT
)

/* Define cb context type */
type HI_UNF_DMX_CB_CONTEXT_TYPE_E int32

const (
	HI_UNF_DMX_CB_CONTEXT_TYPE_SHARED  HI_UNF_DMX_CB_CONTEXT_TYPE_E = iota /* public shared context thread */
	HI_UNF_DMX_CB_CONTEXT_TYPE_PRIVATE                                     /* private context thread */
	HI_UNF_DMX_CB_CONTEXT_TYPE_BUTT
)

/* Declare section/pes/post cb function interface */
type HI_UNF_DMX_CHAN_BUF_CB_FUNC *func(hChannel HI_HANDLE, u32AcquiredNum HI_U32, pstBuf *HI_UNF_DMX_DATA_S, pUserData *HI_VOID) HI_S32

/* Define cb descriptor */
type HI_UNF_DMX_CB_DESC_S struct {
	enContextType HI_UNF_DMX_CB_CONTEXT_TYPE_E /* cb context type */
	pfnChanBufCb  HI_UNF_DMX_CHAN_BUF_CB_FUNC  /* section/pes/post cb function */
	pUserData     *HI_VOID                     /* user private data */
}

type HI_MPI_DMX_BUF_STATUS_S struct {
	u32BufSize  HI_U32 /* buffer size */
	u32UsedSize HI_U32 /* buffer used size */
	u32BufRptr  HI_U32 /* buffer read pointer */
	u32BufWptr  HI_U32 /* buffer written pointer */
}

type DMX_BUF_FLAG_E int32

const (
	DMX_MMZ_BUF DMX_BUF_FLAG_E = iota
	DMX_MMU_BUF
	DMX_SECURE_BUF
)

type DMX_BUF_S struct {
	VirAddr *HI_U8 /* Virtual address of a buffer. */
	PhyAddr HI_U32 /* Physical address of a buffer. */
	Size    HI_U32 /* Buffer size, in the unit of byte. */
	Flag    DMX_BUF_FLAG_E
}

type DMX_PoolBuf_Attr_S struct {
	BufPhyAddr HI_U32
	BufSize    HI_U32
	BufFlag    DMX_BUF_FLAG_E
}

type DMX_PORT_MODE_E int32

const (
	DMX_PORT_MODE_TUNER DMX_PORT_MODE_E = iota
	DMX_PORT_MODE_RAM
	DMX_PORT_MODE_TAG
	DMX_PORT_MODE_RMX
	DMX_PORT_MODE_BUTT
)

type DMX_MMZ_BUF_S struct {
	VirAddr *HI_U8 /* Virtual address of a buffer. */
	PhyAddr HI_U32 /* Physical address of a buffer. */
	Size    HI_U32 /* Buffer size, in the unit of byte. */
	Flag    DMX_BUF_FLAG_E
}

type DMX_DATA_BUF_S struct {
	BufKerAddr *HI_U8
	BufPhyAddr HI_U32
	BufLen     HI_U32
}

type DMX_Stream_S struct {
	pu8BufVirAddr               *HI_U8
	u32BufPhyAddr               HI_U32
	u32BufLen                   HI_U32
	u32PtsMs                    HI_U32
	u32Index                    HI_U32
	u32DispTime                 HI_U32 // add for pvr
	u32DispEnableFlag           HI_U32
	u32DispFrameDistance        HI_U32
	u32DistanceBeforeFirstFrame HI_U32
	u32GopNum                   HI_U32
}

type DMX_UserMsg_S struct {
	u32BufStartAddr HI_U32
	u32MsgLen       HI_U32
	enDataType      HI_UNF_DMX_DATA_TYPE_E /* the data packet type */
}

type HI_DRV_DMX_BUF_STATUS_S struct {
	u32BufSize  HI_U32 /* buffer size */
	u32UsedSize HI_U32 /* buffer used size */
	u32BufRptr  HI_U32 /* buffer read pointer */
	u32BufWptr  HI_U32 /* buffer written pointer */
}

type DMX_Port_GetAttr_S struct {
	PortMode DMX_PORT_MODE_E
	PortId   HI_U32
	PortAttr HI_UNF_DMX_PORT_ATTR_S
}

type DMX_Port_SetAttr_S = DMX_Port_GetAttr_S

type DMX_Tag_GetAttr_S struct {
	DmxId   HI_U32
	TagAttr HI_UNF_DMX_TAG_ATTR_S
}

type DMX_Tag_SetAttr_S = DMX_Tag_GetAttr_S

type DMX_TSO_Port_Attr_S struct {
	PortId   HI_U32
	PortAttr HI_UNF_DMX_TSO_PORT_ATTR_S
}

type DMX_Port_Attach_S struct {
	PortMode DMX_PORT_MODE_E
	PortId   HI_U32
	DmxId    HI_U32
}

type DMX_PortPacketNum_S struct {
	PortMode     DMX_PORT_MODE_E
	PortId       HI_U32
	TsPackCnt    HI_U32
	ErrTsPackCnt HI_U32
}

type DMX_Port_GetId_S = DMX_Port_Attach_S

type DMX_TsBufInit_S struct {
	PortId     HI_U32
	BufPhyAddr HI_U32
	BufSize    HI_U32
	BufFlag    DMX_BUF_FLAG_E
}

type DMX_TsBufGet_S struct {
	PortId     HI_U32
	ReqLen     HI_U32
	BufPhyAddr HI_U32
	BufSize    HI_U32
	TimeoutMs  HI_U32
}

type DMX_TsBufPush_S struct {
	PortId     HI_U32
	BufPhyAddr HI_U32
	BufSize    HI_U32
}

type DMX_TsBufRel_S struct {
	PortId     HI_U32
	BufPhyAddr HI_U32
	BufSize    HI_U32
}

type DMX_TsBufPut_S struct {
	PortId       HI_U32
	ValidDataLen HI_U32
	StartPos     HI_U32
}

type DMX_TsBufStaGet_S struct {
	PortId HI_U32
	Status HI_UNF_DMX_TSBUF_STATUS_S
}

type DMX_ChanNew_S struct {
	u32DemuxId HI_U32
	u32Pid     HI_U32
	stChAttr   HI_UNF_DMX_CHAN_ATTR_S
	hChannel   HI_HANDLE
	BufPhyAddr HI_U32
	BufSize    HI_U32
	BufFlag    DMX_BUF_FLAG_E
}

type DMX_GetChan_Attr_S struct {
	hChannel HI_HANDLE
	stChAttr HI_UNF_DMX_CHAN_ATTR_S
}

type DMX_SetChan_Attr_S = DMX_GetChan_Attr_S

type DMX_ChanPIDSet_S struct {
	hChannel HI_HANDLE
	u32Pid   HI_U32
}

type DMX_ChanPIDGet_S struct {
	hChannel HI_HANDLE
	u32Pid   HI_U32
}

type DMX_ChanStatusGet_S struct {
	hChannel HI_HANDLE
	stStatus HI_UNF_DMX_CHAN_STATUS_S
}

type DMX_ChannelIdGet_S struct {
	u32DmxId HI_U32
	u32Pid   HI_U32
	hChannel HI_HANDLE
}

type DMX_FreeChanGet_S struct {
	u32DmxId     HI_U32
	u32FreeCount HI_U32
}

type DMX_ScrambledFlagGet_S struct {
	hChannel       HI_HANDLE
	enScrambleFlag HI_UNF_DMX_SCRAMBLED_FLAG_E
}

type DMX_NewFilter_S struct {
	DmxId      HI_U32
	FilterAttr HI_UNF_DMX_FILTER_ATTR_S
	Filter     HI_HANDLE
}

type DMX_FilterSet_S struct {
	Filter     HI_HANDLE
	FilterAttr HI_UNF_DMX_FILTER_ATTR_S
}

type DMX_FilterGet_S = DMX_FilterSet_S

type DMX_FilterAttach_S struct {
	Filter  HI_HANDLE
	Channel HI_HANDLE
}

type DMX_FilterDetach_S = DMX_FilterAttach_S
type DMX_FilterChannelIDGet_S = DMX_FilterAttach_S

type DMX_FreeFilterGet_S struct {
	DmxId     HI_U32
	FreeCount HI_U32
}

//#if defined(CHIP_TYPE_hi3798cv200_a) || defined(CHIP_TYPE_hi3798cv200_b) ||defined(CHIP_TYPE_hi3798cv200) || defined(CHIP_TYPE_hi3716mv410) || defined(CHIP_TYPE_hi3716mv420)
type DMX_GetDataFlag_S struct {
	ValidChannel    *HI_HANDLE /* channel has data ready */
	ValidChannelNum *HI_U32    /* channel has data number */
	u32TimeOutMs    HI_U32
}

type DMX_Compat_GetDataFlag_S struct {
	ValidChannel    HI_U32 /* channel has data ready */
	ValidChannelNum HI_U32 /* channel has data number */
	u32TimeOutMs    HI_U32
}

type DMX_SelectDataFlag_S struct {
	channel         *HI_HANDLE /* channel handles to check */
	channelnum      HI_U32     /* channel number to check */
	ValidChannel    *HI_HANDLE /* channel has data ready */
	ValidChannelNum *HI_U32    /* channel has data number */
	u32TimeOutMs    HI_U32     /* timeout time in MS */
}

type DMX_Compat_SelectDataFlag_S struct {
	channel         HI_U32 /* channel handles to check */
	channelnum      HI_U32 /* channel number to check */
	ValidChannel    HI_U32 /* channel has data ready */
	ValidChannelNum HI_U32 /* channel has data number */
	u32TimeOutMs    HI_U32 /* timeout time in MS */
}

//#else
//type DMX_GetDataFlag_S struct {
//	u32Flag      [3]HI_U32
//	u32TimeOutMs HI_U32
//}

//type DMX_SelectDataFlag_S struct {
//	channel      *HI_HANDLE /* channel handles to check */
//	channelnum   HI_U32     /* channel number to check */
//	u32Flag      [3]HI_U32  /* dataflag */
//	u32TimeOutMs HI_U32     /* timeout time in MS */
//}
//#endif
