package usecase

import (
	"app/app/domain"
	"github.com/goccy/go-json"
)

type toolingUsecase struct {
	toolingRepository domain.ToolingRepository
}

func NewToolingUsecase(toolingRepository domain.ToolingRepository) domain.ToolingUsecase {
	return &toolingUsecase{
		toolingRepository: toolingRepository,
	}
}

func (tu *toolingUsecase) HydrateProperties(request domain.ToolingRequest) (*domain.Tooling, error) {
	var tooling = domain.Tooling{
		Marking: request.Name,
		Type:    *request.Type,
	}

	request.Name = ""
	request.Type = nil
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	tooling.Properties = string(data)
	return &tooling, nil
}

func (tu *toolingUsecase) Create(tooling *domain.Tooling) error {
	return tu.toolingRepository.Create(tooling)
}

func (tu *toolingUsecase) GetAll() ([]domain.Tooling, error) {
	return tu.toolingRepository.Fetch("")
}

//func (tu *toolingUsecase) GetByConditions(conditions domain.ToolingRequest) ([]domain.Tooling, error) {
//	var text bytes.Buffer
//	var queries []string
//
//	if &conditions.Name != nil {
//		queries = append(queries, "Name = '"+conditions.Name+"'")
//	}
//
//	if &conditions.Type != nil {
//		queries = append(queries, "Type = '"+conditions.Type.Name+"'")
//	}
//
//	if &conditions.Amount != nil {
//		queries = append(queries, "properties->\"$.amount\" = "+strconv.Itoa(conditions.Amount))
//	}
//
//	if len(queries) > 0 {
//		text.WriteString(strings.Join(queries, ", "))
//	}
//
//	return tu.toolingRepository.Fetch(text.String())
//}

func (tu *toolingUsecase) Delete(marking string) error {
	return tu.toolingRepository.Remove(marking)
}

func (tu *toolingUsecase) Update(tooling *domain.Tooling) error {
	return tu.toolingRepository.Edit(tooling)
}
