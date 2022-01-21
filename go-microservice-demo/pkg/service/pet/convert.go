package pet

import (
	"github.com/win5do/golang-microservice-demo/pkg/api/petpb"
	"github.com/win5do/golang-microservice-demo/pkg/model"
)

func ModelPet2PbPet(in *model.Pet) *petpb.Pet {
	return &petpb.Pet{
		Id:        in.Id,
		CreatedAt: time2Pb(in.CreatedAt),
		UpdatedAt: time2Pb(in.UpdatedAt),
		Name:      in.Name,
		Type:      in.Type,
		Age:       in.Age,
		Sex:       in.Sex,
		Owned:     in.Owned,
	}
}

func PbPet2ModelPet(in *petpb.Pet) *model.Pet {
	return &model.Pet{
		Common: model.Common{
			Id:        in.Id,
			CreatedAt: pb2Time(in.CreatedAt),
			UpdatedAt: pb2Time(in.UpdatedAt),
		},
		Name:  in.Name,
		Type:  in.Type,
		Age:   in.Age,
		Sex:   in.Sex,
		Owned: in.Owned,
	}
}

func ModelPet2PbPetList(in []*model.Pet) []*petpb.Pet {
	var out []*petpb.Pet
	for _, v := range in {
		out = append(out, ModelPet2PbPet(v))
	}
	return out
}

func ModelOwner2PbOwner(in *model.Owner) *petpb.Owner {
	return &petpb.Owner{
		Id:        in.Id,
		CreatedAt: time2Pb(in.CreatedAt),
		UpdatedAt: time2Pb(in.UpdatedAt),
		Name:      in.Name,
		Age:       in.Age,
		Sex:       in.Sex,
		Phone:     in.Phone,
	}
}

func PbOwner2ModelOwner(in *petpb.Owner) *model.Owner {
	return &model.Owner{
		Common: model.Common{
			Id:        in.Id,
			CreatedAt: pb2Time(in.CreatedAt),
			UpdatedAt: pb2Time(in.UpdatedAt),
		},
		Name:  in.Name,
		Age:   in.Age,
		Sex:   in.Sex,
		Phone: in.Phone,
	}
}

func ModelOwner2PbOwnerList(in []*model.Owner) []*petpb.Owner {
	var out []*petpb.Owner
	for _, v := range in {
		out = append(out, ModelOwner2PbOwner(v))
	}
	return out
}
