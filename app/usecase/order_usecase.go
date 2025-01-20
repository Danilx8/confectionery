package usecase

import (
	"app/app/domain"
	"fmt"
	"strings"
	"time"
)

type orderUsecase struct {
	orderRepository   domain.OrderRepository
	historyRepository domain.OrdersHistoryRepository
}

func NewOrderUsecase(orderRepository domain.OrderRepository, historyRepository domain.OrdersHistoryRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository:   orderRepository,
		historyRepository: historyRepository,
	}
}

func (o orderUsecase) Create(order *domain.Order) error {
	var id strings.Builder
	id.WriteString(strings.Join(strings.Split(time.Now().Format(time.DateOnly), "-")[:], ""))
	fullName := strings.Split(order.Orderer.FullName, " ")

	firstLetter := func(name string) string {
		if len(name) > 0 {
			return string([]rune(name)[0])
		} else {
			return "_"
		}
	}
	id.WriteString(firstLetter(fullName[0])) // First letter of a name
	id.WriteString(firstLetter(fullName[1])) // First letter of a surname

	var count int64
	var err error
	if count, err = o.orderRepository.CountToday(); err != nil {
		return err
	}
	if count > 99 {
		count = 100 - count
	}
	id.WriteString(fmt.Sprintf("%02d", count)) // Adds 0 to a number if necessary
	order.ID = id.String()
	order.Date = time.Now()

	if err = o.orderRepository.Create(order); err != nil {
		return err
	}
	return nil
}

func (o orderUsecase) Fetch() ([]domain.Order, error) {
	if orders, err := o.orderRepository.Fetch(); err != nil {
		return nil, err
	} else {
		return orders, nil
	}
}

func (o orderUsecase) FetchOwn(login string) ([]domain.Order, error) {
	if orders, err := o.orderRepository.FetchByUser(login); err != nil {
		return nil, err
	} else {
		return orders, err
	}
}

func (o orderUsecase) GetByID(id string) (*domain.Order, error) {
	if order, err := o.orderRepository.FetchById(id); err != nil {
		return nil, err
	} else {
		return order, nil
	}
}

func (o orderUsecase) FetchByStatus(status string) ([]domain.Order, error) {
	if orders, err := o.orderRepository.FetchByStatus(status); err != nil {
		return nil, err
	} else {
		return orders, nil
	}
}

func (o orderUsecase) Update(order *domain.Order) error {
	if oldOrder, err := o.orderRepository.FetchById(order.ID); err != nil {
		return err
	} else if oldOrder.Status != order.Status {
		if err = o.historyRepository.Create(&domain.OrdersHistory{
			Order:     *order,
			Date:      time.Now().Format(time.DateOnly),
			OldStatus: oldOrder.Status,
			NewStatus: order.Status,
			Time:      time.Now().Format(time.DateTime),
		}); err != nil {
			return err
		}
	}

	if err := o.orderRepository.Update(order); err != nil {
		return err
	}
	return nil
}

func (o orderUsecase) Delete(id string) error {
	if err := o.orderRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (o orderUsecase) MapOrder(order *domain.Order) domain.OrderResponse {
	return domain.OrderResponse{
		ID:                     order.ID,
		Date:                   order.Date.String(),
		Name:                   order.Name,
		Status:                 order.Status,
		Price:                  order.Price,
		OrdererName:            order.OrdererName,
		ExpectedFulfilmentDate: order.ExpectedFulfilmentDate.String(),
		AssignedManagerName:    order.AssignedManagerName,
	}
}
