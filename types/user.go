package types

type User struct {
	UserID   UserID
	Username string
}

func NewUser(username string) *User {
	return &User{
		UserID:   NewUserID(),
		Username: username,
	}
}

func (u User) String() string {
	return u.Username
}

func (u User) Validate() error {
	return u.UserID.Validate()
}
