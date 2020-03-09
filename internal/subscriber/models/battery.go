package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"time"
)

const TopicBaseBattery = "battery"
const TopicBattery = "battery/+/shadow/update"
const TableNameBattery = "battery"
const CollectionNameBattery = "battery"

type Battery struct {
	TIME      string    `firestore:"TIME" bigquery:"TIME" json:"TIME"`
	TID       string    `firestore:"TID" bigquery:"TID" json:"TID"`
	BV        float64   `firestore:"BV" bigquery:"BV" json:"BV"`
	BI        float64   `firestore:"BI" bigquery:"BI" json:"BI"`
	SOC       int       `firestore:"SOC" bigquery:"SOC" json:"SOC"`
	TS        int       `firestore:"TS" bigquery:"TS" json:"TS"`
	BV1       float64   `firestore:"BV1" bigquery:"BV1" json:"BV1"`
	BI1       float64   `firestore:"BI1" bigquery:"BI1" json:"BI1"`
	BSC1      int       `firestore:"BSC1" bigquery:"BSC1" json:"BSC1"`
	TB1       int       `firestore:"TB1" bigquery:"TB1" json:"TB1"`
	CAB1      int       `firestore:"CAB1" bigquery:"CAB1" json:"CAB1"`
	CKB1      int       `firestore:"CKB1" bigquery:"CKB1" json:"CKB1"`
	DAB1      int       `firestore:"DAB1" bigquery:"DAB1" json:"DAB1"`
	DKB1      int       `firestore:"DKB1" bigquery:"DKB1" json:"DKB1"`
	SHB1      int       `firestore:"SHB1" bigquery:"SHB1" json:"SHB1"`
	CCB1      int       `firestore:"CCB1" bigquery:"CCB1" json:"CCB1"`
	CHR1      float64   `firestore:"CHR1" bigquery:"CHR1" json:"CHR1"`
	DHR1      float64   `firestore:"DHR1" bigquery:"DHR1" json:"DHR1"`
	BV2       float64   `firestore:"BV2" bigquery:"BV2" json:"BV2"`
	BI2       float64   `firestore:"BI2" bigquery:"BI2" json:"BI2"`
	BSC2      int       `firestore:"BSC2" bigquery:"BSC2" json:"BSC2"`
	TB2       int       `firestore:"TB2" bigquery:"TB2" json:"TB2"`
	CAB2      int       `firestore:"CAB2" bigquery:"CAB2" json:"CAB2"`
	CKB2      int       `firestore:"CKB2" bigquery:"CKB2" json:"CKB2"`
	DAB2      int       `firestore:"DAB2" bigquery:"DAB2" json:"DAB2"`
	DKB2      int       `firestore:"DKB2" bigquery:"DKB2" json:"DKB2"`
	SHB2      int       `firestore:"SHB2" bigquery:"SHB2" json:"SHB2"`
	CCB2      int       `firestore:"CCB2" bigquery:"CCB2" json:"CCB2"`
	CHR2      float64   `firestore:"CHR2" bigquery:"CHR2" json:"CHR2"`
	DHR2      float64   `firestore:"DHR2" bigquery:"DHR2" json:"DHR2"`
	BV3       float64   `firestore:"BV3" bigquery:"BV3" json:"BV3"`
	BI3       float64   `firestore:"BI3" bigquery:"BI3" json:"BI3"`
	BSC3      int       `firestore:"BSC3" bigquery:"BSC3" json:"BSC3"`
	TB3       int       `firestore:"TB3" bigquery:"TB3" json:"TB3"`
	CAB3      int       `firestore:"CAB3" bigquery:"CAB3" json:"CAB3"`
	CKB3      int       `firestore:"CKB3" bigquery:"CKB3" json:"CKB3"`
	DAB3      int       `firestore:"DAB3" bigquery:"DAB3" json:"DAB3"`
	DKB3      int       `firestore:"DKB3" bigquery:"DKB3" json:"DKB3"`
	SHB3      int       `firestore:"SHB3" bigquery:"SHB3" json:"SHB3"`
	CCB3      int       `firestore:"CCB3" bigquery:"CCB3" json:"CCB3"`
	CHR3      float64   `firestore:"CHR3" bigquery:"CHR3" json:"CHR3"`
	DHR3      float64   `firestore:"DHR3" bigquery:"DHR3" json:"DHR3"`
	TDT       float64   `firestore:"TDT" bigquery:"TDT" json:"TDT"`
	SPD       float64   `firestore:"SPD" bigquery:"SPD" json:"SPD"`
	LNG       string    `firestore:"LNG" bigquery:"LNG" json:"LNG"`
	LNS       string    `firestore:"LNS" bigquery:"LNS" json:"LNS"`
	LAT       string    `firestore:"LAT" bigquery:"LAT" json:"LAT"`
	LAS       string    `firestore:"LAS" bigquery:"LAS" json:"LAS"`
	AL1       int       `firestore:"AL1" bigquery:"AL1" json:"AL1"`
	AL2       int       `firestore:"AL2" bigquery:"AL2" json:"AL2"`
	AL3       int       `firestore:"AL3" bigquery:"AL3" json:"AL3"`
	AL4       int       `firestore:"AL4" bigquery:"AL4" json:"AL4"`
	ST1       int       `firestore:"ST1" bigquery:"ST1" json:"ST1"`
	ST2       int       `firestore:"ST2" bigquery:"ST2" json:"ST2"`
	EC        int       `firestore:"EC" bigquery:"EC" json:"EC"`
	SNB1      string    `firestore:"SNB1" bigquery:"SNB1" json:"SNB1"`
	SNB2      string    `firestore:"SNB2" bigquery:"SNB2" json:"SNB2"`
	BM        int       `firestore:"BM" bigquery:"BM" json:"BM"`
	FV        float64   `firestore:"FV" bigquery:"FV" json:"FV"`
	B1CV      float64   `firestore:"B1CV" bigquery:"B1CV" json:"B1CV"`
	B2CV      float64   `firestore:"B2CV" bigquery:"B2CV" json:"B2CV"`
	B3CV      float64   `firestore:"B3CV" bigquery:"B3CV" json:"B3CV"`
	Year      int       `firestore:"year" bigquery:"-" json:"-"`
	Month     int       `firestore:"month" bigquery:"-" json:"-"`
	Day       int       `firestore:"day" bigquery:"-" json:"-"`
	CreatedAt time.Time `firestore:"-" bigquery:"CreatedAt" json:"-"`
	UpdatedAt time.Time `firestore:"UpdatedAt" bigquery:"-" json:"-"`
}

func NewBattery(payload []byte) *Battery {
	b := Battery{}
	_ = json.Unmarshal(payload, &b)
	now := time.Now()
	b.Year = now.Year()
	b.Month = int(now.Month())
	b.Day = now.Day()
	b.CreatedAt = now
	b.UpdatedAt = now
	return &b
}

func (t *Battery) FirestoreUpdate(ctx context.Context, db *firestore.Client, docId string) (*firestore.WriteResult, error) {
	return db.Collection(CollectionNameBattery).Doc(docId).Set(ctx, t)
}
