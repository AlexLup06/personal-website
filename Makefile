all: npm-install webpack tailwind compress-static-files

npm-install:
	@echo "Installing npm packages..."
	@cd frontend && npm install
	@echo "npm packages installed!"

webpack:
	@echo "Running Webpack..."
	@cd frontend && npm run prod:webpack
	@echo "Webpack finished!"
	
tailwind:
	@echo "Compiling Tailwind CSS..."
	@cd frontend && npx tailwindcss build -i src/css/style.css -o src/css/tailwind.css
	@echo "Tailwind CSS compiled!"

compress-static-files:
	@echo "Compressing static files..."
	@gzip -k -f frontend/src/css/tailwind.css -c > frontend/public/tailwind.css.gz
	@gzip -k -f frontend/src/js/bundle.js -c > frontend/public/bundle.js.gz
	@gzip -k -f frontend/src/js/htmx.min.js -c > frontend/public/htmx.min.js.gz
	@gzip -k -f frontend/src/js/alpine.min.js -c > frontend/public/alpine.min.js.gz
	@echo "Static files compressed!"

run-tailwind:
	@echo "Running Tailwind CSS..."
	@cd frontend && npx tailwindcss build -i src/css/style.css -o src/css/tailwind.css --watch
	@echo "Tailwind CSS running!"

dev:
	@echo "Running Tailwind CSS..."
	@cd frontend && npx tailwindcss build -i src/css/style.css -o src/css/tailwind.css --watch
	@echo "Tailwind CSS running!"