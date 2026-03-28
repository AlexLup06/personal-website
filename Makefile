DOCKER_COMPOSE_FILE = docker-compose.yaml
DOCKER_COMPOSE_DEV  = docker compose -f $(DOCKER_COMPOSE_FILE)

.PHONY: dev dev-tailwind connect-db migrate-up db-clean db-truncate-table db-reset admin-by-email \
	test test-up test-db-create test-migrate test-run test-down test-reset

dev:
	@if command -v tmux >/dev/null 2>&1; then \
		echo "Starting dev environment with tmux..."; \
		tmux new-session -d -s personal-website \
			'$(DOCKER_COMPOSE_DEV) up' \; \
			split-window -h \
			'cd frontend && npm run dev:tailwind' \; \
			attach; \
	else \
		echo ""; \
		echo "tmux not found."; \
		echo ""; \
		echo "Please run the following in two terminals:"; \
		echo "  1) $(DOCKER_COMPOSE_DEV) up"; \
		echo "  2) cd frontend && npm run dev:tailwind"; \
		echo ""; \
	fi

dev-tailwind:
	cd frontend && npm run dev:tailwind