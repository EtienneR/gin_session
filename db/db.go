package db

// Les imports de librairies
import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// La structure de données
type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id"`
	Name     string `gorm:"not null;unique" form:"username"` // Utilisateur unique!
	Password string `gorm:"not null" form:"password"`
}

// Connexion à la BDD SQLite
func InitDb() *gorm.DB {
	// Ouverture de la connexion vers la BDD SQLite
	db, err := gorm.Open("sqlite3", "./data.db")
	// Afficher les requêtes SQL (facultatif)
	db.LogMode(true)

	// Création de la table "users"
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	if err != nil {
		panic(err)
	}

	return db
}
