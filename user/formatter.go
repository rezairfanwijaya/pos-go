package user

type userFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func UserFormatter(user User, token string) userFormatter {
	return userFormatter{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}
}
