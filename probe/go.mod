module github.com/thunkier/thunkmetrc/probe

go 1.25.4

require (
	github.com/joho/godotenv v1.5.1
	github.com/thunkier/thunkmetrc/mockserver v0.0.0-00010101000000-000000000000
	github.com/thunkier/thunkmetrc/tools v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/thunkier/thunkmetrc/tools => ../tools

replace github.com/thunkier/thunkmetrc/mockserver => ../mockserver
