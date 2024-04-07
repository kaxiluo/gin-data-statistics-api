package request

type UserReach struct {
	Appcode string `form:"appcode" json:"appcode" binding:"required"`
	Date    string `form:"date" json:"date" binding:"-"`
	UserId  int64  `form:"user_id" json:"user_id" binding:"required"`
	Channel string `form:"channel" json:"channel" binding:"required"`
	EventId int64  `form:"event_id" json:"event_id" binding:"required"`
}
