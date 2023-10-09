package repo

import "github.com/HavvokLab/true-solar-monitoring/model"

type mockSolarmanCredentialRepo struct{}

func NewMockSolarmanCredentialRepo() SolarmanCredentialRepo {
	return &mockSolarmanCredentialRepo{}
}

func (m *mockSolarmanCredentialRepo) GetCredentials() ([]model.SolarmanCredential, error) {
	return []model.SolarmanCredential{
		{
			ID:        1,
			Username:  "bignode.invt.th@gmail.com",
			Password:  "123456*",
			AppSecret: "222c202135013aee622c71cdf8c47757",
			AppID:     "202010143565002",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	}, nil
}
