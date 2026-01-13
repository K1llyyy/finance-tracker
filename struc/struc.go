package struc

import "time"

type Struc struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreateADT   time.Time `json:"createadt"`
	UpdateADT   time.Time `json:"updateadt"`
}
