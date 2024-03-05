package models

import "github.com/4kayDev/admitad-integration/internal/utils/jsoner"

type Deeplink struct {
	Link string `json:"https"`
}

func (d *Deeplink) String() string {
	return jsoner.Jsonify(d)
}
