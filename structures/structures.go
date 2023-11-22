package structures

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "user"
}

type Response struct {
	Status  string `json:"status"`
	Data    User   `json:"data"`
	Message string `json:"message"`
}
