package controller

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Berita struct {
	ID      int    `json:"id"`
	Tanggal string `json:"tanggal"`
	Title   string `json:"title"`
	Isi     string `json:"isi"`
}
