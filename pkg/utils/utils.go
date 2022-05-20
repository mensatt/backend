package utils

import "os"

func GetDatabaseHost(username string, password string) string {
	return "postgres://" + username + ":" + password + "@" + os.Getenv("DATABASE_HOST") + "/" + os.Getenv("DATABASE_NAME")
}
