package steps

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type PrintStep struct {
	name    string
	content string
}

func (s *PrintStep) GetName() string {
	return s.name
}

func (s *PrintStep) Run(ctx context.Context) error {
	log.Println(s.content)
	return nil
}

func (s *PrintStep) Skip(ctx context.Context) bool {
	return false
}
