package dtos

type CreateSeatDto struct {
	Row   string  `json:"row" binding:"required,max=1"`
	Col   int     `json:"col" binding:"required,min=1"`
	Price float64 `json:"price" binding:"required,min=1"`
}
