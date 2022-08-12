package request

type User struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}
type Artikel struct {
	ArtikelId   string `json:"artikel_id"`
	Judul       string `json:"judul"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}

type URL struct {
	Host string
	Port string
}
