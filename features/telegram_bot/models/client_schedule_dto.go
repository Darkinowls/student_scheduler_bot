package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    scheduleDTO, err := UnmarshalScheduleDTO(bytes)
//    bytes, err = scheduleDTO.Marshal()

import (
	"encoding/json"
	"errors"
)

func UnmarshalClientScheduleDTO(data []byte) (*ClientScheduleDTO, error) {
	var r *ClientScheduleDTO
	err := json.Unmarshal(data, &r)
	if err != nil {
		err = errors.New("Json syntaxis error")
	}
	return r, err
}

func (r ClientScheduleDTO) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ClientScheduleDTO struct {
	ScheduleFirstWeek  []ClientScheduleWeek `json:"scheduleFirstWeek"`
	ScheduleSecondWeek []ClientScheduleWeek `json:"scheduleSecondWeek"`
}

type ClientScheduleWeek struct {
	Day   string       `json:"day"`
	Pairs []ClientPair `json:"pairs"`
}

type ClientPair struct {
	Time  string  `json:"time"`
	Name  string  `json:"name"`
	Place *string `json:"place"`
}
