package modal

import (
	modal "go-testing-poc/pkg/abstract"
)

type Role int

const (
	UserRole Role = iota + 1
	AdminRole
)

type User struct {
	modal.AbstractModal
	Username string
	Email    *string
	Role     Role
}
