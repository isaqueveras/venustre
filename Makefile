templ:
	templ generate -include-timestamp

run: templ
	go run example/main.go