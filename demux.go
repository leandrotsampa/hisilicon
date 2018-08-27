package hisilicon

import (
	"errors"
	"os"
	"strings"
	"sync"
	"unsafe"
)

type DMX_CB_DESC_S struct {
	Handle HI_HANDLE
	Raw    HI_UNF_DMX_CB_DESC_S
	//Context pthread_t
}

var (
	demux       = HiDevice{name: "hi_demux", fd: nil, InUse: 0}
	g_stPoolBuf DMX_BUF_S
	g_stTsBuf   struct {
		RamPortCnt HI_U32
		BufDesc    []DMX_BUF_S
	}
	g_stChanBuf struct {
		DmxCnt  HI_U32
		ChnCnt  HI_U32
		BufDesc []DMX_BUF_S

		CbDescEx        []DMX_CB_DESC_S
		SharedChanCbCxt struct {
			//Context pthread_t
			Ref HI_U32
		}
		CbLock sync.Mutex
	}
	g_stRecBuf struct {
		DmxCnt  HI_U32
		RecCnt  HI_U32
		BufDesc []DMX_BUF_S
	}

	/* global */
	CMD_DEMUX_GET_POOLBUF_ADDR uintptr
	CMD_DEMUX_GET_CAPABILITY   uintptr
	CMD_DEMUX_SET_PUSI         uintptr
	CMD_DEMUX_SET_TEI          uintptr
	CMD_DEMUX_TSI_ATTACH_TSO   uintptr
	CMD_DEMUX_GET_RESUME_COUNT uintptr

	/* TS PORT */
	CMD_DEMUX_PORT_GET_ATTR     uintptr /* get port attr */
	CMD_DEMUX_PORT_SET_ATTR     uintptr /* set port attr */
	CMD_DEMUX_PORT_ATTACH       uintptr /* attach ts port to demux */
	CMD_DEMUX_PORT_DETACH       uintptr /* detach ts port from demux */
	CMD_DEMUX_PORT_GETID        uintptr /* get ts port id of demux */
	CMD_DEMUX_PORT_GETPACKETNUM uintptr /* get ts pack counter */
	CMD_DEMUX_TSO_PORT_GET_ATTR uintptr /* get TSO port attr */
	CMD_DEMUX_TSO_PORT_SET_ATTR uintptr /* Set TSO port attr */
	CMD_DEMUX_DMX_GET_TAG_ATTR  uintptr /* get port tag attrs */
	CMD_DEMUX_DMX_SET_TAG_ATTR  uintptr /* set port tag attrs */

	/* Ts Buffer */
	CMD_DEMUX_TS_BUFFER_INIT       uintptr /* TS Buffer init */
	CMD_DEMUX_TS_BUFFER_DEINIT     uintptr /* TS Buffer deinit */
	CMD_DEMUX_TS_BUFFER_GET        uintptr /* Get TS Buffer */
	CMD_DEMUX_TS_BUFFER_PUT        uintptr /* Put TS Buffer */
	CMD_DEMUX_TS_BUFFER_RESET      uintptr /* Reset TS Buffer */
	CMD_DEMUX_TS_BUFFER_GET_STATUS uintptr /* Get TS Buffer status */
	CMD_DEMUX_TS_BUFFER_PUSH       uintptr /* Push TS Buffer status*/
	CMD_DEMUX_TS_BUFFER_RELEASE    uintptr /* Release TS Buffer status*/

	/* Channel */
	CMD_DEMUX_CHAN_NEW           uintptr /* apply for a free channel */
	CMD_DEMUX_CHAN_NEW2          uintptr /* apply for a free channel */
	CMD_DEMUX_CHAN_DEL           uintptr /* delete an allocated channel */
	CMD_DEMUX_CHAN_OPEN          uintptr /* open channel */
	CMD_DEMUX_CHAN_CLOSE         uintptr /* close channel */
	CMD_DEMUX_CHAN_ATTR_GET      uintptr
	CMD_DEMUX_CHAN_ATTR_SET      uintptr
	CMD_DEMUX_GET_CHAN_STATUS    uintptr /* get channel open/close status */
	CMD_DEMUX_PID_SET            uintptr /* set pid of channel */
	CMD_DEMUX_PID_GET            uintptr /* get pid of channel */
	CMD_DEMUX_CHANID_GET         uintptr /* get channel id with the designated pid */
	CMD_DEMUX_FREECHAN_GET       uintptr /* get free channel counter */
	CMD_DEMUX_SCRAMBLEFLAG_GET   uintptr /* get scrambed flag of audio channel */
	CMD_DEMUX_CHAN_SET_EOS_FLAG  uintptr
	CMD_DEMUX_CHAN_CC_REPEAT_SET uintptr /* set channel CC repeat attr*/
	CMD_DEMUX_GET_CHAN_TSCNT     uintptr /* get channel ts count */

	/* Filter */
	CMD_DEMUX_FLT_NEW        uintptr /* apply for a free filter */
	CMD_DEMUX_FLT_DEL        uintptr /* delete an allocated filter */
	CMD_DEMUX_FLT_SET        uintptr /* set fiter parameter */
	CMD_DEMUX_FLT_GET        uintptr /* get fiter parameter */
	CMD_DEMUX_FLT_ATTACH     uintptr /* attach a filter to a channel */
	CMD_DEMUX_FLT_DETACH     uintptr /* detach a filter from a channel */
	CMD_DEMUX_FREEFLT_GET    uintptr /* get free filter coute */
	CMD_DEMUX_FLT_DELALL     uintptr /* delete all filters on a channel */
	CMD_DEMUX_FLT_CHANID_GET uintptr

	/* data receive */
	CMD_DEMUX_GET_DATA_FLAG           uintptr /* get data flag of dma buffer */
	CMD_DEMUX_COMPAT_GET_DATA_FLAG    uintptr /* get data flag of dma buffer */
	CMD_DEMUX_ACQUIRE_MSG             uintptr
	CMD_DEMUX_COMPAT_ACQUIRE_MSG      uintptr
	CMD_DEMUX_RELEASE_MSG             uintptr
	CMD_DEMUX_COMPAT_RELEASE_MSG      uintptr
	CMD_DEMUX_SELECT_DATA_FLAG        uintptr
	CMD_DEMUX_COMPAT_SELECT_DATA_FLAG uintptr

	/* PCR */
	CMD_DEMUX_PCR_NEW       uintptr /* set pcr pid */
	CMD_DEMUX_PCR_DEL       uintptr /* set pcr pid */
	CMD_DEMUX_PCRPID_SET    uintptr /* set pcr pid */
	CMD_DEMUX_PCRPID_GET    uintptr /* get pcr pid */
	CMD_DEMUX_CURPCR_GET    uintptr /* get pcr count */
	CMD_DEMUX_PCRSYN_ATTACH uintptr /* attach pcr channel and sync handle */
	CMD_DEMUX_PCRSYN_DETACH uintptr /* detach pcr channel and sync handle */

	/* AV */
	CMD_DEMUX_PES_BUFFER_GETSTAT uintptr /* Get PES Buffer status */
	CMD_DEMUX_ES_BUFFER_GET      uintptr /* Get ES Buffer */
	CMD_DEMUX_ES_BUFFER_PUT      uintptr /* Put ES Buffer */

	/* REC */
	CMD_DEMUX_REC_CHAN_CREATE                    uintptr
	CMD_DEMUX_REC_CHAN_DESTROY                   uintptr
	CMD_DEMUX_REC_CHAN_ADD_PID                   uintptr
	CMD_DEMUX_REC_CHAN_DEL_PID                   uintptr
	CMD_DEMUX_REC_CHAN_DEL_ALL_PID               uintptr
	CMD_DEMUX_REC_CHAN_ADD_EXCLUDE_PID           uintptr
	CMD_DEMUX_REC_CHAN_DEL_EXCLUDE_PID           uintptr
	CMD_DEMUX_REC_CHAN_CANCEL_EXCLUDE            uintptr
	CMD_DEMUX_REC_CHAN_START                     uintptr
	CMD_DEMUX_REC_CHAN_STOP                      uintptr
	CMD_DEMUX_REC_CHAN_ACQUIRE_DATA              uintptr
	CMD_DEMUX_REC_CHAN_RELEASE_DATA              uintptr
	CMD_DEMUX_REC_CHAN_ACQUIRE_INDEX             uintptr
	CMD_DEMUX_REC_CHAN_GET_BUF_STATUS            uintptr
	CMD_DEMUX_REC_CHAN_ACQUIRE_DATA_INDEX        uintptr
	CMD_DEMUX_COMPAT_REC_CHAN_ACQUIRE_DATA_INDEX uintptr
	CMD_DEMUX_REC_CHAN_RELEASE_DATA_INDEX        uintptr
	CMD_DEMUX_COMPAT_REC_CHAN_RELEASE_DATA_INDEX uintptr

	/* RMX */
	CMD_REMUX_CREATE                uintptr
	CMD_REMUX_DESTROY               uintptr
	CMD_REMUX_GET_ATTR              uintptr
	CMD_REMUX_SET_ATTR              uintptr
	CMD_REMUX_START                 uintptr
	CMD_REMUX_STOP                  uintptr
	CMD_REMUX_ADD_PUMP              uintptr
	CMD_REMUX_DEL_PUMP              uintptr
	CMD_REMUX_GET_PUMP_DEFAULT_ATTR uintptr
	CMD_REMUX_GET_PUMP_ATTR         uintptr
	CMD_REMUX_SET_PUMP_ATTR         uintptr
)

