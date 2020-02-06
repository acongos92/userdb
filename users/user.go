package users

//user model
import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100);not_null"`
}
