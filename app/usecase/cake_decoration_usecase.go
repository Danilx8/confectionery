package usecase

import "app/app/domain"

type cakeDecorationUsecase struct {
	cakeDecorationRepository domain.CakeDecorationRepository
}

func NewCakeDeoctaionUsecase(cakeDecorationRepository domain.CakeDecorationRepository) domain.CakeDecorationUsecase {
	return &cakeDecorationUsecase{
		cakeDecorationRepository: cakeDecorationRepository,
	}
}

func (i *cakeDecorationUsecase) GetAll() ([]domain.CakeDecoration, error) {
	return i.cakeDecorationRepository.FetchAll()
}

func (i *cakeDecorationUsecase) Edit(cakeDecoration *domain.CakeDecoration) error {
	return i.cakeDecorationRepository.Edit(cakeDecoration)
}

func (i *cakeDecorationUsecase) Delete(article string) error {
	return i.cakeDecorationRepository.Delete(article)
}
