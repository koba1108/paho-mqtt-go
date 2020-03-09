package models

import (
	"encoding/json"
	"time"
)

const TopicBaseCharger = "charger"
const TopicCharger = "charger/+/shadow/update"
const TableNameCharger = "charger"
const CollectionNameCharger = "charger"

type Charger struct {
	CsID      string    `firestore:"cs_id" bigquery:"cs_id" json:"cs_id"`
	EmV1      string    `firestore:"em_v1" bigquery:"em_v1" json:"em_v1"`
	EmV2      string    `firestore:"em_v2" bigquery:"em_v2" json:"em_v2"`
	EmV3      string    `firestore:"em_v3" bigquery:"em_v3" json:"em_v3"`
	EmI1      string    `firestore:"em_i1" bigquery:"em_i1" json:"em_i1"`
	EmI2      string    `firestore:"em_i2" bigquery:"em_i2" json:"em_i2"`
	EmI3      string    `firestore:"em_i3" bigquery:"em_i3" json:"em_i3"`
	EmP1      string    `firestore:"em_p1" bigquery:"em_p1" json:"em_p1"`
	EmP2      string    `firestore:"em_p2" bigquery:"em_p2" json:"em_p2"`
	EmP3      string    `firestore:"em_p3" bigquery:"em_p3" json:"em_p3"`
	EmE1      string    `firestore:"em_e1" bigquery:"em_e1" json:"em_e1"`
	EmE2      string    `firestore:"em_e2" bigquery:"em_e2" json:"em_e2"`
	EmE3      string    `firestore:"em_e3" bigquery:"em_e3" json:"em_e3"`
	S1Sts     string    `firestore:"s1_sts" bigquery:"s1_sts" json:"s1_sts"`
	S2Sts     string    `firestore:"s2_sts" bigquery:"s2_sts" json:"s2_sts"`
	S3Sts     string    `firestore:"s3_sts" bigquery:"s3_sts" json:"s3_sts"`
	HltSts    string    `firestore:"hlt_sts" bigquery:"hlt_sts" json:"hlt_sts"`
	EmgSts    string    `firestore:"emg_sts" bigquery:"emg_sts" json:"emg_sts"`
	BleID     string    `firestore:"ble_id" bigquery:"ble_id" json:"ble_id"`
	Year      int       `firestore:"year" bigquery:"-" json:"-"`
	Month     int       `firestore:"month" bigquery:"-" json:"-"`
	Day       int       `firestore:"day" bigquery:"-" json:"-"`
	CreatedAt time.Time `firestore:"-" bigquery:"CreatedAt" json:"-"`
	UpdatedAt time.Time `firestore:"UpdatedAt" bigquery:"-" json:"-"`
}

func NewCharger(payload []byte) *Charger {
	c := Charger{}
	_ = json.Unmarshal(payload, &c)
	now := time.Now()
	c.Year = now.Year()
	c.Month = int(now.Month())
	c.Day = now.Day()
	c.CreatedAt = now
	c.UpdatedAt = now
	return &c
}
