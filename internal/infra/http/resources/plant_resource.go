package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type PlantDto struct {
	Id      uint64           `json:"id"`
	UserId  uint64           `json:"userid"`
	Name    string           `json:"name"`
	City    string           `json:"city"`
	Address string           `json:"address"`
	Type    domain.PlantType `json:"type"`
}

func (d PlantDto) DomainToDto(p domain.Plant) PlantDto {
	return PlantDto{
		Id:      p.Id,
		UserId:  p.UserId,
		Name:    p.Name,
		City:    p.City,
		Address: p.Address,
		Type:    p.Type,
	}
}

func (d PlantDto) DomainToDtoCollection(plants []domain.Plant) []PlantDto {
	plantsDto := make([]PlantDto, len(plants))
	for i, p := range plants {
		plantsDto[i] = d.DomainToDto(p)
	}
	return plantsDto
}
