package handler

import (
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/service/meja_service"
	"github.com/gin-gonic/gin"
)

type MejaHandler struct {
	mejaService meja_service.Meja
}

func NewMejaHandler(mejaService meja_service.Meja) *MejaHandler {
	return &MejaHandler{mejaService}
}

func (h *MejaHandler)GetAllMeja(c *gin.Context)  {
	mejas, err := h.mejaService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	    return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mejas,
	})
}

func (h *MejaHandler)GetMejaById(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)// convert int to string
	helper.PanicIfError(err)

	meja, err := h.mejaService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": meja,
	})
}

func (h *MejaHandler)CreateMeja(c *gin.Context)  {
	var mejaRequest models.MejaRequest

	err := c.ShouldBindJSON(&mejaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	meja, err := h.mejaService.Create(mejaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": meja,
	})
}

func (h *MejaHandler)UpdateMeja(c *gin.Context)  {
	var mejaRequest models.MejaRequest

	err := c.ShouldBindJSON(&mejaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	meja, err := h.mejaService.Update(id, mejaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": meja,
	})
}

func (h *MejaHandler)DeleteMeja(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	meja, err := h.mejaService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
		"data": meja,
	})
}