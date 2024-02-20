package database

// _ "github.com/lib/pq"
// -----------------------------
// _ en el import se utiliza para importar un paquete únicamente por su efecto secundario, es decir, para ejecutar su código de inicialización
// sin utilizar explícitamente ninguna de sus funciones o variables en el código actual.
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dbname, user, password string) error {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	d, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db = d
	return nil
}

func GetDB() *sql.DB {
	return db
}
