package dtos

type SignUpDto struct {
	Email    string `json:"email"    binding:"required,email"          example:"joedoe@gmail.com"`
	Username string `json:"username" binding:"required,alphanum,min=3" example:"coolusername123"`
	Password string `json:"password" binding:"required,min=8,max=64"   example:"S3cUrEP@55w0rd"`
}

type SignUpResponse struct {
	Code   int    `json:"code"   example:"201"`
	Status string `json:"status" example:"success"`
	Data   string `json:"data"   example:"jwt token"`
}
