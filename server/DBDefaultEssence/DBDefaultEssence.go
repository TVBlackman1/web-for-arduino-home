package DBDefaultEssence

type User struct {
	ID       int8   `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}