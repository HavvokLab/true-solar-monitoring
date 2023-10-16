package repo

import "github.com/HavvokLab/true-solar-monitoring/model"

type mockKStarCredentialRepo struct{}

func NewMockKStarCredentialRepo() KStarCredentialRepo {
	return &mockKStarCredentialRepo{}
}

func (*mockKStarCredentialRepo) GetCredentials() ([]model.KStarCredential, error) {
	return []model.KStarCredential{
		{
			ID:       0,
			Username: "u2.kst",
			Password: "Truec[8mugiup18",
		},
	}, nil
}
