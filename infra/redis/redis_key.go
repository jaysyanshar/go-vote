package redis

import "fmt"

const (
	keyUser = "user"
)

func GetKeyUserFindByEmail(email string) string {
	key := fmt.Sprintf("%s:email:%s", keyUser, email)
	return key
}

func GetKeyUserFindById(id int64) string {
	key := fmt.Sprintf("%s:id:%d", keyUser, id)
	return key
}
