// Code generated by "stringer -type Op -trimprefix O"; DO NOT EDIT.

package gc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OPACK-4]
	_ = x[OLITERAL-5]
	_ = x[OADD-6]
	_ = x[OSUB-7]
	_ = x[OOR-8]
	_ = x[OXOR-9]
	_ = x[OADDSTR-10]
	_ = x[OADDR-11]
	_ = x[OANDAND-12]
	_ = x[OAPPEND-13]
	_ = x[OBYTES2STR-14]
	_ = x[OBYTES2STRTMP-15]
	_ = x[ORUNES2STR-16]
	_ = x[OSTR2BYTES-17]
	_ = x[OSTR2BYTESTMP-18]
	_ = x[OSTR2RUNES-19]
	_ = x[OAS-20]
	_ = x[OAS2-21]
	_ = x[OAS2DOTTYPE-22]
	_ = x[OAS2FUNC-23]
	_ = x[OAS2MAPR-24]
	_ = x[OAS2RECV-25]
	_ = x[OASOP-26]
	_ = x[OCALL-27]
	_ = x[OCALLFUNC-28]
	_ = x[OCALLMETH-29]
	_ = x[OCALLINTER-30]
	_ = x[OCALLPART-31]
	_ = x[OCAP-32]
	_ = x[OCLOSE-33]
	_ = x[OCLOSURE-34]
	_ = x[OCOMPLIT-35]
	_ = x[OMAPLIT-36]
	_ = x[OSTRUCTLIT-37]
	_ = x[OARRAYLIT-38]
	_ = x[OSLICELIT-39]
	_ = x[OPTRLIT-40]
	_ = x[OCONV-41]
	_ = x[OCONVIFACE-42]
	_ = x[OCONVNOP-43]
	_ = x[OCOPY-44]
	_ = x[ODCL-45]
	_ = x[ODCLFUNC-46]
	_ = x[ODCLFIELD-47]
	_ = x[ODCLCONST-48]
	_ = x[ODCLTYPE-49]
	_ = x[ODELETE-50]
	_ = x[ODOT-51]
	_ = x[ODOTPTR-52]
	_ = x[ODOTMETH-53]
	_ = x[ODOTINTER-54]
	_ = x[OXDOT-55]
	_ = x[ODOTTYPE-56]
	_ = x[ODOTTYPE2-57]
	_ = x[OEQ-58]
	_ = x[ONE-59]
	_ = x[OLT-60]
	_ = x[OLE-61]
	_ = x[OGE-62]
	_ = x[OGT-63]
	_ = x[ODEREF-64]
	_ = x[OINDEX-65]
	_ = x[OINDEXMAP-66]
	_ = x[OKEY-67]
	_ = x[OSTRUCTKEY-68]
	_ = x[OLEN-69]
	_ = x[OMAKE-70]
	_ = x[OMAKECHAN-71]
	_ = x[OMAKEMAP-72]
	_ = x[OMAKESLICE-73]
	_ = x[OMUL-74]
	_ = x[ODIV-75]
	_ = x[OMOD-76]
	_ = x[OLSH-77]
	_ = x[ORSH-78]
	_ = x[OAND-79]
	_ = x[OANDNOT-80]
	_ = x[ONEW-81]
	_ = x[ONEWOBJ-82]
	_ = x[ONOT-83]
	_ = x[OBITNOT-84]
	_ = x[OPLUS-85]
	_ = x[ONEG-86]
	_ = x[OOROR-87]
	_ = x[OPANIC-88]
	_ = x[OPRINT-89]
	_ = x[OPRINTN-90]
	_ = x[OPAREN-91]
	_ = x[OSEND-92]
	_ = x[OSLICE-93]
	_ = x[OSLICEARR-94]
	_ = x[OSLICESTR-95]
	_ = x[OSLICE3-96]
	_ = x[OSLICE3ARR-97]
	_ = x[OSLICEHEADER-98]
	_ = x[ORECOVER-99]
	_ = x[ORECV-100]
	_ = x[ORUNESTR-101]
	_ = x[OSELRECV-102]
	_ = x[OSELRECV2-103]
	_ = x[OIOTA-104]
	_ = x[OREAL-105]
	_ = x[OIMAG-106]
	_ = x[OCOMPLEX-107]
	_ = x[OALIGNOF-108]
	_ = x[OOFFSETOF-109]
	_ = x[OSIZEOF-110]
	_ = x[OBLOCK-111]
	_ = x[OBREAK-112]
	_ = x[OCASE-113]
	_ = x[OCONTINUE-114]
	_ = x[ODEFER-115]
	_ = x[OEMPTY-116]
	_ = x[OFALL-117]
	_ = x[OFOR-118]
	_ = x[OFORUNTIL-119]
	_ = x[OGOTO-120]
	_ = x[OIF-121]
	_ = x[OLABEL-122]
	_ = x[OGO-123]
	_ = x[ORANGE-124]
	_ = x[ORETURN-125]
	_ = x[OSELECT-126]
	_ = x[OSWITCH-127]
	_ = x[OTYPESW-128]
	_ = x[OTCHAN-129]
	_ = x[OTMAP-130]
	_ = x[OTSTRUCT-131]
	_ = x[OTINTER-132]
	_ = x[OTFUNC-133]
	_ = x[OTARRAY-134]
	_ = x[ODDD-135]
	_ = x[ODDDARG-136]
	_ = x[OINLCALL-137]
	_ = x[OEFACE-138]
	_ = x[OITAB-139]
	_ = x[OIDATA-140]
	_ = x[OSPTR-141]
	_ = x[OCLOSUREVAR-142]
	_ = x[OCFUNC-143]
	_ = x[OCHECKNIL-144]
	_ = x[OVARDEF-145]
	_ = x[OVARKILL-146]
	_ = x[OVARLIVE-147]
	_ = x[ORESULT-148]
	_ = x[OINLMARK-149]
	_ = x[ORETJMP-150]
	_ = x[OGETG-151]
	_ = x[OEND-152]
}

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLFIELDDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMULDIVMODLSHRSHANDANDNOTNEWNEWOBJNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECVRUNESTRSELRECVSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFBLOCKBREAKCASECONTINUEDEFEREMPTYFALLFORFORUNTILGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYDDDDDDARGINLCALLEFACEITABIDATASPTRCLOSUREVARCFUNCCHECKNILVARDEFVARKILLVARLIVERESULTINLMARKRETJMPGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 36, 39, 45, 49, 55, 61, 70, 82, 91, 100, 112, 121, 123, 126, 136, 143, 150, 157, 161, 165, 173, 181, 190, 198, 201, 206, 213, 220, 226, 235, 243, 251, 257, 261, 270, 277, 281, 284, 291, 299, 307, 314, 320, 323, 329, 336, 344, 348, 355, 363, 365, 367, 369, 371, 373, 375, 380, 385, 393, 396, 405, 408, 412, 420, 427, 436, 439, 442, 445, 448, 451, 454, 460, 463, 469, 472, 478, 482, 485, 489, 494, 499, 505, 510, 514, 519, 527, 535, 541, 550, 561, 568, 572, 579, 586, 594, 598, 602, 606, 613, 620, 628, 634, 639, 644, 648, 656, 661, 666, 670, 673, 681, 685, 687, 692, 694, 699, 705, 711, 717, 723, 728, 732, 739, 745, 750, 756, 759, 765, 772, 777, 781, 786, 790, 800, 805, 813, 819, 826, 833, 839, 846, 852, 856, 859}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
