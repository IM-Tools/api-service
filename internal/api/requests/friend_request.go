package requests

type UpdateFriendRequest struct {
	ID     string `json:"id" validate:"required"`
	Status int    `json:"status" validate:"required,gte=1,lte=2"`
}
type CreateFriendRequest struct {
	ToId        string `json:"to_id" validate:"required"`
	Information string `json:"information" validate:"required"`
}

type QueryUserRequest struct {
	Email string `json:"email" validate:"omitempty,email"`
}
