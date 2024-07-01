package model

type NewTransaction struct {
	ConsumerID int     `json:"consumer_id" binding:"required"`
	LimitID    int     `json:"limit_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	AssetName  string  `json:"asset_name" binding:"required"`
}
