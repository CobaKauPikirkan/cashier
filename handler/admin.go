package handler

import (
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/service/admin_service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService admin_service.Admin
}

func NewAdminHandler(adminService admin_service.Admin) *AdminHandler {
	return &AdminHandler{adminService}
}

func (h *AdminHandler)FindAllUser(c *gin.Context)  {
	users, err := h.adminService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *AdminHandler)FindByUser(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)// convert int to string
	helper.PanicIfError(err)

	users, err := h.adminService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *AdminHandler)UpdateUser(c *gin.Context)  {
	var userRequest models.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword := HashPassword(userRequest.Password)
	userRequest.Password = hashedPassword

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)// convert int to string
	helper.PanicIfError(err)

	users, err := h.adminService.Update(id, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *AdminHandler)DeleteUser(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)// convert int to string
	helper.PanicIfError(err)

	users, err := h.adminService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}