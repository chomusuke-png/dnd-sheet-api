package character

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetAll() ([]Character, error) {
	return service.repository.FindAll()
}

func (service *Service) GetByID(id uint) (*Character, error) {
	return service.repository.FindByID(id)
}

func (service *Service) Create(character *Character) error {
	return service.repository.Create(character)
}

func (service *Service) Update(character *Character) error {
	return service.repository.Update(character)
}

func (service *Service) Delete(id uint) error {
	return service.repository.Delete(id)
}
