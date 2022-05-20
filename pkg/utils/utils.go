package utils

import "os"

func GetDatabaseHost(password string) string {
	return "postgres://" + os.Getenv("DATABASE_USERNAME") + ":" + password + "@" + os.Getenv("DATABASE_HOST") + "/" + os.Getenv("DATABASE_NAME")
}
