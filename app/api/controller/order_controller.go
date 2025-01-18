package controller

import (
	"app/app/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
	"time"
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
	// Заказчик и менеджер по работе с клиентами:
	// Могут создавать заказы
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
		order.Status = domain.StatusName[domain.New]
	} // no other option is available by design

	order.Examples = request.Examples

	if err = oc.OrderUsecase.Create(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Get godoc
// @Summary	Get all orders
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/all [get]
func (oc *OrderController) Get(c *gin.Context) {
	//Директор
	//может просматривать все заказы,

	if orders, err := oc.OrderUsecase.Fetch(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No orders are placed yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetOwn godoc
// @Summary	Get all orders of current logged-in user
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/own [get]
func (oc *OrderController) GetOwn(c *gin.Context) {
	if login, exists := c.Get("x-user-login"); !exists {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Couldn't get your login"})
		return
	} else if user, err := oc.UserUsecase.GetUserByLogin(login.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if orders, err := oc.OrderUsecase.FetchOwn(user.Login); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetNew godoc
// @Summary	Get all orders with status "New"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/new [get]
func (oc OrderController) GetNew(c *gin.Context) {
	//Менеджер по работе с клиентами
	//видит все свободные (со статусом «Новый») и все свои заказы
	//(менеджером которых он является);

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.New]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No new orders are placed yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetCancelled godoc
// @Summary	Get all orders with status "Cancelled"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/cancelled [get]
func (oc OrderController) GetCancelled(c *gin.Context) {
	//Директор
	//видит все заказы

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Cancelled]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No cancelled orders are placed yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetSpecification godoc
// @Summary	Get all orders with status "Specification"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/specification [get]
func (oc OrderController) GetSpecification(c *gin.Context) {
	//Директор
	//Видит все заказы

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Specification]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetConfirmation godoc
// @Summary	Get all orders with status "Confirmation"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/confirmation [get]
func (oc OrderController) GetConfirmation(c *gin.Context) {
	//Директор
	//Видит все заказы

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Confirmation]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetSupplement godoc
// @Summary	Get all orders with status "Supplement"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/supplement [get]
func (oc OrderController) GetSupplement(c *gin.Context) {
	//Менеджер по закупкам
	//видит все заказы в статусе «Закупка»,

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Supplement]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetProduction godoc
// @Summary	Get all orders with status "Production"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/production [get]
func (oc OrderController) GetProduction(c *gin.Context) {
	// Мастер
	//видит все заказы, имеющие статус «Производство»

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Production]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetAssurance godoc
// @Summary	Get all orders with status "Assurance"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/assurance [get]
func (oc OrderController) GetAssurance(c *gin.Context) {
	// Мастер
	// видит все заказы, имеющие статус «Контроль»

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Assurance]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetReady godoc
// @Summary	Get all orders with status "Ready"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/ready [get]
func (oc OrderController) GetReady(c *gin.Context) {
	// Директор

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Ready]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// GetComplete godoc
// @Summary	Get all orders with status "Complete"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/complete [get]
func (oc OrderController) GetComplete(c *gin.Context) {
	// Директор

	if orders, err := oc.OrderUsecase.FetchByStatus(domain.StatusName[domain.Complete]); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	} else if len(orders) == 0 {
		c.JSON(http.StatusNoContent, domain.SuccessResponse{Message: "No order are being specified yet"})
		return
	} else {
		var response []domain.OrderResponse
		for _, order := range orders {
			response = append(response, oc.OrderUsecase.MapOrder(&order))
		}
		c.JSON(http.StatusOK, response)
		return
	}
}

// AcceptNewOrder godoc
// @Summary	Set status of a given order from new to "Specification"
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/accept [post]
func (oc OrderController) AcceptNewOrder(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//может принять свободный (со статусом «Новый») заказ (то
	//есть стать его менеджером), после чего статус заказа меняется
	//на «Составление спецификации»;

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.New] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not new"})
		return
	}

	order.Status = domain.StatusName[domain.Specification]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is sent to specification",
		id)})
}

// Cancel godoc
// @Summary	Cancel an order
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "scheme of order request"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/cancel [post]
func (oc *OrderController) Cancel(c *gin.Context) {
	//ЗАКАЗЧИК:
	//может отклонить заказ до присвоения ему статуса «Закупка».

	//Менеджер по работе с клиентами:
	//может отклонить «Новый» заказ клиента с указанием причины;
	//может менять статус заказов «Подтверждение» на статус
	//«Отклонен», если клиент без объяснения причины
	//отказывается от заявки, или на статус «Закупка», если клиент
	//согласен с условиями выполнения заказа и вносит предоплату;

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order
	role, exists := c.Get("x-user-role")
	if !exists {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Current user role not found in context"})
		return
	}

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if role == domain.RoleName[domain.Client] && order.Status != domain.StatusName[domain.New] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "You can't update the order after it was accepted"})
		return
	} else if role == domain.RoleName[domain.ClientManager] && !slices.Contains([]string{
		domain.StatusName[domain.New],
		domain.StatusName[domain.Confirmation],
	}, order.Status) {
		c.JSON(http.StatusConflict, domain.ErrorResponse{
			Message: "You can cancel orders only with status \"New\" or \"Confirmation\"",
		})
	}

	order.Status = domain.StatusName[domain.Cancelled]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is cancelled", id)})
}

