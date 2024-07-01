package model

type NewConsumer struct {
	NIK          string  `json:"nik" binding:"required"`
	FullName     string  `json:"full_name" binding:"required"`
	LegalName    string  `json:"legal_name" binding:"required"`
	PlaceOfBirth string  `json:"place_of_birth" binding:"required"`
	DateOfBirth  string  `json:"date_of_birth" binding:"required"`
	Salary       float64 `json:"salary" binding:"required"`
	IdCardPhoto  string  `json:"id_card_photo" binding:"required"`
	SelfiePhoto  string  `json:"selfie_photo" binding:"required"`
}
