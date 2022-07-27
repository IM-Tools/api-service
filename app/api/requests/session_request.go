package requests

type SessionStore struct {
	Id   int64 `json:"id" validate:"required"`
	Type int   `json:"type" validate:"required,gte=1,lte=2"`
}
