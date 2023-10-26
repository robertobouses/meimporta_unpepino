package app

import "log"

func (s *Service) DeleteCultivosAll() error {
	err := s.repo.DeleteCultivosAll()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}
