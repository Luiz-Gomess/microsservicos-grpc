module github.com/Luiz-Gomess/microservices/order

go 1.2.3

require (
	github.com/Luiz-Gomess/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.6.0
)

replace github.com/Luiz-Gomess/microservices-proto/golang/order => /workspaces/microsservicos-grpc/microservices/order
