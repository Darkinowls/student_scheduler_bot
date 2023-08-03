// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    scheduleEntity, err := UnmarshalScheduleEntity(bytes)
//    bytes, err = scheduleEntity.Marshal()

package models

import "encoding/json"

func UnmarshalScheduleEntity(data []byte) (ScheduleEntity, error) {
	var r ScheduleEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ScheduleEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ScheduleEntity struct {
	Pairs []PairEntity `json:"pairs"`
}

type PairEntity struct {
	ChatID int64  `json:"chat_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}
