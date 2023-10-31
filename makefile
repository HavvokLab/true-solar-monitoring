run:
	go run ./cmd/api/main.go

test:
	go run ./cmd/test/main.go

apply_schema:
	atlas schema apply --url "sqlite://database.db" --to "file://script/schema.hcl"

dp:
	go run ./cmd/daily_production/main.go

mp:
	go run ./cmd/monthly_production/main.go

pa:
	go run ./cmd/performance_alarm/main.go

mock_invt:
	go run ./cmd/mock_solarman/main.go

daily_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o daily ./cmd/daily_production/main.go

monthly_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o monthly ./cmd/monthly_production/main.go

pa_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o performance_alarm ./cmd/performance_alarm/main.go

mock_invt_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o mock_solarman ./cmd/mock_solarman/main.go

invt_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o solarman ./cmd/solarman/main.go

invt_alarm_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o solarman_alarm ./cmd/solarman_alarm/main.go

huawei:
	go run ./cmd/huawei/main.go

huawei_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o huawei ./cmd/huawei/main.go

prod_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o production ./cmd/production/main.go

kstar:
	go run ./cmd/kstar/main.go

kstar_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o kstar ./cmd/kstar/main.go

register_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o register ./cmd/register/main.go

alarm_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o alarm ./cmd/low_performance/main.go

growatt:
	go run ./cmd/growatt/main.go

growatt_build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o growatt ./cmd/growatt/main.go

plant_agg:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external" -o plant_agg ./cmd/plant_aggregate/main.go