package repo

import "github.com/HavvokLab/true-solar-monitoring/model"

type mockKStarCredentialRepo struct{}

func NewMockKStarCredentialRepo() KStarCredentialRepo {
	return &mockKStarCredentialRepo{}
}

func (*mockKStarCredentialRepo) FindAll() ([]model.KStarCredential, error) {
	return []model.KStarCredential{
		{
			ID:       0,
			Username: "u2.kst",
			Password: "Truec[8mugiup18",
		},
	}, nil
}

func (*mockKStarCredentialRepo) Create(credential *model.KStarCredential) error {
	return nil
}

func (*mockKStarCredentialRepo) Update(id int64, credential *model.KStarCredential) error {
	return nil
}

func (*mockKStarCredentialRepo) Delete(id int64) error {
	return nil
}
