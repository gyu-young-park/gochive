package api

type RequestPostQueryParam struct {
	Id     string `form:"id"`
	Origin string `form:"origin" binding:"required"`
	Limit  string `form:"limit"`
}
