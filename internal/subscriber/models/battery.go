package models

const TopicBaseBattery = "battery"
const TopicBattery = "battery/+/shadow/update"

type Battery struct {
	TIME string  `json:"TIME"`
	TID  string  `json:"TID"`
	BV   float64 `json:"BV"`
	BI   float64 `json:"BI"`
	SOC  int     `json:"SOC"`
	TS   int     `json:"TS"`
	BV1  float64 `json:"BV1"`
	BI1  float64 `json:"BI1"`
	BSC1 int     `json:"BSC1"`
	TB1  int     `json:"TB1"`
	CAB1 int     `json:"CAB1"`
	CKB1 int     `json:"CKB1"`
	DAB1 int     `json:"DAB1"`
	DKB1 int     `json:"DKB1"`
	SHB1 int     `json:"SHB1"`
	CCB1 int     `json:"CCB1"`
	CHR1 float64 `json:"CHR1"`
	DHR1 float64 `json:"DHR1"`
	BV2  float64 `json:"BV2"`
	BI2  float64 `json:"BI2"`
	BSC2 int     `json:"BSC2"`
	TB2  int     `json:"TB2"`
	CAB2 int     `json:"CAB2"`
	CKB2 int     `json:"CKB2"`
	DAB2 int     `json:"DAB2"`
	DKB2 int     `json:"DKB2"`
	SHB2 int     `json:"SHB2"`
	CCB2 int     `json:"CCB2"`
	CHR2 float64 `json:"CHR2"`
	DHR2 float64 `json:"DHR2"`
	TDT  float64 `json:"TDT"`
	SPD  float64 `json:"SPD"`
	LNG  string  `json:"LNG"`
	LNS  string  `json:"LNS"`
	LAT  string  `json:"LAT"`
	LAS  string  `json:"LAS"`
	AL1  int     `json:"AL1"`
	AL2  int     `json:"AL2"`
	AL4  int     `json:"AL4"`
	ST1  int     `json:"ST1"`
	ST2  int     `json:"ST2"`
	EC   int     `json:"EC"`
	SNB1 string  `json:"SNB1"`
	SNB2 string  `json:"SNB2"`
	BM   int     `json:"BM"`
	FV   float64 `json:"FV"`
}

func (t *Battery) Insert() error {
	// todo: insert to data store
	return nil
}
