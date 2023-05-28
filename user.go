package snake

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Score	 uint64	`json:"score"`
}

type UserPublic struct {
	Username string `json:"username" binding:"required"`
	Score	 uint64	`json:"score" binding:"required"`
}
