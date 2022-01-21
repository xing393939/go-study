package pet

import (
	"context"
	"github.com/win5do/golang-microservice-demo/pkg/model"

	"github.com/win5do/golang-microservice-demo/pkg/repository/db/dbcore"
)

type petDomain struct{}

func NewPetDomain() *petDomain {
	return &petDomain{}
}

func (*petDomain) PetDb(ctx context.Context) model.IPetDb {
	return &petDb{dbcore.GetDB(ctx)}
}

func (*petDomain) OwnerDb(ctx context.Context) model.IOwnerDb {
	return &ownerDb{dbcore.GetDB(ctx)}
}

func (*petDomain) OwnerPetDb(ctx context.Context) model.IOwnerPetDb {
	return &ownerPetDb{dbcore.GetDB(ctx)}
}
