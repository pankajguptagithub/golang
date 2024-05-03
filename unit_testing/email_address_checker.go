package linkedinlearning

import "strings"

func IsLinkedInEmployee(address string)bool{
	return strings.HasSuffix(address,"@linkedin.com")
}
