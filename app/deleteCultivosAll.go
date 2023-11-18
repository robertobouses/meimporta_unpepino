package app

import "log"

func (s *Service) DeleteCropsAll() error {
	err := s.repo.DeleteCropsAll()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}
