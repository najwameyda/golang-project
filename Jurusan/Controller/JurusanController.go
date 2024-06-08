package Controller

import (
	"Jurusan/Model"
	"Jurusan/Model/JurusanModel"
	"Jurusan/Repository/JurusanRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJurusanByID(c *gin.Context) {
	var request JurusanModel.Jurusan
	var response Model.BaseResponseModel

	response = JurusanRepository.GetJurusanByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertJurusan(c *gin.Context) {
	var request JurusanModel.Jurusan
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = JurusanRepository.InsertJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateJurusan(c *gin.Context) {
	var request JurusanModel.Jurusan
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = JurusanRepository.UpdateJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteJurusan(c *gin.Context) {
	var request JurusanModel.Jurusan
	var response Model.BaseResponseModel

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Model.BaseResponseModel{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = JurusanRepository.UpdateJurusan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
