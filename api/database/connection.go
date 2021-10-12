package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/KlareTeam/interview-challenges/go/pricematic/models"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// TODO: Se declaran las migraciones a ejecutar cuando se ejecute la aplicación
//			Están incluidas los modelos Product y Price
func InitialMigration(db *gorm.DB ) {
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Price{})
}

// TODO: Establece la conexión con la base de datos a partir de los parámetros leidos en el .env
func Connection() *gorm.DB {
	var (
		host = os.Getenv("DATABASE_HOST")
		user = os.Getenv("DATABASE_USER")
		port = os.Getenv("DATABASE_PORT")
		password = os.Getenv("DATABASE_PASSWORD") 
		name = os.Getenv("DATABASE_NAME")
	)

	DB, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))

	if err != nil {  
		log.Fatalf("Error in connect the DB %v", err)  
		return nil 
	}  

	if err := DB.DB().Ping(); err != nil { 
		log.Fatalln("Error in make ping the DB " + err.Error())  
		return nil 
	} 

	if DB.Error != nil { 
		log.Fatalln("Any Error in connect the DB " + err.Error()) 
		return nil 
	}  

	return DB
}

// TODO: Esta función lee el archivo data.json y crea registros iniciales (Seeds) en la tabla Product
func DataSeed(db *gorm.DB)  {
	jsonFile, err := os.Open("data.json")
	
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products []models.Product
	var product models.Product

	json.Unmarshal([]byte(byteValue), &products)

	for i := 0; i < len(products); i++ {
		product.ID = uuid.NewV4()
		product.Name = products[i].Name
		product.Price = products[i].Price
		db.Create(&product)
  }
}
