docker build -t spiral-backend:1.0.0 .
docker run -e ENV=local -p 8080:8080 -d spiral-backend:1.0.0