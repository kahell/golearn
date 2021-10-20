package transaction

import (
	"golearn/campaign"
	"golearn/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	Campaign   campaign.Campaign
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
