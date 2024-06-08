package JurusanRepository

import (
	"Jurusan/Model"
	connection "Jurusan/Model/Connection"
	"Jurusan/Model/JurusanModel"
)

func InsertJurusan(jurusan JurusanModel.Jurusan) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "INSERT INTO mahasiswa (id, nim, jurusan, nama_mk) VALUES (?, ?, ?, ?)"

	tempResult := db.Exec(query, jurusan.ID, jurusan.Nim, jurusan.Jurusan, jurusan.NamaMk)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          nil,
		}
	}

	return result
}
func UpdateJurusan(jurusan JurusanModel.Jurusan) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "UPDATE mahasiswa SET id = ?, nim = ?, jurusan = ?, nama_mk = ? WHERE = ?"

	tempResult := db.Exec(query, jurusan.ID, jurusan.Nim, jurusan.Jurusan, jurusan.NamaMk)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          nil,
			}
		}
	}

	return result
}

func DeleteJurusan(request JurusanModel.Jurusan) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	db := connection.DB
	query = "DELETE FROM mahasiswa WHERE id = ?"

	tempResult := db.Exec(query, request.ID)

	if tempResult.Error != nil {

		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Model.BaseResponseModel{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Model.BaseResponseModel{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}

func GetJurusanByID(request JurusanModel.Jurusan) Model.BaseResponseModel {
	var query string
	var result Model.BaseResponseModel
	var JurusanList []JurusanModel.Jurusan
	db := connection.DB

	if request.ID != 0 {
		query = "SELECT * FROM mahasiwa WHERE id = ?"
	} else {
		query = "SELECT * FROM mahasiswa"
	}

	tempResult := db.Raw(query).Find(&JurusanList)

	if tempResult.Error != nil {
		result = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Model.BaseResponseModel{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:          JurusanList,
		}
	}

	return result
}
