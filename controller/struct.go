package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// presensi
type Karyawan struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama        string     `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	PhoneNumber string     `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	Jabatan     string     `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
	Jam_kerja   []JamKerja `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string   `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
}

type Presensi struct {
	// ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Longitude    float64 `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64 `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string  `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Phone_number string  `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty" example:"MASUK"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

// ================================================================================================
// DHS
type Dhs struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mahasiswa  Mahasiswa          `bson:"mahasiswa,omitempty" json:"mahasiswa,omitempty"`
	MataKuliah []MataKuliah       `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm          int                `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Fakultas     Fakultas           `bson:"fakultas,omitempty" json:"fakultas,omitempty"`
	ProgramStudi ProgramStudi       `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	DosenWali    Dosen              `bson:"dosen_wali,omitempty" json:"dosen_wali,omitempty"`
}

type MataKuliah struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeMatkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama       string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Dosen      Dosen              `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Sks        int                `bson:"sks,omitempty" json:"sks,omitempty"`
	Nilai      string             `bson:"nilai,omitempty" json:"nilai,omitempty"`
}

type Dosen struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeDosen   string             `bson:"kode_dosen,omitempty" json:"kode_dosen,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type Fakultas struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeFakultas string             `bson:"kode_fakultas,omitempty" json:"kode_fakultas,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type ProgramStudi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeProgramStudi string             `bson:"kode_program_studi,omitempty" json:"kode_program_studi,omitempty"`
	Nama             string             `bson:"nama,omitempty" json:"nama,omitempty"`
}
