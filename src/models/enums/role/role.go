package role

type Role string

const (
	RoleClient  Role = "client"
	RoleAdmin   Role = "admin"
	RoleChef    Role = "chef"
	RoleCashier Role = "cashier"
)
