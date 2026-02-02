# 🔔 Fiscal Reminders

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CI](https://github.com/komen205/fiscal-reminders/actions/workflows/ci.yml/badge.svg)](https://github.com/komen205/fiscal-reminders/actions/workflows/ci.yml)

> 🇵🇹 Notificações automáticas para prazos fiscais portugueses via [ntfy.sh](https://ntfy.sh)

**Nunca mais percas um prazo fiscal!** Recebe alertas automáticos no telemóvel dias antes de cada deadline.

---

## 🎯 Para Quem?

- 👨‍💻 **Freelancers** e trabalhadores independentes
- 🏢 **ENI** (Empresário em Nome Individual)
- 📊 **Contabilistas** que gerem múltiplos clientes
- 🇵🇹 Qualquer pessoa com obrigações fiscais em Portugal

## 📅 Prazos Monitorizados

### Declaração Trimestral Segurança Social

| Trimestre | Período | Prazo |
|-----------|---------|-------|
| Q4 | Out-Dez | 31 Janeiro |
| Q1 | Jan-Mar | 30 Abril |
| Q2 | Abr-Jun | 31 Julho |
| Q3 | Jul-Set | 31 Outubro |

### IVA Trimestral

| Trimestre | Período | Prazo |
|-----------|---------|-------|
| 1º | Jan-Mar | 20 Maio |
| 2º | Abr-Jun | 20 Agosto |
| 3º | Jul-Set | 20 Novembro |
| 4º | Out-Dez | 20 Fevereiro |

### Outros

| Obrigação | Prazo |
|-----------|-------|
| 💰 Pagamento SegSoc | Dia 20 de cada mês |
| 📝 IRS Anual | 1 Abril - 30 Junho |

## 📁 Project Structure

```
fiscal-reminders/
├── cmd/
│   └── fiscal-reminders/     # Application entry point
├── internal/
│   ├── config/               # Configuration loading
│   ├── deadline/             # Deadline definitions & checker
│   └── notifier/             # ntfy notification sender
├── deployments/
│   ├── docker/               # Dockerfile & docker-compose
│   └── systemd/              # Systemd service file
├── scripts/                  # Installation scripts
├── configs/                  # Configuration examples
└── .github/workflows/        # CI/CD
```

## 🚀 Quick Start

### Docker (recomendado)

```bash
docker run -d \
  --name fiscal-reminders \
  --restart unless-stopped \
  -e NTFY_TOPIC=meu-topico-secreto \
  ghcr.io/komen205/fiscal-reminders:latest
```

### Docker Compose

```bash
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders

export NTFY_TOPIC="meu-topico-privado"
docker-compose -f deployments/docker/docker-compose.yml up -d
```

### Make (desenvolvimento)

```bash
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders

make build     # Compila
make run       # Executa
make test      # Testes
make help      # Ver todos os comandos
```

### Systemd (Linux)

```bash
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders
sudo ./scripts/install.sh
```

## ⚙️ Configuração

```bash
cp configs/config.example.json config.json
```

```json
{
  "ntfy_topic": "fiscal-reminders",
  "ntfy_server": "https://ntfy.sh",
  "check_interval_hours": 12,
  "days_before_alert": [7, 3, 1, 0]
}
```

| Campo | Descrição | Default |
|-------|-----------|---------|
| `ntfy_topic` | Nome do tópico ntfy | `fiscal-reminders` |
| `ntfy_server` | Servidor ntfy | `https://ntfy.sh` |
| `check_interval_hours` | Frequência verificação (horas) | `12` |
| `days_before_alert` | Dias antes para alertar | `[7, 3, 1, 0]` |

### Environment Variables

```bash
NTFY_TOPIC=meu-topico
NTFY_SERVER=https://ntfy.sh
NTFY_USER=username        # opcional, para auth
NTFY_PASS=password        # opcional, para auth
```

## 📱 Receber Notificações

### 1. Instalar App

- [Android (Play Store)](https://play.google.com/store/apps/details?id=io.heckel.ntfy)
- [iOS (App Store)](https://apps.apple.com/app/ntfy/id1625396347)
- [F-Droid](https://f-droid.org/packages/io.heckel.ntfy/)

### 2. Subscrever Tópico

Abre a app → "+" → Introduz o teu tópico (ex: `fiscal-reminders`)

## 🛠️ Development

```bash
# Clone
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders

# Build
make build

# Run tests
make test

# Run with coverage
make test-cover

# Format code
make fmt

# Lint
make lint

# Build all platforms
make build-all
```

## 🗺️ Roadmap

- [ ] 📱 Integração Telegram
- [ ] 💬 Integração Discord  
- [ ] 📅 Export iCal (.ics)
- [ ] 🌐 Interface web com dashboard
- [ ] 🇧🇷 Suporte prazos Brasil
- [ ] 🔔 Notificações push nativas
- [ ] 📊 Histórico de notificações

## 🤝 Contribuir

Contribuições são bem-vindas! Vê [CONTRIBUTING.md](CONTRIBUTING.md).

## 📄 Licença

[MIT](LICENSE) - usa livremente!

---

<p align="center">
  Feito com ❤️ em 🇵🇹 Portugal
</p>
