package usecase

import "app/app/domain"

type ingredientUsecase struct {
	ingredientRepository domain.IngredientRepository
}

func NewIngredientUsecase(ingredientRepository domain.IngredientRepository) domain.IngredientUsecase {
	return &ingredientUsecase{
		ingredientRepository: ingredientRepository,
	}
}

func (i *ingredientUsecase) GetAll() ([]domain.Ingredient, error) {
	return i.ingredientRepository.Fetch()
}

func (i *ingredientUsecase) Edit(ingredient *domain.Ingredient) error {
	return i.ingredientRepository.Edit(ingredient)
}

func (i *ingredientUsecase) Delete(article string) error {
	return i.ingredientRepository.Delete(article)
}
