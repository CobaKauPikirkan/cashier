package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/service/manager_service"
	"github.com/gin-gonic/gin"
)

type ManagerHandler struct {
	managerService manager_service.Manager
}

func NewManagerHandler(managerService manager_service.Manager) *ManagerHandler {
	return &ManagerHandler{managerService}
}

func (h *ManagerHandler) FindAllTransaksiManager(c *gin.Context) {
	results, err := h.managerService.FindAllTransaksi()
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

func (h *ManagerHandler) FindMostFavouriteProduct(c *gin.Context) {
	results, err := h.managerService.FindMostFavourite()
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

func (h *ManagerHandler)FindByIdKasir(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)
	result, err := h.managerService.FindByIdKasir(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *ManagerHandler)FindByTgl(c *gin.Context)  {
	tgl := c.Query("tgl")
	filter := c.Query("filter")
	fmt.Println(tgl)
	result, err := h.managerService.FindByTanggal(tgl, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}