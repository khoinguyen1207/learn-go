package request

type GetUserByIdParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}
