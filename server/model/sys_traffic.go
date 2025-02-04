package model

import "time"

// sspanel 上报用户的流量使用情况
type TrafficInfo struct {
	UserID int64 `json:"user_id" gorm:"column:user_id"`
	U      int64 `json:"u" gorm:"column:u"` //bit
	D      int64 `json:"d" gorm:"column:d"` //bit
}

// sspanel 上报用户的流量使用情况
type TrafficReq struct {
	Data []TrafficInfo
}

// 数据库 traffic_log 流量统计表
type TrafficLog struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"      gorm:"primary_key"`
	NodeID    int64     `json:"node_id" gorm:"comment:节点ID"`
	U         int64     `json:"u"       gorm:"comment:上行流量 bit"`
	D         int64     `json:"d"       gorm:"comment:下行流量 bit"`
}
