package models

type Mahasiswa struct {
	Npm_mahasiswa   string `json:"npm_mahasiswa" gorm:"type:varchar(100);primaryKey"`
	Nama_mahasiswa  string `json:"nama_mahasiswa" gorm:"type:varchar(100)"`
	Email_mahasiswa string `json:"email_mahasiswa" gorm:"type:varchar(100)"`
	Id_jurusan      uint   `json:"id_jurusan"`
	Id_fakultas     uint   `json:"id_fakultas"`
}
