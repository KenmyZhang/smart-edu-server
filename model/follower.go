package model

type Follower struct {
	UserId      string `json:"user_id"`
	FollowId    string `json:"follow_id"`
	CreatedTime int64  `json:"created_time"`
	DeletedTime int64  `json:"created_time"`
}
