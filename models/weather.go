package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Weather struct {
	uadmin.Model
	Date        time.Time
	Temperature float64
	Status      string // > 28 sunny else possiblity of rain
}

func (w Weather) String() string {
	return w.Status
}

func (w *Weather) Save() {
	if w.Temperature > 28.0 {
		w.Status = "Sunny"
	} else {
		w.Status = "Possibility of Rain"
	}
	uadmin.Save(w)
}
