package repo

import "github.com/HavvokLab/true-solar-monitoring/model"

type mockSolarmanCredentialRepo struct{}

func NewMockSolarmanCredentialRepo() SolarmanCredentialRepo {
	return &mockSolarmanCredentialRepo{}
}

func (m *mockSolarmanCredentialRepo) FindAll() ([]model.SolarmanCredential, error) {
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

func (m *mockSolarmanCredentialRepo) Create(credential *model.SolarmanCredential) error {
	return nil
}

func (m *mockSolarmanCredentialRepo) Update(id int64, credential *model.SolarmanCredential) error {
	return nil
}

func (m *mockSolarmanCredentialRepo) Delete(id int64) error {
	return nil
}
