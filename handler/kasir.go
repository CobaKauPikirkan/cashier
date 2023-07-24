package handler

import (
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/service/kasir_service"
	"github.com/gin-gonic/gin"
)

type KasirHandler struct {
	kasirService kasir_service.Kasir
}

func NewKasirHandler(kasirService kasir_service.Kasir) *KasirHandler {
	return &KasirHandler{kasirService}
}

func (h *KasirHandler) FindAllTransaksi(c *gin.Context) {
	results, err := h.kasirService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}
func (h *KasirHandler)FindTransaksiByID(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)// convert int to string
	helper.PanicIfError(err)

	transaksi, detail, err := h.kasirService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"transaksi": transaksi,
			"detail": detail,
		},
	})
}

func (h *KasirHandler) CreateTransaksi(c *gin.Context) {
	var request models.KasirRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transaksi, detail, err := h.kasirService.CreateTransaksi(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"transaksi": transaksi,
			"detail": detail,
		},
	})
}

func (h *KasirHandler)UpdateTransaksi(c *gin.Context)  {
	var request models.KasirRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	transaksi, detail, err := h.kasirService.UpdateTransaksi(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"transaksi": transaksi,
			"detail": detail,
		},
	})
}

func (h *KasirHandler) FindMejaAvaible(c *gin.Context) {
	results, err := h.kasirService.FindMejaAvaible()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func (h *KasirHandler)UpdateMejaAvaible(c *gin.Context)  {
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

	meja, err := h.kasirService.UpdateMeja(id, mejaRequest)
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
