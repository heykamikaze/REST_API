package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"` //нижнее подчеркивание в монго системное поле которое само генерит айдив системеб омит если оно пустое
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"` //не отдавать никому
	Email        string `json:"email" bson:"email"`
}

//берем данные, спускать в сервис,отдавать в сторедж

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
