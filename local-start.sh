#Starting backend
echo "Starting backend..."
cd backend
sh local-start.sh
cd ..

#Starting frontend
echo "Starting frontend..."
cd frontend
sh local-start.sh
cd ..