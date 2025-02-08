package verification

import (
	"net/mail"
	"strconv"
	"strings"
)

type correctData struct {
	FirstName string
	LastName string
}

func Trim(str string) string {
	trimedStr := strings.TrimSpace(str)
	return trimedStr
}

func ChantainsSpecialChars(str string) bool {
	return strings.ContainsAny(str, "@#$%^&*_-+'`")
}

func VerificationData(FirstName, LastName, Email string, NationalID, TicketCount, remainingTicket int) ([]string,correctData) {
	var errors []string
	var Name correctData
	num := strconv.Itoa(NationalID)

	if len(FirstName) < 2 {
		errors = append(errors, "First name must be at least 2 characters.")
	}

	specialChars := ChantainsSpecialChars(FirstName)
	if specialChars {
		errors = append(errors, "please enter correct First name")
	}

	trimedName := Trim(FirstName)
	Name.FirstName = trimedName

	specialChars = ChantainsSpecialChars(LastName)
	if specialChars {
		errors = append(errors, "please enter correct Last name")
	}

	if len(LastName) < 3 {
		errors = append(errors, "last name must be at least 3 characters.")
	}

	trimedLastName := Trim(LastName)
	Name.LastName = trimedLastName

	if len(num) != 8 && num == "00000000" {
		errors = append(errors, "National ID must be exactly 8 characters.")
	}

	if _, err := mail.ParseAddress(Email); err != nil {
		errors = append(errors, "Your Email Is Incorrect")
	}

	if TicketCount <= 0 {
		errors = append(errors, "Ticket count must be greater than 0.")
	}

	if TicketCount > remainingTicket {
		errors = append(errors, "Ticket count must be less than Remaining Ticket.")
	}


	return errors,Name
}
