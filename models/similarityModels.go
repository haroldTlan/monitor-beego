package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Similar struct {
	Id         int64     `orm:"column(id);pk"                         json:"id"`
	FaceId     int64     `orm:"column(face_id)"                           json:"face_id"`
	PictureId  int64     `orm:"column(picture_id)"               json:"picture_id"`
	Threshold  float64   `orm:"column(threshold)"              json:"threshold"`
	Similarity float64   `orm:"column(similarity)"                  json:"similarity"`
	Created    time.Time `orm:"column(createtime);type(datetime)"     json:"createtime"`
}

func (s *Similar) TableName() string {
	return "similar"
}

func init() {
	orm.RegisterModel(new(Similar))
}
