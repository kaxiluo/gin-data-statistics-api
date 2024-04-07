package models

type UserReachMeta struct {
	ID
	Appcode string `json:"appcode"`
	Date    string `json:"date" gorm:"type:date;column:date"`
	UserId  int64  `json:"user_id"`
	Channel string `json:"channel"`
	EventId int64  `json:"event_id"`
	CreatedAt
}

func (UserReachMeta) TableName() string {
	return "user_reach_meta"
}
