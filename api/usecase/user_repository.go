package usecase

import "github.com/Okinawas/blog_api/api/domain"

type UserRepository interface {
  Store(domain.User) (int, error)
  FindById(int)(domain.User, error)
  FindAll()(domain.Users, error)
}

