package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const PlantsTableName = "plants"

type plant struct {
	Id        uint64           `db:"id,omitempty"`
	UserId    uint64           `db:"user_id"`
	Name      string           `db:"name"`
	City      string           `db:"city"`
	Address   string           `db:"address"`
	Type      domain.PlantType `db:"type"`
	CreatedAt time.Time        `db:"created_at"`
	UpdatedAt time.Time        `db:"updated_at"`
	DeletedAt *time.Time       `db:"deleted_at"`
}

type PlantRepository interface {
	Save(p domain.Plant) (domain.Plant, error)
	FindList(uId uint64) ([]domain.Plant, error)
	FindById(id uint64) (domain.Plant, error)
	Update(plant domain.Plant) (domain.Plant, error)
	Delete(id uint64) error
}

type plantRepository struct {
	coll db.Collection
	sess db.Session
}

func NewPlantRepository(sess db.Session) PlantRepository {
	return plantRepository{
		coll: sess.Collection(PlantsTableName),
		sess: sess,
	}
}

func (r plantRepository) Save(p domain.Plant) (domain.Plant, error) {
	pl := r.mapDomainToModel(p)
	pl.CreatedAt = time.Now()
	pl.UpdatedAt = time.Now()

	err := r.coll.InsertReturning(&pl)
	if err != nil {
		return domain.Plant{}, err
	}

	p = r.mapModelToDomain(pl)
	return p, nil
}

func (r plantRepository) FindList(uId uint64) ([]domain.Plant, error) {
	var plants []plant

	err := r.coll.Find(db.Cond{
		"user_id":    uId,
		"deleted_at": nil,
	}).All(&plants)
	if err != nil {
		return nil, err
	}

	ps := r.mapModelToDomainCollection(plants)
	return ps, nil
}

func (r plantRepository) FindById(id uint64) (domain.Plant, error) {
	var p plant

	err := r.coll.Find(db.Cond{
		"id":         id,
		"deleted_at": nil,
	}).One(&p)
	if err != nil {
		return domain.Plant{}, err
	}

	pl := r.mapModelToDomain(p)
	return pl, nil
}

func (r plantRepository) Update(plant domain.Plant) (domain.Plant, error) {
	p := r.mapDomainToModel(plant)
	p.UpdatedAt = time.Now()
	err := r.coll.Find(db.Cond{"id": p.Id, "deleted_at": nil}).Update(&p)
	if err != nil {
		return domain.Plant{}, err
	}
	return r.mapModelToDomain(p), nil
}

func (r plantRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_at": nil}).Update(map[string]interface{}{"deleted_at": time.Now()})
}

func (r plantRepository) mapDomainToModel(p domain.Plant) plant {
	return plant{
		Id:        p.Id,
		UserId:    p.UserId,
		Name:      p.Name,
		City:      p.City,
		Address:   p.Address,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func (r plantRepository) mapModelToDomain(p plant) domain.Plant {
	return domain.Plant{
		Id:        p.Id,
		UserId:    p.UserId,
		Name:      p.Name,
		City:      p.City,
		Address:   p.Address,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func (r plantRepository) mapModelToDomainCollection(plants []plant) []domain.Plant {
	ps := make([]domain.Plant, len(plants))
	for i, p := range plants {
		ps[i] = r.mapModelToDomain(p)
	}
	return ps
}
