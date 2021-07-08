package installer

import (
	"context"
	"fmt"
)

// Step defines the interface for a individual step in a installer
type Step interface {
	GetName() string
	Run(ctx context.Context) error
	Skip(ctx context.Context) bool
}

// Installer defines a installer instance use to install some set
// of deployable resources in a set of defines `Step`s
type Installer struct {
	steps []Step
}

// NewInstaller returns a new installer instance with the specified
// name and set of steps
func NewInstaller(name string, steps ...Step) *Installer {
	return &Installer{
		steps: steps,
	}
}

// Install runs each of the defined steps (unless the step's skip triggers)
// and returns a error if a step fails, aborting the install
func (i *Installer) Install(ctx context.Context) error {
	for _, step := range i.steps {
		if step.Skip(ctx) {
			continue
		}

		err := step.Run(ctx)
		if err != nil {
			return fmt.Errorf(
				"error executing step %s: %w",
				step.GetName(),
				err,
			)
		}
	}
	return nil
}
