package worker

import (
	"github/gyu-young-park/go-archive/repository"
)

type iWork interface {
	do(store *repository.Storer)
}

type Service struct {
	store *repository.Storer
	works []iWork
}

func NewWorker(store *repository.Storer) *Service {
	s := &Service{
		store: store,
	}
	s.ready()
	return s
}

func (s *Service) ready() {
	s.register(newMediumWork())
}

func (s *Service) register(work iWork) {
	s.works = append(s.works, work)
}

func (s *Service) Execute() {
	for _, work := range s.works {
		work.do(s.store)
	}
}
