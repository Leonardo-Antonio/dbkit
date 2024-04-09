# DB KIT

- DBKit: Facilita el desarrollo de aplicaciones que interactúan con bases de datos en Go. Simplifica consultas y manipulación de datos, permitiendo código limpio y eficiente, y reduciendo el tiempo de desarrollo. Mejora la mantenibilidad del código

## Install
```bash
go get github.com/Leonardo-Antonio/dbkit
```

## Use

#### Insert Item
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/dbkit"
)

type User struct {
	UserId             uint32 `json:"userId" db:"userId"`
	ProductId          string `json:"productId" db:"productId"`
	Email              string `json:"email" db:"email"`
	IcCode             string `json:"icCode" db:"icCode"`
	Phone              string `json:"phone" db:"phone"`
	Agent              string `json:"agent" db:"agent"`
	Device             string `json:"device" db:"device"`
	SocialNetwork      string `json:"socialNetwork" db:"socialNetwork"`
	Username           string `json:"username" db:"username"`
	Password           string `json:"password" db:"password"`
	FirstName          string `json:"firstName" db:"firstName"`
	LastName           string `json:"lastName" db:"lastName"`
	Gender             uint8  `json:"gender" db:"gender"`
	BirthDate          string `json:"birthDate" db:"birthDate"`
	BirthYear          uint16 `json:"birthYear" db:"birthYear"`
	BirthMonth         uint8  `json:"birthMonth" db:"birthMonth"`
	CountryId          string `json:"countryId" db:"countryId"`
	Region             string `json:"region" db:"region"`
	PostalCode         string `json:"postalCode" db:"postalCode"`
	Address            string `json:"address" db:"address"`
	AddressLine1       string `json:"addressLine1" db:"addressLine1"`
	AddressLine2       string `json:"addressLine2" db:"addressLine2"`
	Verified           uint8  `json:"verified" db:"verified"`
	Active             uint8  `json:"active" db:"active"`
	Ban                uint8  `json:"ban" db:"ban"`
	Deleted            uint8  `json:"deleted" db:"deleted"`
	FailedLogins       uint8  `json:"failedLogins" db:"failedLogins"`
	RegIp              string `json:"regIp" db:"regIp"`
	RegDate            string `json:"regDate" db:"regDate"`
	RegDatetime        string `json:"regDatetime" db:"regDatetime"`
	RegTimestamp       int64  `json:"regTimestamp" db:"regTimestamp"`
	UpdDate            string `json:"updDate" db:"updDate"`
	UpdDatetime        string `json:"updDatetime" db:"updDatetime"`
	UpdTimestamp       int64  `json:"updTimestamp" db:"updTimestamp"`
	LastLoginDate      string `json:"lastLoginDate" db:"lastLoginDate"`
	LastLoginDatetime  string `json:"lastLoginDatetime" db:"lastLoginDatetime"`
	LastLoginTimestamp int64  `json:"lastLoginTimestamp" db:"lastLoginTimestamp"`
}

func main() {
	db, err := sql.Open("mysql", "your_url_conn")
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		UserId:             1,
		ProductId:          "mangox",
		Email:              "leo@gmail.com",
		IcCode:             "+51",
		Phone:              "1234567890",
		Agent:              "web",
		Device:             "desktop",
		SocialNetwork:      "facebook",
		Username:           "username123",
		Password:           "password123",
		FirstName:          "John",
		LastName:           "Doe",
		Gender:             1,
		BirthDate:          "1990-01-01",
		BirthYear:          1990,
		BirthMonth:         1,
		CountryId:          "PER",
		Region:             "Lima",
		PostalCode:         "LIM01",
		Address:            "123 Main St",
		AddressLine1:       "Apt 4B",
		AddressLine2:       "",
		Verified:           1,
		Active:             1,
		Ban:                0,
		Deleted:            0,
		FailedLogins:       0,
		RegIp:              "192.168.1.1",
		RegDate:            "2024-04-08",
		RegDatetime:        "2024-04-08T18:39:08",
		RegTimestamp:       1678007948,
		UpdDate:            "2024-04-08",
		UpdDatetime:        "2024-04-08T18:39:08",
		UpdTimestamp:       1678007948,
		LastLoginDate:      "2024-04-08",
		LastLoginDatetime:  "2024-04-08T18:39:08",
		LastLoginTimestamp: 1678007948,
	}

	kit := dbkit.New(dbkit.MYSQL, db)
	result, err := kit.NewItem("UsersDB.Users", user)
	fmt.Println(result, err) // out: result (LastInsertId, RowsAffected), error
}
```


#### Get Items
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/dbkit"
)

type User struct {
	UserId             uint32 `json:"userId" db:"userId"`
	ProductId          string `json:"productId" db:"productId"`
	Email              string `json:"email" db:"email"`
	IcCode             string `json:"icCode" db:"icCode"`
	Phone              string `json:"phone" db:"phone"`
	Agent              string `json:"agent" db:"agent"`
	Device             string `json:"device" db:"device"`
	SocialNetwork      string `json:"socialNetwork" db:"socialNetwork"`
	Username           string `json:"username" db:"username"`
	Password           string `json:"password" db:"password"`
	FirstName          string `json:"firstName" db:"firstName"`
	LastName           string `json:"lastName" db:"lastName"`
	Gender             uint8  `json:"gender" db:"gender"`
	BirthDate          string `json:"birthDate" db:"birthDate"`
	BirthYear          uint16 `json:"birthYear" db:"birthYear"`
	BirthMonth         uint8  `json:"birthMonth" db:"birthMonth"`
	CountryId          string `json:"countryId" db:"countryId"`
	Region             string `json:"region" db:"region"`
	PostalCode         string `json:"postalCode" db:"postalCode"`
	Address            string `json:"address" db:"address"`
	AddressLine1       string `json:"addressLine1" db:"addressLine1"`
	AddressLine2       string `json:"addressLine2" db:"addressLine2"`
	Verified           uint8  `json:"verified" db:"verified"`
	Active             uint8  `json:"active" db:"active"`
	Ban                uint8  `json:"ban" db:"ban"`
	Deleted            uint8  `json:"deleted" db:"deleted"`
	FailedLogins       uint8  `json:"failedLogins" db:"failedLogins"`
	RegIp              string `json:"regIp" db:"regIp"`
	RegDate            string `json:"regDate" db:"regDate"`
	RegDatetime        string `json:"regDatetime" db:"regDatetime"`
	RegTimestamp       int64  `json:"regTimestamp" db:"regTimestamp"`
	UpdDate            string `json:"updDate" db:"updDate"`
	UpdDatetime        string `json:"updDatetime" db:"updDatetime"`
	UpdTimestamp       int64  `json:"updTimestamp" db:"updTimestamp"`
	LastLoginDate      string `json:"lastLoginDate" db:"lastLoginDate"`
	LastLoginDatetime  string `json:"lastLoginDatetime" db:"lastLoginDatetime"`
	LastLoginTimestamp int64  `json:"lastLoginTimestamp" db:"lastLoginTimestamp"`
}

func main() {
	db, err := sql.Open("mysql", "your_url_conn")
	if err != nil {
		log.Fatal(err)
	}

	// get items db
	info := kit.SelectItems(dbkit.QueryInfo{
		Query: "SELECT * FROM UsersDB.Users",
	})

	items := dbkit.ParseToStruct[User](info.Items)
	fmt.Println(items) // out: []User{...}
}
```