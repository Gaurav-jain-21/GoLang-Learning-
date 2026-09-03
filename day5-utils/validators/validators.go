package validators

import "strings"

func IsValidEmail(email string) bool{
	return strings.Contains(email,"@") && strings.Contains(email,".")

}
func IsStrongPassword(password string) bool{
	return len(password)>=6
}