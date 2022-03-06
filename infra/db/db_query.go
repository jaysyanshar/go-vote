package db

// users group
const (
	QueryUserInsert      = "insert into users (name, email, password) values (?, ?, ?);"
	QueryUserFindByEmail = "select id, name, email, password from users where email = ?;"
	QueryUserFindById    = "select id, name, email, password from users where password = ?;"
)
