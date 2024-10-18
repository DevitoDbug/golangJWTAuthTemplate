package models

type User struct {
	FirstName  string `json:"firstName" validate:"required,alpha"`
	SecondName string `json:"secondName" validate:"required,alpha"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password"`
}

type PublicUserInfo struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

func (s *User) GetPublicUserInfo() PublicUserInfo {
	user := PublicUserInfo{
		FirstName:  s.FirstName,
		SecondName: s.SecondName,
		Email:      s.Email,
	}
	return user
}
