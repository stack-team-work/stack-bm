#!/bin/bash

set -e

APP_NAME="stack-bm"
DEPLOY_DIR="/opt/${APP_NAME}"
BINARY="${APP_NAME}_linux_amd64"
BUILD_DIR="build"

echo "=== Deploying ${APP_NAME} ==="

if [ ! -f "${BUILD_DIR}/${BINARY}" ]; then
    echo "Binary not found. Building..."
    bash build.sh
fi

sudo mkdir -p ${DEPLOY_DIR}
sudo cp ${BUILD_DIR}/${BINARY} ${DEPLOY_DIR}/${APP_NAME}
sudo cp .env ${DEPLOY_DIR}/.env
sudo cp stack-bm.service /etc/systemd/system/

sudo chmod +x ${DEPLOY_DIR}/${APP_NAME}

sudo systemctl daemon-reload
sudo systemctl enable ${APP_NAME}
sudo systemctl restart ${APP_NAME}

echo "=== Deploy complete ==="
echo "Check status: sudo systemctl status ${APP_NAME}"
echo "View logs: sudo journalctl -u ${APP_NAME} -f"
