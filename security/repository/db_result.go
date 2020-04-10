package repository

type LoginResult struct {
	ID        string `db:"id"`
	Email     string `db:"e_mail"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
