package verification

import "strings"

// تابع برای اعتبارسنجی داده‌ها
func VerificationData(FirstName, LastName, Email, NationalID string, TicketCount int) []string {
	var errors []string

	// بررسی اعتبار نام‌ها
	if len(FirstName) < 4 || len(LastName) < 4 {
		errors = append(errors, "First name and last name must be at least 4 characters.")
	}

	// بررسی شماره ملی
	if len(NationalID) != 8 {
		errors = append(errors, "National ID must be exactly 8 characters.")
	}

	// بررسی ایمیل
	if !strings.Contains(Email, "@") {
		errors = append(errors, "Email must contain '@'.")
	}

	// بررسی تعداد بلیت‌ها
	if TicketCount <= 0 {
		errors = append(errors, "Ticket count must be greater than 0.")
	}

	return errors
}
