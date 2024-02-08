package models

import "github.com/4kayDev/admitad-integration/internal/utils/jsoner"

type PublisherRecord struct {
	LeadsSum    int     `json:"leads_sum"`
	PaymentsSum float64 `json:"payments_sum"`
	ViewsCount  int     `json:"views"`
	DeclinedSum float64 `json:"payment_sum_declined"`
	ActionsSum  int     `json:"actions_sum_total"`
	ApprovedSum float64 `json:"payment_sum_approved"`
	Currency    string  `json:"currency"`
	SalesSum    int     `json:"21694727"`
	OpenSum     float64 `json:"payment_sum_open"`
	ClicksCount int     `json:"clicks"`
}

func (r *PublisherRecord) String() string {
	return jsoner.Jsonify(r)
}
