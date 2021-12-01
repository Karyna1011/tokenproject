package service

import (
  "github.com/go-chi/chi"
  "gitlab.com/distributed_lab/ape"
  "gitlab.com/tokend/subgroup/tokenproject/internal/config"
  "gitlab.com/tokend/subgroup/tokenproject/internal/data/postgres"
  "gitlab.com/tokend/subgroup/tokenproject/internal/service/handlers"
)

func (s *service) router(cfg config.Config) chi.Router {
  r := chi.NewRouter()

  r.Use(
    ape.RecoverMiddleware(s.log),
    ape.LoganMiddleware(s.log),
    ape.CtxMiddleware(
      handlers.CtxLog(s.log),
      handlers.CtxToken(postgres.NewTokenQ(cfg.DB())),
    ),
  )
  r.Route("/integrations/project", func(r chi.Router) {
   r.Get("/list", handlers.List)
   r.Get("/add", handlers.Add)
   // r.Get("/get/{id}", handlers.GetByIndex)
  })

  return r
}

