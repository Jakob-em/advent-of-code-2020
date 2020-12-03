
DAY = $(shell date +'%-d')
TODAYS_FILE = calendar/day_$(DAY)/task.go
TODAYS_TEST_FILE = calendar/day_$(DAY)/task_test.go

generate:
	go run generator/template.go -d $(DAY)

run:
	go run $(TODAYS_FILE)

test:
	go test $(TODAYS_FILE) $(TODAYS_TEST_FILE)