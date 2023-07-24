package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/service/menu_service"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuService menu_service.Menu
}

func NewMenuHandler(menuService menu_service.Menu) *MenuHandler {
	return &MenuHandler{menuService}
}

func (h *MenuHandler) GetAllMenu(c *gin.Context) {
	Menus, err := h.menuService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Menus,
	})
}

func (h *MenuHandler) GetMenuById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString) // convert int to string
	helper.PanicIfError(err)

	Menu, err := h.menuService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Menu,
	})
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	// Bind the JSON request body to a struct
	var menuRequest struct {
		NamaMenu  string `form:"nama_menu" binding:"required"`
		Jenis     string `form:"jenis" binding:"required"`
		Deskripsi string `form:"deskripsi" binding:"required"`
		Harga     string `form:"harga" binding:"required"`
	}
	err := c.ShouldBind(&menuRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert the harga string to an integer
	harga, err := strconv.Atoi(menuRequest.Harga)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid harga value",
		})
		return
	}
	jenis, err := strconv.Atoi(menuRequest.Jenis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid harga value",
		})
		return
	}

	// Get the file from the form data
	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the file to the "uploads" directory
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create a new MenuRequest struct with the file name and other fields
	menu := models.MenuRequest{
		NamaMenu:  menuRequest.NamaMenu,
		Deskripsi: menuRequest.Deskripsi,
		Jenis:     jenis,
		Harga:     harga,
		Gambar:    file.Filename,
	}
	fmt.Println(menu)
	// Call the service to create the menu
	createdMenu, err := h.menuService.Create(menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the created menu as JSON response
	c.JSON(http.StatusOK, gin.H{
		"data": createdMenu,
	})
}

func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	// Bind the JSON request body to a struct
	var menuRequest struct {
		NamaMenu  string `form:"nama_menu" binding:"required"`
		Jenis     string `form:"jenis" binding:"required"`
		Deskripsi string `form:"deskripsi" binding:"required"`
		Harga     string `form:"harga" binding:"required"`
	}
	err := c.ShouldBind(&menuRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert the harga string to an integer
	harga, err := strconv.Atoi(menuRequest.Harga)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid harga value",
		})
		return
	}
	jenis, err := strconv.Atoi(menuRequest.Jenis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid harga value",
		})
		return
	}

	// Get the file from the form data
	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the file to the "uploads" directory
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString) //convert int to string
	helper.PanicIfError(err)

	// Create a new MenuRequest struct with the file name and other fields
	menu := models.MenuRequest{
		NamaMenu:  menuRequest.NamaMenu,
		Deskripsi: menuRequest.Deskripsi,
		Jenis:     jenis,
		Harga:     harga,
		Gambar:    file.Filename,
	}
	fmt.Println(menu)
	// Call the service to create the menu
	createdMenu, err := h.menuService.Update(id, menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the created menu as JSON response
	c.JSON(http.StatusOK, gin.H{
		"data": createdMenu,
	})
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString) //convert int to string
	helper.PanicIfError(err)

	meja, err := h.menuService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
		"data":    meja,
	})
}
