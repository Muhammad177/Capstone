package likes

import "context"

type likesUsecase struct {
	likesRepository Repository
}

func NewLikesUsecase(repository Repository) Usecase {
	return &likesUsecase{
		likesRepository: repository,
	}
}

func (usecase *likesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.likesRepository.GetAll(ctx)
}

func (usecase *likesUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.likesRepository.GetByID(ctx, id)
}

func (usecase *likesUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.likesRepository.GetByName(ctx, name)
}

func (usecase *likesUsecase) Create(ctx context.Context, likesDomain *Domain) (Domain, error) {
	return usecase.likesRepository.Create(ctx, likesDomain)
}

func (usecase *likesUsecase) Update(ctx context.Context, likesDomain *Domain, id string) (Domain, error) {
	return usecase.likesRepository.Update(ctx, likesDomain, id)
}

func (usecase *likesUsecase) Delete(ctx context.Context, id string) error {
	return usecase.likesRepository.Delete(ctx, id)
}

func (usecase *likesUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.likesRepository.Restore(ctx, id)
}

func (usecase *likesUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.likesRepository.ForceDelete(ctx, id)
}
