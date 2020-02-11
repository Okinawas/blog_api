package usecase

import "github/nntakuya/matsun_blogs/app/domain"

type UserRepository interface {
  Store(domain.User) (int, error)
  FindById(int)(domain.User, error)
  FindAll()(domain.Users, error)
}

