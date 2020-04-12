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
							:phonenumber,
							:password ,
							:firstname,
							:lastname,
							:createdby)
					RETURNING id`

	insertPrivilege = `INSERT INTO common.privilege(
							user_id,
							resource_id,
							resource_type_id,
							role_id)
					VALUES(	:userid,
							:resourceid,
							:resourcetypeid,
							:roleid)`

	userByEmail = `SELECT id, password, first_name, last_name FROM common.user WHERE email = $1`

	userByID = `SELECT email, first_name, last_name, active FROM common.user WHERE id = $1`

	changePassword = `UPDATE common.user SET  password = $2 WHERE id = $1`

	userPrivileges = `SELECT r.code, r.name FROM common.role r
						INNER JOIN common.privilege p ON r.id = p.role_id
						WHERE p.user_id::TEXT = $1 and p.resource_id::TEXT = $2`


)
