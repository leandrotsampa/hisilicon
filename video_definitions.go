package hisilicon

/* Structs Definitions */
/* Defines the stream type supported by the decoder. */
type HI_UNF_VCODEC_TYPE_E int32

const (
	HI_UNF_VCODEC_TYPE_MPEG2     HI_UNF_VCODEC_TYPE_E = iota /* MPEG2 */
	HI_UNF_VCODEC_TYPE_MPEG4                                 /* MPEG4 DIVX4 DIVX5 */
	HI_UNF_VCODEC_TYPE_AVS                                   /* AVS */
	HI_UNF_VCODEC_TYPE_H263                                  /* H263 */
	HI_UNF_VCODEC_TYPE_H264                                  /* H264 */
	HI_UNF_VCODEC_TYPE_REAL8                                 /* REAL */
	HI_UNF_VCODEC_TYPE_REAL9                                 /* REAL */
	HI_UNF_VCODEC_TYPE_VC1                                   /* VC-1 */
	HI_UNF_VCODEC_TYPE_VP6                                   /* VP6 */
	HI_UNF_VCODEC_TYPE_VP6F                                  /* VP6F */
	HI_UNF_VCODEC_TYPE_VP6A                                  /* VP6A */
	HI_UNF_VCODEC_TYPE_MJPEG                                 /* MJPEG */
	HI_UNF_VCODEC_TYPE_SORENSON                              /* SORENSON SPARK */
	HI_UNF_VCODEC_TYPE_DIVX3                                 /* DIVX3 */
	HI_UNF_VCODEC_TYPE_RAW                                   /* RAW */
	HI_UNF_VCODEC_TYPE_JPEG                                  /* JPEG, added for VENC */
	HI_UNF_VCODEC_TYPE_VP8                                   /* VP8 */
	HI_UNF_VCODEC_TYPE_MSMPEG4V1                             /* MS private MPEG4 */
	HI_UNF_VCODEC_TYPE_MSMPEG4V2
	HI_UNF_VCODEC_TYPE_MSVIDEO1 /* MS video */
	HI_UNF_VCODEC_TYPE_WMV1
	HI_UNF_VCODEC_TYPE_WMV2
	HI_UNF_VCODEC_TYPE_RV10
	HI_UNF_VCODEC_TYPE_RV20
	HI_UNF_VCODEC_TYPE_SVQ1 /* Apple video */
	HI_UNF_VCODEC_TYPE_SVQ3 /* Apple video */
	HI_UNF_VCODEC_TYPE_H261
	HI_UNF_VCODEC_TYPE_VP3
	HI_UNF_VCODEC_TYPE_VP5
	HI_UNF_VCODEC_TYPE_CINEPAK
	HI_UNF_VCODEC_TYPE_INDEO2
	HI_UNF_VCODEC_TYPE_INDEO3
	HI_UNF_VCODEC_TYPE_INDEO4
	HI_UNF_VCODEC_TYPE_INDEO5
	HI_UNF_VCODEC_TYPE_MJPEGB
	HI_UNF_VCODEC_TYPE_MVC
	HI_UNF_VCODEC_TYPE_HEVC
	HI_UNF_VCODEC_TYPE_DV
	HI_UNF_VCODEC_TYPE_VP9
	HI_UNF_VCODEC_TYPE_BUTT
)

/* Defines the type of the video frame. */
type HI_UNF_VIDEO_FRAME_TYPE_E int32

const (
	HI_UNF_FRAME_TYPE_UNKNOWN HI_UNF_VIDEO_FRAME_TYPE_E = iota /* Unknown */
	HI_UNF_FRAME_TYPE_I                                        /* I frame */
	HI_UNF_FRAME_TYPE_P                                        /* P frame */
	HI_UNF_FRAME_TYPE_B                                        /* B frame */
	HI_UNF_FRAME_TYPE_BUTT
)
