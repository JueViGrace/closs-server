package types

import (
	"fmt"
)

type Role string

func (r Role) String() string {
	return string(r)
}

func RoleFromString(s string) (Role, error) {
	switch {
	case s == RoleCustomer.String():
		return RoleCustomer, nil
	case s == RoleSalesman.String():
		return RoleSalesman, nil
	case s == RoleManager.String():
		return RoleManager, nil
	case s == RoleAdmin.String():
		return RoleAdmin, nil
	default:
		return "", fmt.Errorf("cannot match %s to Role\n", s)
	}
}

const (
	RoleCustomer Role = "customer"
	RoleSalesman Role = "salesman"
	RoleManager  Role = "manager"
	RoleAdmin    Role = "admin"
)
