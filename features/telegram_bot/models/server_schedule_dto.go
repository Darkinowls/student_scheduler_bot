package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    serverScheduleDto, err := UnmarshalServerScheduleDto(bytes)
//    bytes, err = serverScheduleDto.Marshal()

import "encoding/json"

func UnmarshalServerScheduleDto(data []byte) (ServerScheduleDto, error) {
	var r ServerScheduleDto
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ServerScheduleDto) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerScheduleDto struct {
	Paging interface{} `json:"paging"`
	Data   ServerData  `json:"data"`
}

type ServerData struct {
	GroupCode          string               `json:"groupCode"`
	ScheduleFirstWeek  []ServerScheduleWeek `json:"scheduleFirstWeek"`
	ScheduleSecondWeek []ServerScheduleWeek `json:"scheduleSecondWeek"`
}

type ServerScheduleWeek struct {
	Day   string       `json:"day"`
	Pairs []ServerPair `json:"pairs"`
}

type ServerPair struct {
	TeacherName string     `json:"teacherName"`
	LecturerID  string     `json:"lecturerId"`
	Type        ServerType `json:"type"`
	Time        string     `json:"time"`
	Name        string     `json:"name"`
	Place       string     `json:"place"`
	Tag         ServerTag  `json:"tag"`
}

type ServerTag string

const (
	Lab  ServerTag = "lab"
	Lec  ServerTag = "lec"
	Prac ServerTag = "prac"
)

type ServerType string

const (
	ЛабOnLine  ServerType = "Лаб on-line"
	ЛекOnLine  ServerType = "Лек on-line"
	ПракOnLine ServerType = "Прак on-line"
)
