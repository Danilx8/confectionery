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
// @Failure 400 {object} domain.ErrorResponse
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
			response = append(response, domain.OrderResponse{
				ID:                     order.ID,
				Date:                   order.Date.String(),
				Name:                   order.Name,
				Status:                 order.Status,
				Price:                  order.Price,
				OrdererName:            order.OrdererName,
				ExpectedFulfilmentDate: order.ExpectedFulfilmentDate.String(),
				AssignedManagerName:    order.AssignedManagerName,
			})
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
		c.JSON(http.StatusOK, orders)
		return
	}
}

// GetNew godoc
// @Summary	Get all orders with status "New"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/new [get]
func (oc OrderController) GetNew(c *gin.Context) {
	//Менеджер по работе с клиентами
	//видит все свободные (со статусом «Новый») и все свои заказы
	//(менеджером которых он является);

	//TODO implement me
	panic("implement me")
}

// GetSpecification godoc
// @Summary	Get all orders with status "Specification"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/specification [get]
func (oc OrderController) GetSpecification(c *gin.Context) {
	//ВСЕ

	//TODO implement me
	panic("implement me")
}

// GetConfirmation godoc
// @Summary	Get all orders with status "Confirmation"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/confirmation [get]
func (oc OrderController) GetConfirmation(c *gin.Context) {
	//ВСЕ

	//TODO implement me
	panic("implement me")
}

// GetSupplement godoc
// @Summary	Get all orders with status "Supplement"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/supplement [get]
func (oc OrderController) GetSupplement(c *gin.Context) {
	//Менеджер по закупкам
	//видит все заказы в статусе «Закупка»,

	// TODO implement me
	panic("implement me")
}

// GetProduction godoc
// @Summary	Get all orders with status "Production"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/production [get]
func (oc OrderController) GetProduction(c *gin.Context) {
	// Мастер
	//видит все заказы, имеющие статус «Производство»

	//TODO implement me
	panic("implement me")
}

// GetAssurance godoc
// @Summary	Get all orders with status "Assurance"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/assurance [get]
func (oc OrderController) GetAssurance(c *gin.Context) {
	// Мастер
	// видит все заказы, имеющие статус «Контроль»

	//TODO implement me
	panic("implement me")
}

// GetReady godoc
// @Summary	Get all orders with status "Ready"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/ready [get]
func (oc OrderController) GetReady(c *gin.Context) {
	// ВСЕ

	//TODO implement me
	panic("implement me")
}

// GetComplete godoc
// @Summary	Get all orders with status "Complete"
// @Tags Orders
// @Produce json
// @Success 200 {object} []domain.OrderResponse
// @Success 204 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/complete [get]
func (oc OrderController) GetComplete(c *gin.Context) {
	// ВСЕ

	//TODO implement me
	panic("implement me")
}

// AcceptNewOrder godoc
// @Summary	Set status of a given order from new to accept
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/accept [post]
func (oc OrderController) AcceptNewOrder(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//может принять свободный (со статусом «Новый») заказ (то
	//есть стать его менеджером), после чего статус заказа меняется
	//на «Составление спецификации»;

	//TODO implement me
	panic("implement me")
}

// SetSupplement godoc
// @Summary	Set status of a given order from accept to supplement
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/supplement [post]
func (oc *OrderController) SetSupplement(c *gin.Context) {
	// Менеджер по работе с клиентами:
	//может менять статус заказов «Подтверждение» на статус
	//«Отклонен», если клиент без объяснения причины
	//отказывается от заявки, или на статус «Закупка», если клиент
	//согласен с условиями выполнения заказа и вносит предоплату;

	//TODO implement me
	panic("implement me")
}

// SetProduction godoc
// @Summary	Set status of a given order from supplement to production
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/production [post]
func (oc OrderController) SetProduction(c *gin.Context) {
	//Менеджер по закупкам
	//после поступления ингредиентов и украшений для торта может
	//менять статус «Закупка» на статус «Производство».

	//TODO implement me
	panic("implement me")
}

// SetComplete godoc
// @Summary	Set status of a given order from ready to complete
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "order id"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/complete [post]
func (oc OrderController) SetComplete(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//после отгрузки заказа в выбранный день клиенту и получения
	//полной оплаты может сменить статус с «Готов» на
	//«Выполнен».

	// TODO implement me
	panic("implement me")
}

func (oc OrderController) SetAssurance(c *gin.Context) {
	//Мастер
	//после выполнения работ по производству заказа может
	//изменить статус на «Контроль»;

	// TODO implement me
	panic("implement me")
}

func (oc *OrderController) Update(c *gin.Context) {
	//Менеджер по работе с клиентами:
	//после составления чертежей и получения данных о сроках и
	//стоимости заказа
	//o должен ввести информацию о заказе: стоимость и
	//плановую дату завершения заказа,
	//o должен поменять статус «Составление спецификации»
	//на «Подтверждение»;

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
	panic("implement me")
}

// Cancel godoc
// @Summary	Cancel an order
// @Tags Orders
// @Accept html/text
// @Produce json
// @Param        data    body   string true  "scheme of order request"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /order/cancel [post]
func (oc *OrderController) Cancel(c *gin.Context) {
	//ЗАКАЗЧИК:
	//может отклонить заказ до присвоения ему статуса «Закупка».

	//Менеджер по работе с клиентами:
	//может отклонить «Новый» заказ клиента с указанием причины;

	//TODO implement me
	panic("implement me")
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

	//TODO implement me
	panic("implement me")
}
