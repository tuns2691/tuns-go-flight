package customer_handler

import (
	"fmt"
	customer_request "gin-tuns_go_flight/api/customer-api/request"
	customer_response "gin-tuns_go_flight/api/customer-api/response"
	"gin-tuns_go_flight/pb"
	"net/http"
	"strings"

	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.RPCCustomerClient
}

func NewCustomerHandler(customerClient pb.RPCCustomerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := customer_request.CreateCustomerRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Role:           req.Role,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBith:     req.DateOfBith,
		IdentityCard:   req.IdentityCard,
		Address:        req.Address,
		MembershipCard: req.MembershipCard,
		Password:       req.Password,
		Status:         req.Status,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.Password)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := Encrypt(req.Password)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.Password = encText
	}

	pRes, err := h.customerClient.CreateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &customer_response.CustomerResponse{
		Id:             pRes.Id,
		Role:           pRes.Role,
		Name:           pRes.Name,
		Email:          pRes.Email,
		PhoneNumber:    pRes.PhoneNumber,
		DateOfBith:     pRes.DateOfBith,
		IdentityCard:   pRes.IdentityCard,
		Address:        pRes.Address,
		MembershipCard: pRes.MembershipCard,
		Status:         pRes.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	req := customer_request.UpdateCustomerRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Id:             req.Id,
		Role:           req.Role,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBith:     req.DateOfBith,
		IdentityCard:   req.IdentityCard,
		Address:        req.Address,
		MembershipCard: req.MembershipCard,
		Status:         req.Status,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.Password)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := Encrypt(req.Password)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.Password = encText
	}

	pRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &customer_response.CustomerResponse{
		Id:             pRes.Id,
		Role:           pRes.Role,
		Name:           pRes.Name,
		Email:          pRes.Email,
		PhoneNumber:    pRes.PhoneNumber,
		DateOfBith:     pRes.DateOfBith,
		IdentityCard:   pRes.IdentityCard,
		Address:        pRes.Address,
		MembershipCard: pRes.MembershipCard,
		Status:         pRes.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) ChangePassword(c *gin.Context) {
	req := customer_request.ChangePasswordRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReqCheck := &pb.CustomerParamId{
		Id: req.Id,
	}
	// Check existed
	pResCheck, err := h.customerClient.FindById(c.Request.Context(), pReqCheck)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	if pResCheck == nil || pResCheck.Id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusNotFound),
			"error":  "ID does not exist in the system",
		})
		return
	}

	// Check old password not match
	decText, err := Decrypt(pResCheck.Password)
	if err != nil {
		fmt.Println("error decrypting your classified text: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	if req.OldPassword != decText {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 99,
			"error":  "Old password not match",
		})
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 99,
			"error":  "New password not match with Confirm password",
		})
		return
	}

	pReq := &pb.ChangePasswordRequest{
		CustomerId: req.Id,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.NewPassword)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := Encrypt(req.NewPassword)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.NewPassword = encText
	}

	pRes, err := h.customerClient.ChangePassword(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pRes,
	})
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(textStr string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(textStr)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(textStr string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(textStr)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
