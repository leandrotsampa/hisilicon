package hisilicon

/* Structs Definitions */
/** Global config structure */
type HI_SYS_CONF_S struct {
	u32Reverse HI_U32 /* Not used, reserved for extension */
}

/* Define the chip type. */
type HI_CHIP_TYPE_E int32

const (
	HI_CHIP_TYPE_HI3716M HI_CHIP_TYPE_E = iota
	HI_CHIP_TYPE_HI3716H
	HI_CHIP_TYPE_HI3716C
	HI_CHIP_TYPE_HI3716CES

	HI_CHIP_TYPE_HI3720
	HI_CHIP_TYPE_HI3712
	HI_CHIP_TYPE_HI3715

	HI_CHIP_TYPE_HI3718M
	HI_CHIP_TYPE_HI3718C
	HI_CHIP_TYPE_HI3719M
	HI_CHIP_TYPE_HI3719C
	HI_CHIP_TYPE_HI3719M_A
)

const (
	HI_CHIP_TYPE_HI3796C HI_CHIP_TYPE_E = iota + 0x20
	HI_CHIP_TYPE_HI3798C
	HI_CHIP_TYPE_HI3796M
	HI_CHIP_TYPE_HI3798M
)
const (
	HI_CHIP_TYPE_HI3796C_A HI_CHIP_TYPE_E = iota + 0x40
	HI_CHIP_TYPE_HI3798C_A
	HI_CHIP_TYPE_HI3798C_B

	HI_CHIP_TYPE_HI3798M_A

	HI_CHIP_TYPE_BUTT
)

/* Define the chip version. */
type HI_CHIP_VERSION_E int32

const (
	HI_CHIP_VERSION_V100 HI_CHIP_VERSION_E = 0x100
	HI_CHIP_VERSION_V101 HI_CHIP_VERSION_E = 0x101
	HI_CHIP_VERSION_V200 HI_CHIP_VERSION_E = 0x200
	HI_CHIP_VERSION_V300 HI_CHIP_VERSION_E = 0x300
	HI_CHIP_VERSION_V400 HI_CHIP_VERSION_E = 0x400
	HI_CHIP_VERSION_V410 HI_CHIP_VERSION_E = 0x410
	HI_CHIP_VERSION_V420 HI_CHIP_VERSION_E = 0x420
	HI_CHIP_VERSION_BUTT
)

type HI_CHIP_PACKAGE_TYPE_E int32

const (
	HI_CHIP_PACKAGE_TYPE_BGA_15_15 HI_CHIP_PACKAGE_TYPE_E = iota
	HI_CHIP_PACKAGE_TYPE_BGA_16_16
	HI_CHIP_PACKAGE_TYPE_BGA_19_19
	HI_CHIP_PACKAGE_TYPE_BGA_23_23
	HI_CHIP_PACKAGE_TYPE_BGA_31_31
	HI_CHIP_PACKAGE_TYPE_QFP_216
	HI_CHIP_PACKAGE_TYPE_BUTT
)

/* Define the chip support attrs */
type HI_CHIP_CAP_E int32

const (
	HI_CHIP_CAP_DOLBY HI_CHIP_CAP_E = iota
	HI_CHIP_CAP_DTS
	HI_CHIP_CAP_ADVCA
	HI_CHIP_CAP_MACROVISION
)

/* System version, that is, the version of the software developer's kit (SDK) */
type HI_SYS_VERSION_S struct {
	enChipTypeSoft     HI_CHIP_TYPE_E    /* Chip type corresponding to the SDK */
	enChipTypeHardWare HI_CHIP_TYPE_E    /* Chip type that is detected when the SDK is running */
	enChipVersion      HI_CHIP_VERSION_E /* Chip version that is detected when the SDK is running */
	aVersion           [80]HI_CHAR       /* Version string of the SDK */
	BootVersion        [80]HI_CHAR       /* Version string of the Boot */
}

/* Define the chip attributes */
type HI_SYS_CHIP_ATTR_S struct {
	bDolbySupport       HI_BOOL /* Whether support dolby or not */
	bDTSSupport         HI_BOOL /* Whether support DTS or not */
	bADVCASupport       HI_BOOL /* Whether support ADVCA or not */
	bMacrovisionSupport HI_BOOL /* Whether support Macrovision or not */
	u64ChipID           HI_U64  /* the unique chipid */
}

/* Structure of an MMZ buffer */
const MAX_BUFFER_NAME_SIZE = 16

type HI_MMZ_BUF_S struct {
	bufname             [MAX_BUFFER_NAME_SIZE]HI_CHAR /* Strings of an MMZ buffer name */
	phyaddr             HI_U32                        /* Physical address of an MMZ buffer */
	kernel_viraddr      *HI_U8                        /* Kernel-state virtual address of an MMZ buffer */
	user_viraddr        *HI_U8                        /* User-state virtual address of an MMZ buffer */
	bufsize             HI_U32                        /* Size of an MMZ buffer */
	overflow_threshold  HI_U32                        /* Overflow threshold of an MMZ buffer, in percentage. For example, the value 100 indicates 100%. */
	underflow_threshold HI_U32                        /* Underflow threshold of an MMZ buffer, in percentage. For example, the value 0 indicates 0%. */
}

type HI_RECT_S struct {
	s32X      HI_S32
	s32Y      HI_S32
	s32Width  HI_S32
	s32Height HI_S32
}

type HI_LAYER_ZORDER_E int32

const (
	HI_LAYER_ZORDER_MOVETOP    HI_LAYER_ZORDER_E = iota /* Move to the top */
	HI_LAYER_ZORDER_MOVEUP                              /* Move up */
	HI_LAYER_ZORDER_MOVEBOTTOM                          /* Move to the bottom */
	HI_LAYER_ZORDER_MOVEDOWN                            /* Move down */
	HI_LAYER_ZORDER_BUTT
)

/* Defines user mode proc show buffer */
type HI_PROC_SHOW_BUFFER_S struct {
	pu8Buf    *HI_U8 /* Buffer address */
	u32Size   HI_U32 /* Buffer size */
	u32Offset HI_U32 /* Offset */
}

/* Proc show function */
type HI_PROC_SHOW_FN *func(pstBuf *HI_PROC_SHOW_BUFFER_S, pPrivData *HI_VOID) HI_S32

/* Proc command function */
type HI_PROC_CMD_FN *func(pstBuf *HI_PROC_SHOW_BUFFER_S, u32Argc HI_U32, pu8Argv *[]HI_U8, pPrivData *HI_VOID) HI_S32

/* Defines user mode proc entry */
type HI_PROC_ENTRY_S struct {
	pszEntryName *HI_CHAR        /* Entry name */
	pszDirectory *HI_CHAR        /* Directory name. If null, the entry will be added to /proc/hisi directory */
	pfnShowProc  HI_PROC_SHOW_FN /* Proc show function */
	pfnCmdProc   HI_PROC_CMD_FN  /* Proc command function */
	pPrivData    *HI_VOID        /* Private data */
}

/* Defines DDR configuration type struct */
type HI_SYS_MEM_CONFIG_S struct {
	u32TotalSize HI_U32 /* Total memory size(MB) */
	u32MMZSize   HI_U32 /* MMZ memory size(MB) */
}
