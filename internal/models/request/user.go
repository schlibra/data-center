package request

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginKey struct {
	Username string `json:"username"`
}
type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}
type UserRegisterKey struct {
	Username string `json:"username"`
}
type UserUpdate struct {
	Nickname string `json:"nickname"`
}
type UserPassword struct {
	Password string `json:"password"`
}
