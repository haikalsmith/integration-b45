package authdto

type LoginResponse struct {
	ID    int    `gorm:"type: varchar(255)" json:"id"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	Role  string `gorm:"type: varchar(255)" json:"role"`
}

type RegisterResponse struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
