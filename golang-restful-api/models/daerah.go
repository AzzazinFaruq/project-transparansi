package models

type Kabupaten struct {
	Id   string `gorm:"primaryKey"`
	Nama string
}

type Kecamatan struct {
	ID_KEC  string `gorm:"primaryKey"`
	ID_KAB   string
	NAMA_KEC   string
	Kabupaten Kabupaten `gorm:"foreignKey:ID_KAB;references:Id"`
}

type Desa struct{
	ID_DESA int64 `gorm:primaryKey`
	ID_KAB string
	ID_KEC string
	NAMA_DES string
	Kabupaten Kabupaten `gorm:"foreignKey:ID_KAB;references:Id"`
	Kecamatan Kecamatan `gorm:"foreignKey:ID_KEC;references:ID_KEC"`
}