
DAY := $(shell date +'%-d')

generate:
	go run generator/template.go -d $(DAY)

run:
	go run calendar/day_${DAY}/task.go
