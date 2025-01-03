all: typescript-to-javascript compress-static-files webpack tailwind build run

typescript-to-javascript:
	@echo "Compiling TypeScript to JavaScript..."
	@cd frontend && npx tsc 
	@echo "TypeScript compiled to JavaScript!"

webpack:
	@echo "Running Webpack..."
	@cd frontend && npm run prod:webpack
	@echo "Webpack finished!"

tailwind:
	@echo "Compiling Tailwind CSS..."
	@cd frontend && npx tailwindcss build src/css/style.css -o src/css/tailwind.css
	@echo "Tailwind CSS compiled!"

compress-static-files:
	@echo "Compressing static files..."
	@gzip -k -f frontend/src/js/build/*.js
	@gzip -k -f frontend/src/css/tailwind.css
	@find ./frontend/public -type f ! -path "./frontend/public/fonts/*" ! -name "*.gz" -exec sh -c 'gzip -c "$1" > "$1.gz"' _ {} \;
	@echo "Static files compressed!"
build:
	@echo "Building the project..."
	@go build -o bin/main main.go
	@echo "Project built!"

run:
	@echo "Running the project..."
	@./bin/main
	@echo "Project running!"