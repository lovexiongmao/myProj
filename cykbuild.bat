@echo off
swag init -g .\api_service\api_service.go

:: go run main.go -migrate
go run main.go