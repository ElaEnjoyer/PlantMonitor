package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type PlantService interface {
	Save(p domain.Plant) (domain.Plant, error)
	FindList(uId uint64) ([]domain.Plant, error)
	Find(id uint64) (interface{}, error)
	Update(p domain.Plant) (domain.Plant, error)
	Delete(id uint64) error
}

type plantService struct {
	plantRepo database.PlantRepository
}

func NewPlantService(pr database.PlantRepository) PlantService {
	return plantService{
		plantRepo: pr,
	}
}

func (s plantService) Save(p domain.Plant) (domain.Plant, error) {
	plant, err := s.plantRepo.Save(p)
	if err != nil {
		log.Printf("plantService.Save(s.plantRepo.Save): %s", err)
		return domain.Plant{}, err
	}

	return plant, nil
}

func (s plantService) FindList(uId uint64) ([]domain.Plant, error) {
	plants, err := s.plantRepo.FindList(uId)
	if err != nil {
		log.Printf("plantService.FindList(s.plantRepo.FindList): %s", err)
		return nil, err
	}

	return plants, nil
}

func (s plantService) Find(id uint64) (interface{}, error) {
	plant, err := s.plantRepo.FindById(id)
	if err != nil {
		log.Printf("plantService.Find(s.plantRepo.FindById): %s", err)
		return domain.Plant{}, err
	}

	return plant, nil
}

func (s plantService) Update(p domain.Plant) (domain.Plant, error) {
	plant, err := s.plantRepo.Update(p)
	if err != nil {
		log.Printf("plantService.Update(s.plantRepo.Update): %s", err)
		return domain.Plant{}, err
	}

	return plant, nil
}

func (s plantService) Delete(id uint64) error {
	err := s.plantRepo.Delete(id)
	if err != nil {
		log.Printf("plantService.Delete(s.plantRepo.Delete): %s", err)
		return err
	}

	return nil
}
