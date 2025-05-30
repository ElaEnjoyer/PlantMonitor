package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type PlantRequest struct {
	Name    string `json:"name" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Type    string `json:"type" validate:"oneof=SOLAR WIND"`
}

func (r PlantRequest) ToDomainModel() (interface{}, error) {
	return domain.Plant{
		Name:    r.Name,
		City:    r.City,
		Address: r.Address,
		Type:    domain.PlantType(r.Type),
	}, nil
}
