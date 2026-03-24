package seed

import (
	"context"
	"log"
)

type Seeder interface {
	Run(ctx context.Context)
}

type Runner struct {
	seeders []Seeder
}

func NewRunner(seeders ...Seeder) *Runner {
	return &Runner{seeders: seeders}
}

func (r *Runner) Run(ctx context.Context) {
	log.Println("[Seeder] running all seeders...")

	for _, s := range r.seeders {
		s.Run(ctx)
	}

	log.Println("✅ all seeders executed")
}
