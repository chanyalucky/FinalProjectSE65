package orm
import (
	"time" 
	"gorm.io/gorm"
)
type Booking struct { // สร้าง ตารางใน database ชื่อ Booking
gorm.Model
UserID string
CarID string
Start time.Time
End time.Time
}