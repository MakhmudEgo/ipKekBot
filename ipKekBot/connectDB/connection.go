package connectDB

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ipKekBot/users"
)

type Ips struct {
	ID          uint    `gorm:"autoIncrement"`
	Query       string  `gorm:"column:query"`
	Status      string  `gorm:"column:status"`
	Country     string  `gorm:"column:country"`
	CountryCode string  `gorm:"column:countryCode"`
	Region      string  `gorm:"column:region"`
	RegionName  string  `gorm:"column:regionName"`
	City        string  `gorm:"column:city"`
	Zip         string  `gorm:"column:zip"`
	Lat         float64 `gorm:"column:lat"`
	Lon         float64 `gorm:"column:lon"`
	Timezone    string  `gorm:"column:timezone"`
	Isp         string  `gorm:"column:isp"`
	Org         string  `gorm:"column:org"`
	As          string  `gorm:"column:as"`
}

func Connect() *gorm.DB {
	configDB := "host=192.168.235.110 " +
		"user=postgres " +
		"password=lol " +
		"dbname=postgres " +
		"port=5432 " +
		"sslmode=disable"

	db, err := gorm.Open(postgres.Open(configDB), &gorm.Config{})
	if err != nil {
		log.Panic("no connect")
	}
	admin := &users.Users{Id: 234899515, Adm: true}
	if err := db.First(&admin); err.Error != nil {
		result := db.Create(&admin)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
	return db
}
