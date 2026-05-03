#!/bin/bash
# EC2 User Data Script - IT KMITL Workshop
# ต้องแทนที่ <PLACEHOLDER> ก่อน deploy จริง
set -e

exec > /var/log/user-data.log 2>&1

# ============================
# Config - แก้ไขค่าเหล่านี้
# ============================
GITHUB_REPO="https://github.com/13D4C/Workshop-IT-KMITL-Cloud-Computing-Deployment-Website-Production.git"
ALB_DNS="<ALB_DNS_NAME>"
DB_HOST="<RDS_ENDPOINT>"
DB_USER="<DB_USERNAME>"
DB_PASSWORD="<DB_PASSWORD>"
DB_NAME="itkmitl"
JWT_SECRET="<CHANGE_THIS_TO_RANDOM_SECRET>"

# ============================
# Install Docker
# ============================
yum update -y
yum install -y docker git
systemctl enable docker
systemctl start docker

# Docker Compose plugin
mkdir -p /usr/local/lib/docker/cli-plugins
curl -SL "https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64" \
    -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose

usermod -aG docker ec2-user

# ============================
# Clone Repository
# ============================
mkdir -p /app
cd /app
git clone "${GITHUB_REPO}" .

# ============================
# Create Production .env
# ============================
cat > /app/.env.prod << EOF
ORIGIN=http://${ALB_DNS}
DB_HOST=${DB_HOST}
DB_PORT=5432
DB_USER=${DB_USER}
DB_PASSWORD=${DB_PASSWORD}
DB_NAME=${DB_NAME}
DB_SSLMODE=require
JWT_SECRET=${JWT_SECRET}
CORS_ORIGIN=http://${ALB_DNS}
EOF

# ============================
# Build & Start Services
# ============================
docker compose -f /app/deploy/docker-compose.prod.yml --env-file /app/.env.prod up -d --build

# ============================
# Systemd Service for auto-restart on reboot
# ============================
cat > /etc/systemd/system/workshop-app.service << 'UNIT'
[Unit]
Description=IT KMITL Workshop Application
After=docker.service
Requires=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=/app
ExecStart=/usr/bin/docker compose -f /app/deploy/docker-compose.prod.yml --env-file /app/.env.prod up -d
ExecStop=/usr/bin/docker compose -f /app/deploy/docker-compose.prod.yml down
Restart=on-failure

[Install]
WantedBy=multi-user.target
UNIT

systemctl daemon-reload
systemctl enable workshop-app

echo "✅ User data script completed successfully"
