FROM node:20-slim

# Install MongoDB, curl, and gnupg
RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    lsb-release \
    && curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | gpg --dearmor -o /usr/share/keyrings/mongodb-server-7.0.gpg \
    && echo "deb [ signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] http://repo.mongodb.org/apt/debian bookworm/mongodb-org/7.0 main" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list \
    && apt-get update \
    && apt-get install -y mongodb-org \
    && rm -rf /var/lib/apt/lists/*

# Create MongoDB data directory
RUN mkdir -p /data/db

WORKDIR /app

# Install app dependencies
COPY package*.json ./
RUN npm install

# Copy source code and build Vite frontend
COPY . .
RUN npm run build

# Create a startup script to launch MongoDB and Node
RUN echo '#!/bin/sh\n\
mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db\n\
node seed.js\nnode server.js' > /app/start.sh && chmod +x /app/start.sh

EXPOSE 8080

ENV PORT=8080
ENV MONGO_URI=mongodb://127.0.0.1:27017/neo_monolith_db

CMD ["/app/start.sh"]