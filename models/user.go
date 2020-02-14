package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

type Role int

const (
	RoleAdministrator Role = iota
	RoleBasicUser
)

func init() {}
