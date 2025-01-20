package usecase

import (
	"app/app/domain"
	"fmt"
)

type itemUsecase struct {
	itemRepository                        domain.ItemRepository
	ingredientSpecificationRepository     domain.IngredientSpecificationRepository
	cakeDecorationSpecificationRepository domain.CakeDecorationSpecificationRepository
	premadeSpecificationRepository        domain.PremadeSpecificationRepository
}

func NewItemUsecase(itemRepository domain.ItemRepository,
	ingredientSpecificationRepository domain.IngredientSpecificationRepository,
	cakeDecorationSpecificationRepository domain.CakeDecorationSpecificationRepository,
	premadeSpecificationRepository domain.PremadeSpecificationRepository) domain.ItemUseCase {
	return &itemUsecase{
		itemRepository:                        itemRepository,
		ingredientSpecificationRepository:     ingredientSpecificationRepository,
		cakeDecorationSpecificationRepository: cakeDecorationSpecificationRepository,
		premadeSpecificationRepository:        premadeSpecificationRepository,
	}
}
func (i itemUsecase) FetchAll() ([]domain.Item, error) {
	if items, err := i.itemRepository.Fetch(); err != nil {
		return []domain.Item{}, err
	} else {
		return items, nil
	}
}

func (i itemUsecase) FetchRequired(item string) (domain.ItemSpecificationsResponse, error) {
	ingredients, err := i.ingredientSpecificationRepository.FetchByItem(item)
	cakeDecorations, err := i.cakeDecorationSpecificationRepository.FetchByItem(item)
	premades, err := i.premadeSpecificationRepository.FetchByItem(item)

	if err != nil {
		return domain.ItemSpecificationsResponse{}, err
	}

	var ingredientsResponse []domain.IngredientSpecificationResponse
	for _, ingredient := range ingredients {
		ingredientsResponse = append(ingredientsResponse, domain.IngredientSpecificationResponse{
			Name:           ingredient.IngredientName,
			RequiredAmount: ingredient.Amount,
		})
	}

	var cakeDecorationsResponse []domain.CakeDecorationSpecificationResponse
	for _, cakeDecoration := range cakeDecorations {
		cakeDecorationsResponse = append(cakeDecorationsResponse, domain.CakeDecorationSpecificationResponse{
			Name:           cakeDecoration.CakeDecorationName,
			RequiredAmount: cakeDecoration.Amount,
		})
	}

	var premadesResponse []domain.PremadeSpecificationResponse
	for _, premade := range premades {
		premadesResponse = append(premadesResponse, domain.PremadeSpecificationResponse{
			Name:           premade.PremadeName,
			RequiredAmount: premade.Amount,
		})
	}

	return domain.ItemSpecificationsResponse{
		Ingredients: ingredientsResponse,
		Decorations: cakeDecorationsResponse,
		Premades:    premadesResponse,
		Steps:       fmt.Sprintf("1. *Приготовление коржей*:\n   * Включите духовку на 180°C.\n   * В глубокой миске взбейте яйца с сахаром до пышной массы.\n   * Добавьте растопленное масло и молоко, аккуратно перемешайте.\n   * Постепенно введите муку, разрыхлитель и ванилин. Тщательно перемешайте до однородной массы.\n   * Вылейте тесто в подготовленную форму и выпекайте 25-30 минут. Готовность проверяйте деревянной шпажкой.\n   * Остудите и разрежьте на коржи.\n\n2. *Приготовление крема*:\n   * В миске взбейте сливки до устойчивых пиков.\n   * В отдельной миске смешайте рикотту с сахарной пудрой и кокосовой стружкой.\n   * Осторожно объедините обе смеси до однородной консистенции.\n\n3. *Сборка торта*:\n   * На нижний корж нанесите частью крема, накройте вторым коржом.\n   * Повторяйте процесс, чередуя коржи и крем, пока не соберете весь торт.\n   * Обмажьте торт оставшимся кремом и посыпьте кокосовой стружкой. Украсить орехами по желанию.\n\n4. *Выставка*:\n   * Дайте торту немного настояться в холодильнике перед подачей, чтобы он хорошо пропитался."),
		Description: "Подготовьте все ингридиенты и чётко следуйте рецепту",
	}, nil
}
