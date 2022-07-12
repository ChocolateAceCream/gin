package service

import (
	"fmt"
	"gin_demo/dao"
)

type Service struct {
	Dao *dao.Dao
}

func New() (s *Service) {
	dao := dao.GetDao()
	s = &Service{
		Dao: dao,
	}
	fmt.Println("New(), s: ", s)
	return
}

// Close Service.
func (s *Service) Close() {
	s.Dao.Close()
}
