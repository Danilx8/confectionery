package usecase

import (
	"app/app/domain"
	"github.com/goccy/go-json"
	"math"
	"time"
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

func (tu *toolingUsecase) GetAll() ([]domain.ToolingResponse, error) {
	toolings, err := tu.toolingRepository.Fetch("")

	if err != nil {
		return nil, err
	}

	var response []domain.ToolingResponse
	for _, tooling := range toolings {
		var properties domain.ToolingRequest
		err = json.Unmarshal([]byte(tooling.Properties), &properties)
		if err != nil {
			continue
		}

		createTime := 0
		if properties.AcquireTime != nil {
			tmp, _ := math.Modf(math.Floor(time.Now().Sub(*properties.AcquireTime).Seconds() / 2600640)) // в месяцах
			createTime = int(tmp)
		}

		response = append(response, domain.ToolingResponse{
			Name:   tooling.Marking,
			Type:   tooling.TypeName,
			Age:    createTime,
			Amount: properties.Amount,
		})
	}

	return response, nil
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
