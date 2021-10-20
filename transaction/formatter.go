package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"current_amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{}
	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt
	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	transactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, campaignTransactionFormatter)
	}

	return transactionsFormatter
}
