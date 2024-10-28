package models

type Kabupaten struct {
	Id            int64  `gorm:"primaryKey;autoIncrement"`
	NamaKabupaten string `json:"nama_kabupaten"`
}

type Kecamatan struct {
	Id            int64     `gorm:"primaryKey;autoIncrement"`
	IdKabupaten   int64     `json:"id_kabupaten"`
	IdProvinsi    int64     `json:"id_provinsi"`
	NamaKecamatan string    `json:"nama_kecamatan"`
	Kabupaten     Kabupaten `gorm:"foreignKey:IdKabupaten;references:Id"`
}

type Desa struct {
	Id            int64     `gorm:"primaryKey;autoIncrement"`
	IdKabupaten   int64     `json:"id_kabupaten"`
	IdKecamatan   int64     `json:"id_kecamatan"`
	NamaDesa      string    `json:"nama_desa"`
	Kabupaten     Kabupaten `gorm:"foreignKey:IdKabupaten;references:Id"`
	Kecamatan     Kecamatan `gorm:"foreignKey:IdKecamatan;references:Id"`
}
