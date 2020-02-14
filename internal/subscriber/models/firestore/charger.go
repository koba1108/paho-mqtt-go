package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
)

const CollectionNameCharger = "charger"

type Charger struct {
	CsID   string `json:"cs_id"`
	EmV1   string `json:"em_v1"`
	EmV2   string `json:"em_v2"`
	EmV3   string `json:"em_v3"`
	EmI1   string `json:"em_i1"`
	EmI2   string `json:"em_i2"`
	EmI3   string `json:"em_i3"`
	EmP1   string `json:"em_p1"`
	EmP2   string `json:"em_p2"`
	EmP3   string `json:"em_p3"`
	EmE1   string `json:"em_e1"`
	EmE2   string `json:"em_e2"`
	EmE3   string `json:"em_e3"`
	S1Sts  string `json:"s1_sts"`
	S2Sts  string `json:"s2_sts"`
	S3Sts  string `json:"s3_sts"`
	HltSts string `json:"hlt_sts"`
	EmgSts string `json:"emg_sts"`
	BleID  string `json:"ble_id"`
}

func (t *Charger) Update(ctx context.Context, db *firestore.Client, docId string) (*firestore.WriteResult, error) {
	// todo: CreateAtを追加してインドのタイムゾーンのtimestamp型にする
	return db.Collection(CollectionNameCharger).Doc(docId).Set(ctx, t)
}
