package request

type ClaimsJWT struct {
	Username  string `json:"username"`
	Role      string `json:"role,omitempty"`
	UserId    int    `json:"id"`
	PegawaiId int    `json:"pegawai_id"`
}
