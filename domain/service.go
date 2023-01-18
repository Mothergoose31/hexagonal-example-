package domain

//Service ...
type Service interface {
	Find(code string) (*Store, error)
	Store(Store *Store) error
	Update(Store *Store) error
	FindAll() ([]*Store, error)
	Delete(code string) error
}

//Repository ...
type Repository interface {
	Find(code string) (*Store, error)
	Store(Store *Store) error
	Update(Store *Store) error
	FindAll() ([]*Store, error)
	Delete(code string) error
}

type service struct {
	Storerepo Repository
}

//NewStoreService ...
func NewStoreService(Storerepo Repository) Service {
	return &service{Storerepo: Storerepo}
}

func (s *service) Find(code string) (*Store, error) {
	return s.Storerepo.Find(code)
}

func (s *service) Store(Store *Store) error {

	return s.Storerepo.Store(Store)

}
func (s *service) Update(Store *Store) error {
	return s.Storerepo.Update(Store)
}

func (s *service) FindAll() ([]*Store, error) {
	return s.Storerepo.FindAll()
}

func (s *service) Delete(code string) error {

	return s.Storerepo.Delete(code)
}
