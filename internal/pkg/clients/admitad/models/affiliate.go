package models

import "github.com/4kayDev/admitad-integration/internal/utils/jsoner"

type Affiliate struct {
	Id                        int        `json:"id"`
	Name                      string     `json:"name"`
	Status                    string     `json:"status"`
	Rating                    string     `json:"rating"`
	ImageURL                  string     `json:"image"`
	Description               string     `json:"description"`
	Traffics                  []Traffic  `json:"traffics"`
	Actions                   []Action   `json:"actions"`
	SiteURL                   string     `json:"site_url"`
	Regions                   []Region   `json:"regions"`
	Currency                  string     `json:"currency"`
	Geotargeting              bool       `json:"geotargeting"`
	IsConnected               bool       `json:"connected"`
	CR                        float32    `json:"cr"`
	ECPC                      float32    `json:"ecpc"`
	EPC                       float32    `json:"epc"`
	CrTrend                   float32    `json:"cr_trend"`
	Categories                []Category `json:"categories"`
	ActionType                string     `json:"action_type"`
	IsIndividualTerms         bool       `json:"individual_terms"`
	IsAllowDeeplink           bool       `json:"allow_deeplink"`
	ActionTestLimit           int        `json:"action_testing_limit"`
	MobileDevice              string     `json:"mobile_device_type"`
	MobileOS                  string     `json:"mobile_os_type"`
	ActionCountries           []string   `json:"action_countries"`
	IsAllowAllCountriesAction bool       `json:"allow_actions_all_countries"`
}

func (a *Affiliate) String() string {
	return jsoner.Jsonify(a)
}

type Traffic struct {
	IsEnabled bool   `json:"enabled"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Id        int    `json:"id"`
}

func (t *Traffic) String() string {
	return jsoner.Jsonify(t)
}

type Action struct {
	PaymentSize string `json:"payment_size"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Id          int    `json:"id"`
}

func (a *Action) String() string {
	return jsoner.Jsonify(a)
}

type Region struct {
	Region string `json:"region"`
}

func (r *Region) String() string {
	return jsoner.Jsonify(r)
}

type Category struct {
	Language string    `json:"language"`
	Name     string    `json:"name"`
	Parent   *Category `json:"parent"`
	Id       int       `json:"id"`
}

func (c *Category) String() string {
	return jsoner.Jsonify(c)
}
