package customerHandler

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/modules/customers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}

	r.GET("/customer", handler.GetAll)
	r.POST("/customer", handler.Insert)
	r.PUT("/customer/:id", handler.Update)
	r.DELETE("/customer/:id", handler.Delete)
	// r = router

	// r.DELETE("customer/", handler.GetAll)// DELETE = untuk hapus
}

func (handler *customerHandler) GetAll(c *gin.Context) {
	results, err := handler.UC.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ //http.StatusInternalServerError = internal object error
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
	})
}

func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Insert(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

// func (h *customerHandler) Update(c *gin.Context) {
// 	var request models.RequestUpdateCustomer

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  http.StatusBadRequest,
// 			"message": err.Error(),
// 			"data":    []string{},
// 		})

// 		return
// 	}

// 	err := h.UC.Update(&request)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 			"data":    []string{},
// 		})

// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{
//			"status":  http.StatusOK,
//			"message": "Updated successfully",
//			"data":    []string{},
//		})
//	}
func (handler *customerHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
			"data":    []string{},
		})
		return
	}

	var request models.RequestUpdateCustomer
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	// Set the ID from the URL to the request struct
	request.Id = uint(id)

	err = handler.UC.Update(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated successfully",
		"data":    []string{},
	})
}

// ...

func (h *customerHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	request := &models.RequestDeleteCustomer{Id: uint(id)}
	err = h.UC.Delete(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted successfully",
		"data":    []string{},
	})
}

// ...

// go get github.com/gin-gonic/gin = untuk install gin
