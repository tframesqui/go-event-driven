#Solution tools

.PHONY: test
test:
	@echo "fuck"

.PHONY: setup
setup:
	@docker compose run worker pip install -r requirements.txt

.PHONY: run_worker
run_worker:
	@docker compose run