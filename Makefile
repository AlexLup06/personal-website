.PHONY: dev dev-tailwind

dev:
	@if command -v tmux >/dev/null 2>&1; then \
		echo "Starting dev environment with tmux..."; \
		tmux new-session -d -s personal-website \
			'docker compose -f docker-compose.dev.yaml up' \; \
			split-window -h \
			'cd frontend && npm run dev:tailwind' \; \
			attach; \
	else \
		echo ""; \
		echo "tmux not found."; \
		echo ""; \
		echo "Please run the following in two terminals:"; \
		echo "  1) docker compose -f docker-compose.dev.yaml up"; \
		echo "  2) cd frontend && npm run dev:tailwind"; \
		echo ""; \
	fi

dev-tailwind:
	cd frontend && npm run dev:tailwind