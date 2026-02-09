module github.com/thunkier/thunkmetrc/tools

go 1.25.4

require (
	github.com/PuerkitoBio/goquery v1.11.0
	github.com/thunkier/thunkmetrc/probe v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/net v0.47.0 // indirect
)

replace github.com/thunkier/thunkmetrc/probe => ../probe
