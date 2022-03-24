package db

// users group
const (
	QueryUserInsert      = "insert into users (name, email, password) values (?, ?, ?);"
	QueryUserFindByEmail = "select id, name, email, password from users where email = ?;"
	QueryUserFindById    = "select id, name, email, password from users where id = ?;"
)

// sessions group
const (
	QueryAuthInsert        = "insert into sessions (userId, ipAddress, expiredAt) values (?, ?, ?);"
	QueryAuthFindById      = "select s.id, u.id, u.email, u.name, s.ipAddress, s.createdAt, s.expiredAt, s.isRevoked from sessions s join users u on u.id = s.userId where s.id = ?;"
	QueryAuthUpdateExpired = "update sessions set expiredAt = ? where id = ?;"
	QueryAuthRevokeById    = "update sessions set isRevoked = true where id = ?;"
)
