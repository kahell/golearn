package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amoung     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
