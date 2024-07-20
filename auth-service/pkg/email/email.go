package email

import (
	"fmt"
)

func SendVerificationEmail(to, code string) error {
	// In a real application, integrate with an email service provider here
	fmt.Printf("Sending verification email to %s with code %s\n", to, code)
	return nil
}
