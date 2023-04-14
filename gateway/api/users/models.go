package users

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}


