package usecase

import (
	"app/app/domain"
	"fmt"
	"strings"
	"time"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
}

func NewOrderUsecase(orderRepository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{orderRepository: orderRepository}
}

func (o orderUsecase) Create(order *domain.Order) error {
	var id strings.Builder
	id.WriteString(time.Now().Format("DDMMYYYY"))
	fullName := strings.Split(order.OrdererName, " ")

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
	if users, err := o.orderRepository.Fetch(); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (o orderUsecase) FetchOwn(login string) ([]domain.Order, error) {
	if users, err := o.orderRepository.FetchByUser(login); err != nil {
		return nil, err
	} else {
		return users, err
	}
}

func (o orderUsecase) GetByID(id string) (*domain.Order, error) {
	if user, err := o.orderRepository.FetchById(id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (o orderUsecase) Update(order *domain.Order) error {
	if err := o.orderRepository.Update(order); err != nil {
		return err
	}
	return nil
}

func (o orderUsecase) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
