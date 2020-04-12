package repository

type LoginResult struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	IsActive  bool   `db:"active"`
}

type PrivilegeResult struct {
	Code string `db:"code"`
	Name string `db:"name"`
}
