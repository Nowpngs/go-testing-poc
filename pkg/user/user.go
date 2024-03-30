package modal

import (
	modal "go-testing-poc/pkg/abstract"
) 

type User struct {
	modal.AbstractModal
	Username string
	Email *string
	
}