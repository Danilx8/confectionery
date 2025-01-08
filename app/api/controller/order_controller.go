package controller

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
	UserUsecase  domain.LoginUsecase
}

func NewOrderController(orderUsecase domain.OrderUsecase, userUsecase domain.LoginUsecase) *OrderController {
	return &OrderController{
		OrderUsecase: orderUsecase,
		UserUsecase:  userUsecase,
	}
}

// Create godoc
// @Summary	Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param        data    body   domain.OrderRequest true  "scheme of order request"
// @Success 200 {object} domain.Order
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/create [post]
func (oc *OrderController) Create(c *gin.Context) {
	var request domain.OrderRequest
	var err error
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var user *domain.User
	if login, exists := c.Get("x-user-login"); !exists {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Couldn't get your login. Try refresh your token"})
		return
	} else if user, err = oc.UserUsecase.GetUserByLogin(login.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var order *domain.Order
	if strings.Contains(user.Role, domain.RoleName[domain.ClientManager]) { // If order is placed by a manager
		order.Status = domain.StatusName[domain.Specification]
		order.AssignedManager = *user
		order.OrdererName = request.Orderer
	} else { // if order is placed by a client
		order.Orderer = *user
		order.AssignedManagerName = request.Manager
	} // no other option is available by design

	order.Examples = request.Examples

	if err = oc.OrderUsecase.Create(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) Update(c *gin.Context) {
	var request domain.OrderRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if order, err := oc.OrderUsecase.GetByID(request.ID); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.New] {
		c.JSON(http.StatusNotAcceptable, domain.ErrorResponse{Message: "You can't modify or delete orders with status other than \"New\""})
		return
	}

	// TODO: доделать обновление заказа
}
