package controller

import (
	"errors"
	"net/http"
	"ngc11/dto"
	"ngc11/entity"
	"ngc11/repository"
	"ngc11/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) User {
	return User{
		DB: db,
	}
}

func (controller User) RegisterUser(c echo.Context) error {
	body := dto.RegisterUserBody{}

	c.Bind(&body)
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	user := entity.Users{
		Username:      body.Username,
		Password:      string(hash),
		DepositAmount: body.DepositAmount,
	}

	// Create to DB
	if err = repository.CreateUser(&user, controller.DB); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return utils.ErrorMessage(c, &utils.ErrDuplicateKey)
		}
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	return c.JSON(http.StatusCreated, user)
}

func (controller User) LoginUser(c echo.Context) error {
	body := dto.LoginUserBody{}

	c.Bind(&body)
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// get user
	user, err := repository.GetUserByUsername(body.Username, controller.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorMessage(c, &utils.ErrUsernameNotFound)
		}
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return utils.ErrorMessage(c, &utils.ErrInvalidPassword)
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrGenerateToken)
	}

	response := dto.LoginResponse{
		Code:        http.StatusOK,
		Message:     "Login Success",
		AccessToken: tokenString,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller User) GetProducts(c echo.Context) error {
	products, err := repository.GetAllProduct(controller.DB)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ErrorMessage(c, &utils.ErrDataNotFound)
	}
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	return c.JSON(http.StatusOK, products)
}

func (controller User) CreateTransaction(c echo.Context) error {
	body := dto.TransactionReq{}

	c.Bind(&body)
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// get username from header
	username := c.Get("user").(string)

	products := []entity.Products{}
	totalAmmount := 0

	// Start DB transaction
	tx := controller.DB.Begin()

	// get user data
	user := entity.Users{}
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		tx.Rollback()
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// get product data
	for _, v := range body.Data {
		p := entity.Products{}

		// get product by id where product stock > order qty
		if err := tx.Where("id = ? AND stock > ?", v.ProductID, v.Qty).First(&p).Error; err != nil {
			tx.Rollback()

			return c.JSON(http.StatusBadRequest, dto.ErrTransactionResponse{
				Code:    http.StatusBadRequest,
				Message: "Product out of stock",
			})
		}

		p.Stock -= v.Qty
		totalAmmount += p.Price * v.Qty
		products = append(products, p)
	}

	// check if balance grater than total ammount
	balance := user.DepositAmount - totalAmmount
	if user.DepositAmount < totalAmmount || balance < 0 {
		tx.Rollback()

		return c.JSON(http.StatusBadRequest, dto.ErrTransactionResponse{
			Code:    http.StatusBadRequest,
			Message: "Not enough balance",
		})
	}

	// decreasing user balance
	if err := tx.Model(&user).Update("deposit_amount", balance).Error; err != nil {
		tx.Rollback()

		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// decreasing product stock
	for _, v := range products {
		if err := tx.Model(&entity.Products{}).Where("id = ?", v.ID).Update("stock", v.Stock).Error; err != nil {
			tx.Rollback()

			return utils.ErrorMessage(c, &utils.ErrInternalServer)
		}
	}

	// add new transaction
	transaction := entity.Transactions{
		UsersID:     user.ID,
		TotalAmount: totalAmmount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// add new transaction_details
	for _, v := range body.Data {
		td := entity.TransactionDetails{}

		td.TransactionsID = transaction.ID
		td.ProductsID = uint(v.ProductID)
		td.Qty = v.Qty

		if err := tx.Create(&td).Error; err != nil {
			tx.Rollback()
			return utils.ErrorMessage(c, &utils.ErrInternalServer)
		}
	}

	// Commit transaction
	tx.Commit()

	response := dto.TransactionResponse{
		Code:     http.StatusOK,
		Products: products,
	}

	return c.JSON(http.StatusOK, response)
}
