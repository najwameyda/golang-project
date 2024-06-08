package JurusanModel

type Jurusan struct {
	ID      int    `json:"Id" form:"Id"`
	Nim     int    `json:"Nim" form:"Nim"`
	Jurusan string `json:"Jurusan" form:"Jurusan"`
	NamaMk  string `json:"NamaMk" form:"NamaMk"`
}
