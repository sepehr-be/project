package data

type ReservationData struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  string `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
}

var ReservationList = make([]ReservationData, 0)
