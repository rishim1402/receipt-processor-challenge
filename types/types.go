package types

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required" validate:"required"`
	Price            string `json:"price" binding:"required,numeric" validate:"required,numeric"`
}

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required" validate:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required,datetime=2006-01-02" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" binding:"required,datetime=15:04" validate:"required,datetime=15:04"`
	Items        []Item `json:"items" binding:"required,dive,required" validate:"required,min=1,dive,required"`
	Total        string `json:"total" binding:"required,numeric" validate:"required,numeric"`
}
