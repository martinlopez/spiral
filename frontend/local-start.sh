docker build -t spiral-frontend:1.0.0 .
docker run -e ENV=local -p 3000:3000 -d spiral-frontend:1.0.0