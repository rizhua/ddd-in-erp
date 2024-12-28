package entity

import "time"

// 根实体：店仓
type Store struct {
	ID       int32
	Name     string
	Icon     string
	OrgID    int32
	Lng      float32
	Lat      float32
	Contact  string
	Tel      string
	Address  string
	OpenTime string
	StopTime string
	Type     int8
	Status   int8
	UpdateAt time.Time
	CreateAt time.Time
}
