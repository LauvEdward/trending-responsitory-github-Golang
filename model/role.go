package model

type Role int

const (
	RoleAdmin Role = iota
	RoleMember
)

func (r Role) String() string {
	return [...]string{"Admin", "Member"}[r]
}