// Specify godoc
// @Summary	Send ready specification and set current order status to "Confirmation"
// @Tags Orders
// @Accept json
// @Produce json
// @Param        data    body   domain.SpecificationRequest true  "scheme of order request"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/specify [post]
func (oc *OrderController) Specify(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//после составления чертежей и получения данных о сроках и
	//стоимости заказа
	//o должен ввести информацию о заказе: стоимость и
	//плановую дату завершения заказа,
	//o должен поменять статус «Составление спецификации»
	//на «Подтверждение»;

	var request domain.SpecificationRequest
	var err error
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var order *domain.Order
	if order, err = oc.OrderUsecase.GetByID(request.ID); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Specification] {
		c.JSON(http.StatusNotAcceptable, domain.ErrorResponse{
			Message: "You can't set specification for orders not in status \"Specification\""},
		)
		return
	}

	order.Price = request.Price
	order.ExpectedFulfilmentDate, err = time.Parse("YYYY-MM-DD hh:mm:ss", request.ExpectedFulfilmentDate)
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" got its specification",
		request.ID)})
}

// SetSupplement godoc
// @Summary	Set status of a given order from accept to supplement
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/supplement [post]
func (oc *OrderController) SetSupplement(c *gin.Context) {
	// Менеджер по работе с клиентами:
	//может менять статус заказов «Подтверждение» на статус
	//«Отклонен», если клиент без объяснения причины
	//отказывается от заявки, или на статус «Закупка», если клиент
	//согласен с условиями выполнения заказа и вносит предоплату;

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Confirmation] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not being confirmed"})
		return
	}

	order.Status = domain.StatusName[domain.Supplement]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is sent to supplement", id)})
}

// SetProduction godoc
// @Summary	Set status of a given order from supplement to production
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/production [post]
func (oc OrderController) SetProduction(c *gin.Context) {
	//Менеджер по закупкам
	//после поступления ингредиентов и украшений для торта может
	//менять статус «Закупка» на статус «Производство».

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Supplement] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not being supplied"})
		return
	}

	order.Status = domain.StatusName[domain.Production]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is sent to production", id)})
}

// SetAssurance godoc
// @Summary	Set status of a given order from production to assurance
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/assurance [post]
func (oc OrderController) SetAssurance(c *gin.Context) {
	//Мастер
	//после выполнения работ по производству заказа может
	//изменить статус на «Контроль»;

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Production] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not in production"})
		return
	}

	order.Status = domain.StatusName[domain.Assurance]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is sent to assurance", id)})
}

// AssureQuality godoc
// @Summary	Make a verdict on quality assurance
// @Tags Orders
// @Accept json
// @Produce json
// @Param        data    body   domain.AssuranceRequest true  "associative array of criteria with true for passed and false for not passed"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/verdict [post]
func (oc OrderController) AssureQuality(c *gin.Context) {
	//Мастер
	//После окончания изготовления изделия мастер должен
	//произвести контроль качества.

	var request domain.AssuranceRequest
	err := c.ShouldBind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(request.ID); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Production] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not new"})
		return
	}

	for name, criteria := range request.Criteria {
		if !criteria {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Message: fmt.Sprintf("Criterion \"%s\" is set to false", name),
			})
			return
		}
	}

	order.Status = domain.StatusName[domain.Ready]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" status is set to ready",
		request.ID)})
}

// SetComplete godoc
// @Summary	Set status of a given order from ready to complete
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/complete [post]
func (oc OrderController) SetComplete(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//после отгрузки заказа в выбранный день клиенту и получения
	//полной оплаты может сменить статус с «Готов» на
	//«Выполнен».

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var order *domain.Order

	if order, err = oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.Ready] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "Order is not ready"})
		return
	}

	order.Status = domain.StatusName[domain.Complete]
	if err = oc.OrderUsecase.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: fmt.Sprintf("Order \"%s\" is complete", id)})
}

// Delete godoc
// @Summary	Delete an order
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "scheme of order request"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/delete [delete]
func (oc *OrderController) Delete(c *gin.Context) {
	//ЗАКАЗЧИК:
	//может удалить новый заказ,

	var id string
	err := c.ShouldBind(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	if order, err := oc.OrderUsecase.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	} else if order.Status != domain.StatusName[domain.New] {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "You can't delete order which isn't new"})
		return
	}

	if err = oc.OrderUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Order deleted"})
}
