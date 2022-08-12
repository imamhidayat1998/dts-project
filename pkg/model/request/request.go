package request

type User struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}


type URL struct {
	Host string
	Port string

}