package storage

import "ServerSite/internal/service"

type UsersPostgresStorage struct {
}

func NewUsersPostgresStorage() *UsersPostgresStorage {
	return &UsersPostgresStorage{}
}

func (s *UsersPostgresStorage) CheckUserCredentials(login, passwordHash string) (bool, int, string, error) {
	//заглушка временная
	if login == "admin" && passwordHash == service.Sha256Hash("admin") {
		return true, 1, "admin", nil

	} else if login == "user" && passwordHash == service.Sha256Hash("user") {
		return true, 2, "user", nil

	} else {
		return false, 3, "", nil
	}
}
func (s *UsersPostgresStorage) CheckId(id int) (bool, string, error) {
	if id == 1 {
		return true, "admin", nil
	} else if id == 2 {
		return true, "user", nil
	} else {
		return false, "", nil
	}
}
