docker run --name redis \
      -p 6379:6379 \
      -d redis:alpine redis-server

# docker exec -it redis redis-cli