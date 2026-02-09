module github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper

go 1.21
require (
    github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client v0.0.0
    golang.org/x/time v0.5.0
)
replace (
    github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client => ../client
)
