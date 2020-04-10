package repository

const (
	insertNewUser = `INSERT INTO common.user(
							email,
							phone_number,
							password ,
							first_name,
							last_name,
							created_by)
					VALUES(	:email,
							:phone_number,
							:password ,
							:first_name,
							:last_name,
							:created_by)
					RETURNING id`

	userByEmail = `SELECT id, first_name, last_name FROM common.user WHERE e_mail = $1 AND password = $2`

	userByID = `SELECT e_mail, first_name, last_name FROM common.user WHERE id = $1 AND password = $2`

	changePassword = `UPDATE common.user SET  password = $2 WHERE id = $1`
)
