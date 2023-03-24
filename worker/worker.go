package worker

type iWork interface {
	do()
}

type Service struct {
	works []iWork
}

func NewWorker() *Service {
	s := &Service{}
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
		go work.do()
	}
}
