package dao

type User struct {
	UserID       uint64 `gorm:"primary_key; not null"`
	Username     string `gorm:"size:15; not null; unique"`
	Password     string `gorm:"size:20; not null"`
	Email        string `gorm:"size:20; not null; unique"`
	Usertype     int    `gorm:"not null;default:1"`
	BasicInfo    string `gorm:"size:100"`
	Interdiction bool   `gorm:"default:false"`

	//User_type Profile `gorm:"ForeignKey:UserRefer"` 指定外键
}

type Student struct {
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255"`
	Age  uint64
}
