#!/bin/bash
set -e

echo "🔔 Installing Fiscal Reminders..."

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "❌ Please run as root (sudo ./scripts/install.sh)"
    exit 1
fi

# Build if binary doesn't exist
if [ ! -f "bin/fiscal-reminders" ]; then
    echo "📦 Building..."
    make build
fi

# Create working directory
mkdir -p /var/lib/fiscal-reminders

# Copy binary
cp bin/fiscal-reminders /usr/local/bin/
chmod +x /usr/local/bin/fiscal-reminders

# Copy config example
if [ -f "configs/config.example.json" ]; then
    cp configs/config.example.json /var/lib/fiscal-reminders/config.json
fi

# Install systemd service
cp deployments/systemd/fiscal-reminders.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable fiscal-reminders

echo ""
echo "✅ Installation complete!"
echo ""
echo "📝 Configure: /var/lib/fiscal-reminders/config.json"
echo "🚀 Start:     sudo systemctl start fiscal-reminders"
echo "📊 Status:    sudo systemctl status fiscal-reminders"
echo "📜 Logs:      sudo journalctl -u fiscal-reminders -f"
