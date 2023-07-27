package config
import(
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	
)
var (
	db * gorm.DB

)
func Connect(){
	d, err := gorm.Open("mysql", "root:Sravani@123@tcp(127.0.0.1:3306)/BookStore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("database connection error")
		panic(err)
	}
	db = d

}
func GetDB() *gorm.DB{
	return db
}