package service

import "context"

type Healthier interface {
	Status(context.Context) error
}

type Nameabe interface {
	Name() string
}

type Service interface {
	Healthier
	Nameabe
	Start(context.Context) error
	Stop(context.Context) error
}

type AsyncService interface {
	Nameabe
	Run(context.Context) (chan bool, error)
}
