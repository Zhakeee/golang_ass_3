package pkg

import (
	md "github.com/LeilaBeken/golang_ass_3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
    dsn := "postgres://Zhakeee:qFClM2j9uHRo@ep-cool-art-311896.us-east-2.aws.neon.tech/neondb"
    // Open a database connection
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&md.Book{})

    return db, nil
}

type book struct{
	*md.Book
}

func (b *book) GetByID(id uint) error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.First(b, id)
    return result.Error
}

func (b *book) Create() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Create(b)
    return result.Error
}

func (b *book) Update() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Save(b)
    return result.Error
}

func (b *book) Delete() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Delete(b)
    return result.Error
}