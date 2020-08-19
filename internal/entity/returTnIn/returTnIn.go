package returtnin

import (
	"time"
)

// TranRCD ...
type TranRCD struct {
	TranrcDRunningID     int       `db:"TranrcD_RunningID" json:"TranrcD_RunningID"`
	TranrcDOutCodeTransf string    `db:"TranrcD_OutCodeTransf" json:"TranrcD_OutCodeTransf"`
	TranrcDNoTransf      string    `db:"TranrcD_NoTransf" json:"TranrcD_NoTransf"`
	TranrcDOutCodeTranrc string    `db:"TranrcD_OutCodeTranrc" json:"TranrcD_OutCodeTranrc"`
	TranrcDNoTranrc      string    `db:"TranrcD_NoTranrc" json:"TranrcD_NoTranrc"`
	TranrcDProcod        string    `db:"TranrcD_Procod" json:"TranrcD_Procod"`
	TranrcDBBatchNumber  string    `db:"TranrcDB_BatchNumber" json:"TranrcDB_BatchNumber"`
	TranrcDBKonsentrasi  int       `db:"TranrcDB_Konsentrasi" json:"TranrcDB_Konsentrasi"`
	TranrcDQuantityRecv  int       `db:"TranrcD_QuantityRecv" json:"TranrcD_QuantityRecv"`
	TranrcDQuantityScan  int       `db:"TranrcD_QuantityScan" json:"TranrcD_QuantityScan"`
	TranrcDQuantityStk   int       `db:"TranrcD_QuantityStk" json:"TranrcD_QuantityStk"`
	TranrcDOutCodeOR     string    `db:"TranrcD_OutCodeOR" json:"TranrcD_OutCodeOR"`
	TranrcDNoOR          string    `db:"TranrcD_NoOR" json:"TranrcD_NoOR"`
	TranrcDQuantityOR    string    `db:"TranrcD_QuantityOR" json:"TranrcD_QuantityOR"`
	TranrcDActiveYN      string    `db:"TranrcD_ActiveYN" json:"TranrcD_ActiveYN"`
	TranrcDUserID        string    `db:"TranrcD_UserId" json:"TranrcD_UserId"`
	TranrcDlastUpdate    time.Time `db:"TranrcD_lastUpdate" json:"TranrcD_lastUpdate"`
	TranrcDDataAktifYN   string    `db:"TranrcD_DataAktifYN" json:"TranrcD_DataAktifYN"`
}

// TranRCH ...
type TranRCH struct {
	TranrcHRunningID     int       `db:"TranrcH_RunningID" json:"TranrcH_RunningID"`
	TranrcHOutCodeTransf string    `db:"TranrcH_OutCodeTransf" json:"TranrcH_OutCodeTransf"`
	TranrcHNoTransf      string    `db:"TranrcH_NoTransf" json:"TranrcH_NoTransf"`
	TranrcHOutCodeTranrc string    `db:"TranrcH_OutCodeTranrc" json:"TranrcH_OutCodeTranrc"`
	TranrcHNoTranrc      string    `db:"TranrcH_NoTranrc" json:"TranrcH_NoTranrc"`
	TranrcHTglTranrc     time.Time `db:"TranrcH_TglTranrc" json:"TranrcH_TglTranrc"`
	TranrcHFlag          string    `db:"TranrcH_Flag" json:"TranrcH_Flag"`
	TranrcHFlagTrf       string    `db:"TranrcH_FlagTrf" json:"TranrcH_FlagTrf"`
	TranrcHTglDwld       time.Time `db:"TranrcH_TglDwld" json:"TranrcH_TglDwld"`
	TranrcHNip           string    `db:"TranrcH_Nip" json:"TranrcH_Nip"`
	TranrcHActiveYN      string    `db:"TranrcH_ActiveYN" json:"TranrcH_ActiveYN"`
	TranrcHUserID        string    `db:"TranrcH_UserId" json:"TranrcH_UserId"`
	TranrcHlastUpdate    time.Time `db:"TranrcH_lastUpdate" json:"TranrcH_lastUpdate"`
	TranrcHDataAktifYN   string    `db:"TranrcH_DataAktifYN" json:"TranrcH_DataAktifYN"`
}

// DetailRC ...
type DetailRC struct {
	KodeProduct      string `db:"TranrcD_Procod" json:"KodeProduct"`
	DeskripsiProduct string `json:"DeskripsiProduct"`
	QuantityScan     int    `db:"TranrcD_QuantityScan" json:"QuantityScan"`
	SellPack         int    `json:"SellPack"`
	QuantityDO       int    `db:"TranrcD_QuantityRecv" json:"QuantityDO"`
	QuantityStok     string `db:"TranrcD_QuantityStk" json:"QuantityStok"`
	ExpDate          string `json:"ExpDate"`
	BatchNumber      string `db:"TranrcDB_BatchNumber" json:"BatchNumber"`
}

// HeaderRC ...
type HeaderRC struct {
	NoReceive      string    `db:"TranrcH_NoTranrc" json:"no_receive"`
	KodePengirim   string    `db:"TranrcH_OutCodeTransf" json:"kode_pengirim"`
	Pengirim       string    `db:"TranrcH_Nip" json:"Pengirim"`
	NomorTransfer  string    `db:"TranrcH_NoTransf" json:"nomor_transfer"`
	StatusPrint    string    `db:"TranrcH_Flag" json:"status_print"`
	TanggalReceive time.Time `db:"TranrcH_TglTranrc" json:"tanggal_receive"`
}

// JSONRC ...
type JSONRC struct {
	TranRCH HeaderRC
	TranRCD []DetailRC
}
