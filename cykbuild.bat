@echo off
swag init -g .\apiservice\route.go

:: go run main.go -migrate
go run main.go