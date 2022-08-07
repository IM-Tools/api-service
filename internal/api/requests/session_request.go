package requests

type SessionStore struct {
	Id   int64 `json:"id" validate:"required"`
	Type int   `json:"type" validate:"required,gte=1,lte=2"`
}

type SessionUpdate struct {
	TopStatus int    `json:"top_status" validate:"required,gte=0,lte=1"`
	Note      string `json:"type"`
}
