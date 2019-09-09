package config

import (
	"fmt"
)

var (
	PORT       = 6800
	DBURL      = ""
	DBHOST     = "db"
	DBNAME     = "Movies"
	DBPASSWORD = "example"
	DBDRIVER   = "postgres"
	DBUSERNAME = "joshua"
	SECRETKEY  = "jesus"
)

func Load() {
	// var err error

	// err := godotenv.Load(filepath.Join("../", ".env"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	// if err != nil {
	// 	log.Println(err)
	// 	PORT = 6800
	// }

	// SECRETKEY = os.Getenv("SECRET_KEY")

	// DBHOST = os.Getenv("DB_HOST")
	// DBUSERNAME = os.Getenv("DB_USER")
	// DBPASSWORD = os.Getenv("DB_PASSWORD")
	// DBDRIVER = os.Getenv("DB_DRIVER")
	// DBNAME = os.Getenv("DB_NAME")
	DBURL = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DBHOST, DBUSERNAME, DBNAME, DBPASSWORD)
}
