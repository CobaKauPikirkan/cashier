package kasir_service

import (
	"fmt"
	"time"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/kasir"
)

type kasirService struct {
	repository kasir.RepositoryKasir
}

func NewKasirService(repository kasir.RepositoryKasir) *kasirService {
	return &kasirService{repository}
}

func (s *kasirService) FindAll() ([]models.Result, error) {
	result, err := s.repository.FindAll()
	helper.PanicIfError(err)

	return result, err
}

func (s *kasirService) FindById(id int) (models.Transaksi, models.DetailTransaksi, error) {
	Transaksi, err := s.repository.FindByIdTransaksi(id)
	helper.PanicIfError(err)

	detail, err := s.repository.FindByIdDetailTransaksi(id)
	helper.PanicIfError(err)
	
	return Transaksi, detail, err
}

func (s *kasirService) CreateTransaksi(request models.KasirRequest) (models.Transaksi, models.DetailTransaksi, error) {
	transaksi := models.Transaksi{
        Tgl_Transaksi:  time.Now(),
        IdUser:         request.IdUser,
        IdMeja:         request.IdMeja,
        Nama_Pelanggan: request.Nama_Pelanggan,
        Status:         request.Status,
    }

    // Create a channel to receive the result of the CreateTransaksi repository method
    transaksiChan := make(chan models.Transaksi)
    go func() {
        result, err := s.repository.CreateTransaksi(transaksi)
        if err != nil {
            // Send the error to the channel if there was an error
            transaksiChan <- models.Transaksi{}
            fmt.Println("Error in CreateTransaksi:", err)
        } else {
            // Send the result to the channel if there was no error
            transaksiChan <- result
        }
    }()

    // Create a channel to receive the result of the SumMenu repository method
    totalChan := make(chan int)
    go func() {
        total, err := s.repository.SumMenu(request.Id_menu)
        if err != nil {
            // Send the error to the channel if there was an error
            totalChan <- 0
            fmt.Println("Error in SumMenu:", err)
        } else {
            // Send the result to the channel if there was no error
            totalChan <- total
        }
    }()


    // Wait for the results from the two channels
    result1 := <-transaksiChan
    total := <-totalChan

    // If either of the channels returned an error, return the error
    if result1.ID == 0 || total == 0 {
        return models.Transaksi{}, models.DetailTransaksi{}, fmt.Errorf("failed to create transaction")
    }

    // Create the detail transaction object
    detail := models.DetailTransaksi{
        Id_transaksi: result1.ID,
        Id_menu:      request.Id_menu,
        Harga:        total,
    }

    // Call the CreateDetailTransaksi repository method
    result2, err := s.repository.CreateDetailTransaksi(detail)
    if err != nil {
        return models.Transaksi{}, models.DetailTransaksi{}, err
    }

    return result1, result2, nil
}

func (s *kasirService) UpdateTransaksi(id int, request models.KasirRequest) (models.Transaksi, models.DetailTransaksi, error) {
	// return result1, result2, err
	Transaksi, err := s.repository.FindByIdTransaksi(id)
	helper.PanicIfError(err)
	Transaksi.Tgl_Transaksi = time.Now()
	Transaksi.IdUser = request.IdUser
	Transaksi.IdMeja = request.IdMeja
	Transaksi.Nama_Pelanggan = request.Nama_Pelanggan
	Transaksi.Status = request.Status

	// Create a channel to receive the result of the CreateTransaksi method
	transaksiChan := make(chan models.Transaksi)
	var transaksiErr error

	// Start a goroutine to execute the CreateTransaksi method
	go func() {
		result1, err := s.repository.UpdateTransaksi(Transaksi)
		if err != nil {
			transaksiErr = err
			return
		}
		transaksiChan <- result1
	}()

	// Call the SumMenu method asynchronously
	menuChan := make(chan int)
	var menuErr error
	go func() {
		total, err := s.repository.SumMenu(request.Id_menu)
		if err != nil {
			menuErr = err
			return
		}
		menuChan <- total
	}()

	// Call the FindByIdDetailTransaksi method asynchronously
	detailChan := make(chan models.DetailTransaksi)
	var detailErr error
	go func() {
		detail, err := s.repository.FindByIdDetailTransaksi(id)
		if err != nil {
			detailErr = err
			return
		}
		detailChan <- detail
	}()

	// Wait for all async methods to complete
	result1 := <-transaksiChan
	total := <-menuChan
	detail := <-detailChan

	// Check for errors from async methods
	if transaksiErr != nil {
		return models.Transaksi{}, models.DetailTransaksi{}, transaksiErr
	}
	if menuErr != nil {
		return models.Transaksi{}, models.DetailTransaksi{}, menuErr
	}
	if detailErr != nil {
		return models.Transaksi{}, models.DetailTransaksi{}, detailErr
	}

	detail.Id_transaksi = result1.ID
	detail.Id_menu = request.Id_menu
	detail.Harga = total

	result2, err := s.repository.UpdateDetailTransaksi(detail)
	helper.PanicIfError(err)

	return result1, result2, err
}

func (s *kasirService) FindMejaAvaible() ([]models.Meja, error) {
	meja, err := s.repository.FindMejaAvaible()
	helper.PanicIfError(err)
	
	return meja, err
}

func (s *kasirService)UpdateMeja(Id int, mejarequest models.MejaRequest) (models.Meja, error) {
	meja, err:=s.repository.FindMeja(Id)
	helper.PanicIfError(err)

	meja.NomorMeja = mejarequest.NomorMeja
	meja.IsAvaible = mejarequest.IsAvaible

	newMeja, err:= s.repository.UpdateMeja(meja)
	helper.PanicIfError(err)

	return newMeja, err
}