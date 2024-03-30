package modal

import (
	modal "go-testing-poc/pkg/abstract"
)

type Role int

const (
	AdminRole Role = iota
	UserRole
)

type User struct {
	modal.AbstractModal
	Username string
	Email    *string
	Role     Role
}
