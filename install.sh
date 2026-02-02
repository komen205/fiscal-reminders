#!/bin/bash
set -e

# Fiscal Reminders - Installation Script
# Run as root or with sudo

APP_NAME="fiscal-reminders"
INSTALL_DIR="/opt/$APP_NAME"
SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"

echo "🔔 Installing Fiscal Reminders..."

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root or with sudo"
    exit 1
fi

# Create user if not exists
if ! id "$APP_NAME" &>/dev/null; then
    useradd -r -s /bin/false $APP_NAME
    echo "✅ Created user: $APP_NAME"
fi

# Create installation directory
mkdir -p $INSTALL_DIR
chown $APP_NAME:$APP_NAME $INSTALL_DIR

# Check if binary exists
if [ ! -f "./$APP_NAME" ]; then
    echo "Building from source..."
    if command -v go &>/dev/null; then
        go build -o $APP_NAME .
    else
        echo "❌ Go not installed. Please build the binary first or install Go."
        exit 1
    fi
fi

# Copy files
cp ./$APP_NAME $INSTALL_DIR/
cp ./config.json $INSTALL_DIR/
chown -R $APP_NAME:$APP_NAME $INSTALL_DIR
chmod +x $INSTALL_DIR/$APP_NAME

# Install systemd service
cp ./$APP_NAME.service $SERVICE_FILE

# Reload systemd and enable service
systemctl daemon-reload
systemctl enable $APP_NAME
systemctl start $APP_NAME

echo ""
echo "✅ Installation complete!"
echo ""
echo "📋 Commands:"
echo "   Status:  systemctl status $APP_NAME"
echo "   Logs:    journalctl -u $APP_NAME -f"
echo "   Stop:    systemctl stop $APP_NAME"
echo "   Restart: systemctl restart $APP_NAME"
echo ""
echo "📱 Subscribe to notifications:"
echo "   ntfy subscribe fiscal-reminders"
echo "   Or open: https://ntfy.sh/fiscal-reminders"
echo ""

