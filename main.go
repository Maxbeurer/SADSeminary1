// Zhen Feng
package main

// Importar los modulos necesarios
// Voy a usar la libreria GORM para manejar la base de datos
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Crear un nuevo tipo struct para ToDoItem
type ToDoItem struct {
	//Field ID : tag json, tag gorm
	ID uint `json:"id" gorm:"primaryKey"`
	//Field Title : tag json
	Title string `json:"title"`
	//Field Status : tag json
	Status string `json:"status"`
}

// Declarar la variable 'db' para apuntar a la connexión de base de datos usando GORM
var db *gorm.DB

func main() {
	// Inicializa el enrutamiento de GIN
	r := gin.Default()

	// Connectar a una base de datos PostgreSQL usando GORM
	var err error
	dsn := "host=db user=postgres password=postgres dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Método AutoMigrate de GORM para crear automáticamente las tablas en la base de datos
	db.AutoMigrate(&ToDoItem{})

	// Definir rutas para REST API
	r.POST("/todos", createToDo)
	r.GET("/todos", getToDos)
	r.PUT("/todos/:id", updateToDo)
	r.DELETE("/todos/:id", deleteToDo)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")

}

// Función CREATE
func createToDo(c *gin.Context) {
	var item ToDoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&item)
	c.JSON(http.StatusOK, item)
}

// Función READ
func getToDos(c *gin.Context) {
	var items []ToDoItem
	db.Find(&items)
	c.JSON(http.StatusOK, items)
}

// Función UPDATE
func updateToDo(c *gin.Context) {
	var item ToDoItem
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&item)
	c.JSON(http.StatusOK, item)
}

// Función DELETE
func deleteToDo(c *gin.Context) {
	var item ToDoItem
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}