/** Internal Function for Demux IOCTL Calls **/
func demuxLoadIoctl() error {
	var id uint32
	var err error

	if demux.fd != nil {
		return nil
	} else if _, err = HI_MODULE_Init(); err != nil {
		return err
	} else if id, err = HI_MODULE_GetModuleID(strings.ToUpper(demux.name)); err != nil {
		return err
	}

	/* global */
	CMD_DEMUX_GET_POOLBUF_ADDR = IoR(uintptr(id), 0x00, unsafe.Sizeof(DMX_PoolBuf_Attr_S{}))
	CMD_DEMUX_GET_CAPABILITY = IoR(uintptr(id), 0x01, unsafe.Sizeof(HI_UNF_DMX_CAPABILITY_S{}))
	CMD_DEMUX_SET_PUSI = IoW(uintptr(id), 0x02, unsafe.Sizeof(HI_UNF_DMX_PUSI_SET_S{}))
	CMD_DEMUX_SET_TEI = IoW(uintptr(id), 0x03, unsafe.Sizeof(HI_UNF_DMX_TEI_SET_S{}))
	CMD_DEMUX_TSI_ATTACH_TSO = IoW(uintptr(id), 0x04, unsafe.Sizeof(HI_UNF_DMX_TSI_ATTACH_TSO_S{}))
	CMD_DEMUX_GET_RESUME_COUNT = IoRW(uintptr(id), 0x05, unsafe.Sizeof(HI_U32(0)))

	/* TS PORT */
	CMD_DEMUX_PORT_GET_ATTR = IoRW(uintptr(id), 0x10, unsafe.Sizeof(DMX_Port_GetAttr_S{}))
	CMD_DEMUX_PORT_SET_ATTR = IoW(uintptr(id), 0x11, unsafe.Sizeof(DMX_Port_SetAttr_S{}))
	CMD_DEMUX_PORT_ATTACH = IoW(uintptr(id), 0x12, unsafe.Sizeof(DMX_Port_Attach_S{}))
	CMD_DEMUX_PORT_DETACH = IoW(uintptr(id), 0x13, unsafe.Sizeof(HI_U32(0)))
	CMD_DEMUX_PORT_GETID = IoRW(uintptr(id), 0x14, unsafe.Sizeof(DMX_Port_GetId_S{}))
	CMD_DEMUX_PORT_GETPACKETNUM = IoRW(uintptr(id), 0x15, unsafe.Sizeof(DMX_PortPacketNum_S{}))
	CMD_DEMUX_TSO_PORT_GET_ATTR = IoRW(uintptr(id), 0x16, unsafe.Sizeof(DMX_TSO_Port_Attr_S{}))
	CMD_DEMUX_TSO_PORT_SET_ATTR = IoW(uintptr(id), 0x17, unsafe.Sizeof(DMX_TSO_Port_Attr_S{}))
	CMD_DEMUX_DMX_GET_TAG_ATTR = IoRW(uintptr(id), 0x18, unsafe.Sizeof(DMX_Tag_GetAttr_S{}))
	CMD_DEMUX_DMX_SET_TAG_ATTR = IoW(uintptr(id), 0x19, unsafe.Sizeof(DMX_Tag_SetAttr_S{}))

	/* Ts Buffer */
	CMD_DEMUX_TS_BUFFER_INIT = IoRW(uintptr(id), 0x20, unsafe.Sizeof(DMX_TsBufInit_S{}))
	CMD_DEMUX_TS_BUFFER_DEINIT = IoW(uintptr(id), 0x21, unsafe.Sizeof(HI_U32(0)))
	CMD_DEMUX_TS_BUFFER_GET = IoRW(uintptr(id), 0x22, unsafe.Sizeof(DMX_TsBufGet_S{}))
	CMD_DEMUX_TS_BUFFER_PUT = IoW(uintptr(id), 0x23, unsafe.Sizeof(DMX_TsBufPut_S{}))
	CMD_DEMUX_TS_BUFFER_RESET = IoW(uintptr(id), 0x24, unsafe.Sizeof(HI_U32(0)))
	CMD_DEMUX_TS_BUFFER_GET_STATUS = IoRW(uintptr(id), 0x25, unsafe.Sizeof(DMX_TsBufStaGet_S{}))
	CMD_DEMUX_TS_BUFFER_PUSH = IoRW(uintptr(id), 0x26, unsafe.Sizeof(DMX_TsBufPush_S{}))
	CMD_DEMUX_TS_BUFFER_RELEASE = IoRW(uintptr(id), 0x27, unsafe.Sizeof(DMX_TsBufRel_S{}))

	/* Channel */
	CMD_DEMUX_CHAN_NEW = IoRW(uintptr(id), 0x30, unsafe.Sizeof(DMX_ChanNew_S{}))
	CMD_DEMUX_CHAN_NEW2 = IoRW(uintptr(id), 0x31, unsafe.Sizeof(DMX_ChanNew_S{}))
	CMD_DEMUX_CHAN_DEL = IoW(uintptr(id), 0x32, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_CHAN_OPEN = IoW(uintptr(id), 0x33, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_CHAN_CLOSE = IoW(uintptr(id), 0x34, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_CHAN_ATTR_GET = IoRW(uintptr(id), 0x35, unsafe.Sizeof(DMX_GetChan_Attr_S{}))
	CMD_DEMUX_CHAN_ATTR_SET = IoW(uintptr(id), 0x36, unsafe.Sizeof(DMX_SetChan_Attr_S{}))
	CMD_DEMUX_GET_CHAN_STATUS = IoRW(uintptr(id), 0x37, unsafe.Sizeof(DMX_ChanStatusGet_S{}))
	CMD_DEMUX_PID_SET = IoW(uintptr(id), 0x38, unsafe.Sizeof(DMX_ChanPIDSet_S{}))
	CMD_DEMUX_PID_GET = IoRW(uintptr(id), 0x39, unsafe.Sizeof(DMX_ChanPIDGet_S{}))
	CMD_DEMUX_CHANID_GET = IoRW(uintptr(id), 0x3A, unsafe.Sizeof(DMX_ChannelIdGet_S{}))
	CMD_DEMUX_FREECHAN_GET = IoRW(uintptr(id), 0x3B, unsafe.Sizeof(DMX_FreeChanGet_S{}))
	CMD_DEMUX_SCRAMBLEFLAG_GET = IoRW(uintptr(id), 0x3C, unsafe.Sizeof(DMX_ScrambledFlagGet_S{}))
	CMD_DEMUX_CHAN_SET_EOS_FLAG = IoRW(uintptr(id), 0x3D, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_CHAN_CC_REPEAT_SET = IoW(uintptr(id), 0x3E, unsafe.Sizeof(DMX_SetChan_CC_REPEAT_S{}))
	CMD_DEMUX_GET_CHAN_TSCNT = IoRW(uintptr(id), 0x3F, unsafe.Sizeof(DMX_ChanChanTsCnt_S{}))

	/* Filter */
	CMD_DEMUX_FLT_NEW = IoRW(uintptr(id), 0x40, unsafe.Sizeof(DMX_NewFilter_S{}))
	CMD_DEMUX_FLT_DEL = IoW(uintptr(id), 0x41, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_FLT_SET = IoW(uintptr(id), 0x42, unsafe.Sizeof(DMX_FilterSet_S{}))
	CMD_DEMUX_FLT_GET = IoRW(uintptr(id), 0x43, unsafe.Sizeof(DMX_FilterGet_S{}))
	CMD_DEMUX_FLT_ATTACH = IoW(uintptr(id), 0x44, unsafe.Sizeof(DMX_FilterAttach_S{}))
	CMD_DEMUX_FLT_DETACH = IoW(uintptr(id), 0x45, unsafe.Sizeof(DMX_FilterDetach_S{}))
	CMD_DEMUX_FREEFLT_GET = IoRW(uintptr(id), 0x46, unsafe.Sizeof(DMX_FreeFilterGet_S{}))
	CMD_DEMUX_FLT_DELALL = IoW(uintptr(id), 0x47, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_FLT_CHANID_GET = IoRW(uintptr(id), 0x48, unsafe.Sizeof(DMX_FilterChannelIDGet_S{}))

	/* data receive */
	CMD_DEMUX_GET_DATA_FLAG = IoRW(uintptr(id), 0x60, unsafe.Sizeof(DMX_GetDataFlag_S{}))
	CMD_DEMUX_COMPAT_GET_DATA_FLAG = IoRW(uintptr(id), 0x60, unsafe.Sizeof(DMX_Compat_GetDataFlag_S{}))
	CMD_DEMUX_ACQUIRE_MSG = IoRW(uintptr(id), 0x61, unsafe.Sizeof(DMX_AcqMsg_S{}))
	CMD_DEMUX_COMPAT_ACQUIRE_MSG = IoRW(uintptr(id), 0x61, unsafe.Sizeof(DMX_Compat_AcqMsg_S{}))
	CMD_DEMUX_RELEASE_MSG = IoW(uintptr(id), 0x62, unsafe.Sizeof(DMX_RelMsg_S{}))
	CMD_DEMUX_COMPAT_RELEASE_MSG = IoW(uintptr(id), 0x62, unsafe.Sizeof(DMX_Compat_RelMsg_S{}))
	CMD_DEMUX_SELECT_DATA_FLAG = IoRW(uintptr(id), 0x63, unsafe.Sizeof(DMX_SelectDataFlag_S{}))
	CMD_DEMUX_COMPAT_SELECT_DATA_FLAG = IoRW(uintptr(id), 0x63, unsafe.Sizeof(DMX_Compat_SelectDataFlag_S{}))

	/* PCR */
	CMD_DEMUX_PCR_NEW = IoRW(uintptr(id), 0x70, unsafe.Sizeof(DMX_NewPcr_S{}))
	CMD_DEMUX_PCR_DEL = IoW(uintptr(id), 0x71, unsafe.Sizeof(HI_U32(0)))
	CMD_DEMUX_PCRPID_SET = IoW(uintptr(id), 0x72, unsafe.Sizeof(DMX_PcrPidSet_S{}))
	CMD_DEMUX_PCRPID_GET = IoRW(uintptr(id), 0x73, unsafe.Sizeof(DMX_PcrPidGet_S{}))
	CMD_DEMUX_CURPCR_GET = IoRW(uintptr(id), 0x74, unsafe.Sizeof(DMX_PcrScrGet_S{}))
	CMD_DEMUX_PCRSYN_ATTACH = IoRW(uintptr(id), 0x75, unsafe.Sizeof(DMX_PCRSYNC_S{}))
	CMD_DEMUX_PCRSYN_DETACH = IoRW(uintptr(id), 0x76, unsafe.Sizeof(DMX_PCRSYNC_S{}))

	/* AV */
	CMD_DEMUX_PES_BUFFER_GETSTAT = IoRW(uintptr(id), 0x80, unsafe.Sizeof(DMX_PesBufStaGet_S{}))
	CMD_DEMUX_ES_BUFFER_GET = IoRW(uintptr(id), 0x81, unsafe.Sizeof(DMX_PesBufGet_S{}))
	CMD_DEMUX_ES_BUFFER_PUT = IoW(uintptr(id), 0x82, unsafe.Sizeof(DMX_PesBufGet_S{}))

	/* REC */
	CMD_DEMUX_REC_CHAN_CREATE = IoRW(uintptr(id), 0x90, unsafe.Sizeof(DMX_Rec_CreateChan_S{}))
	CMD_DEMUX_REC_CHAN_DESTROY = IoW(uintptr(id), 0x91, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_REC_CHAN_ADD_PID = IoRW(uintptr(id), 0x92, unsafe.Sizeof(DMX_Rec_AddPid_S{}))
	CMD_DEMUX_REC_CHAN_DEL_PID = IoW(uintptr(id), 0x93, unsafe.Sizeof(DMX_Rec_DelPid_S{}))
	CMD_DEMUX_REC_CHAN_DEL_ALL_PID = IoW(uintptr(id), 0x94, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_REC_CHAN_ADD_EXCLUDE_PID = IoW(uintptr(id), 0x95, unsafe.Sizeof(DMX_Rec_ExcludePid_S{}))
	CMD_DEMUX_REC_CHAN_DEL_EXCLUDE_PID = IoW(uintptr(id), 0x96, unsafe.Sizeof(DMX_Rec_ExcludePid_S{}))
	CMD_DEMUX_REC_CHAN_CANCEL_EXCLUDE = IoW(uintptr(id), 0x97, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_REC_CHAN_START = IoW(uintptr(id), 0x98, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_REC_CHAN_STOP = IoW(uintptr(id), 0x99, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_DEMUX_REC_CHAN_ACQUIRE_DATA = IoRW(uintptr(id), 0x9A, unsafe.Sizeof(DMX_Rec_AcquireData_S{}))
	CMD_DEMUX_REC_CHAN_RELEASE_DATA = IoW(uintptr(id), 0x9B, unsafe.Sizeof(DMX_Rec_ReleaseData_S{}))
	CMD_DEMUX_REC_CHAN_ACQUIRE_INDEX = IoRW(uintptr(id), 0x9C, unsafe.Sizeof(DMX_Rec_AcquireIndex_S{}))
	CMD_DEMUX_REC_CHAN_GET_BUF_STATUS = IoRW(uintptr(id), 0x9D, unsafe.Sizeof(DMX_Rec_BufStatus_S{}))
	CMD_DEMUX_REC_CHAN_ACQUIRE_DATA_INDEX = IoRW(uintptr(id), 0x9E, unsafe.Sizeof(DMX_Rec_ProcessDataIndex_S{}))
	CMD_DEMUX_COMPAT_REC_CHAN_ACQUIRE_DATA_INDEX = IoRW(uintptr(id), 0x9E, unsafe.Sizeof(DMX_Compat_Rec_ProcessDataIndex_S{}))
	CMD_DEMUX_REC_CHAN_RELEASE_DATA_INDEX = IoRW(uintptr(id), 0x9F, unsafe.Sizeof(DMX_Rec_ProcessDataIndex_S{}))
	CMD_DEMUX_COMPAT_REC_CHAN_RELEASE_DATA_INDEX = IoRW(uintptr(id), 0x9F, unsafe.Sizeof(DMX_Compat_Rec_ProcessDataIndex_S{}))

	/* RMX */
	CMD_REMUX_CREATE = IoRW(uintptr(id), 0xA0, unsafe.Sizeof(RMX_Create_S{}))
	CMD_REMUX_DESTROY = IoW(uintptr(id), 0xA1, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_REMUX_GET_ATTR = IoRW(uintptr(id), 0xA2, unsafe.Sizeof(RMX_Attr_S{}))
	CMD_REMUX_SET_ATTR = IoRW(uintptr(id), 0xA3, unsafe.Sizeof(RMX_Attr_S{}))
	CMD_REMUX_START = IoW(uintptr(id), 0xA4, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_REMUX_STOP = IoW(uintptr(id), 0xA5, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_REMUX_ADD_PUMP = IoRW(uintptr(id), 0xA6, unsafe.Sizeof(RMX_Add_Pump_S{}))
	CMD_REMUX_DEL_PUMP = IoRW(uintptr(id), 0xA7, unsafe.Sizeof(HI_HANDLE(0)))
	CMD_REMUX_GET_PUMP_DEFAULT_ATTR = IoRW(uintptr(id), 0xA8, unsafe.Sizeof(RMX_Pump_Attr_S{}))
	CMD_REMUX_GET_PUMP_ATTR = IoRW(uintptr(id), 0xA9, unsafe.Sizeof(RMX_Pump_Attr_S{}))
	CMD_REMUX_SET_PUMP_ATTR = IoRW(uintptr(id), 0xAA, unsafe.Sizeof(RMX_Pump_Attr_S{}))

	HI_MODULE_DeInit()
	return nil
}

func demuxCall(op uintptr, arg interface{}) (bool, error) {
	if demux.fd == nil {
		return false, errors.New("Demux Device not initialized.")
	}

	if err := Ioctl(demux.fd.Fd(), op, arg); err != nil {
		return false, err
	}

	return true, nil
}

/*************************************************************
Function:       HI_UNF_DMX_Init
Description:    open demux device, and do the basical initialization
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_DMX_Init() (bool, error) {
	demux.mu.Lock()
	defer demux.mu.Unlock()

	if demux.fd != nil {
		demux.InUse++
		return true, nil
	}

	var err error
	if err = demuxLoadIoctl(); err != nil {
		return false, err
	} else if demux.fd, err = os.OpenFile("/dev/"+demux.name, os.O_RDWR, 0); err != nil {
		return false, err
	}

	Cap := HI_UNF_DMX_CAPABILITY_S{}
	if _, err := HI_UNF_DMX_GetCapability(&Cap); err != nil {
		demux.fd.Close()
		return false, err
	}

	/* save ts buff desc info */
	g_stTsBuf.RamPortCnt = Cap.u32RamPortNum
	g_stTsBuf.BufDesc = make([]DMX_BUF_S, g_stTsBuf.RamPortCnt)

	/* save channel buf desc and handle info */
	g_stChanBuf.DmxCnt = Cap.u32DmxNum
	g_stChanBuf.ChnCnt = Cap.u32ChannelNum
	g_stChanBuf.BufDesc = make([]DMX_BUF_S, g_stChanBuf.DmxCnt*g_stChanBuf.ChnCnt)

	/* save call back desc. */
	g_stChanBuf.CbDescEx = make([]DMX_CB_DESC_S, g_stChanBuf.DmxCnt*g_stChanBuf.ChnCnt)

	for i := 0; i < len(g_stChanBuf.CbDescEx); i++ {
		CbDescEx := &g_stChanBuf.CbDescEx[i]

		CbDescEx.Handle = HI_INVALID_HANDLE
		//CbDescEx.Context = -1
		CbDescEx.Raw.enContextType = HI_UNF_DMX_CB_CONTEXT_TYPE_BUTT
		CbDescEx.Raw.pfnChanBufCb = nil
		CbDescEx.Raw.pUserData = nil
	}

	//g_stChanBuf.SharedChanCbCxt.Context = -1
	g_stChanBuf.SharedChanCbCxt.Ref = 0

	/* save rec buf desc info */
	g_stRecBuf.DmxCnt = Cap.u32DmxNum
	g_stRecBuf.RecCnt = Cap.u32RecChnNum
	g_stRecBuf.BufDesc = make([]DMX_BUF_S, g_stRecBuf.DmxCnt*g_stRecBuf.RecCnt)

	PoolBufParam := DMX_PoolBuf_Attr_S{}
	if _, err := demuxCall(CMD_DEMUX_GET_POOLBUF_ADDR, &PoolBufParam); err != nil {
		demux.fd.Close()
		return false, err
	}

	g_stPoolBuf.PhyAddr = PoolBufParam.BufPhyAddr
	g_stPoolBuf.Size = PoolBufParam.BufSize
	g_stPoolBuf.Flag = PoolBufParam.BufFlag

	//g_stPoolBuf.VirAddr = DmxMmap(g_stPoolBuf.PhyAddr, g_stPoolBuf.Flag)
	/*if (0 == g_stPoolBuf.VirAddr)
	  {
	      HI_ERR_DEMUX("Pool buffer mmap error\n");
	      ret =  HI_ERR_DMX_MMAP_FAILED;
	      goto out6;
	  }*/

	return true, nil
}

/*************************************************************
Function:       HI_UNF_DMX_DeInit
Description:    close ir device
Calls:
Data Accessed:
Data Updated:   NA
Input:
Output:
Return:         bool
                error
Others:         NA
*************************************************************/
func HI_UNF_DMX_DeInit() (bool, error) {
	demux.mu.Lock()
	defer demux.mu.Unlock()

	if demux.fd == nil {
		return true, nil
	} else if demux.InUse > 0 {
		demux.InUse--
		return true, nil
	}

	if err := demux.fd.Close(); err != nil {
		return false, err
	}

	return true, nil
}

func HI_UNF_DMX_GetCapability(pstCap *HI_UNF_DMX_CAPABILITY_S) (bool, error) {
	return demuxCall(CMD_DEMUX_GET_CAPABILITY, pstCap)
}
