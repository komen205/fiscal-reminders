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

# Editar tópico ntfy
export NTFY_TOPIC="meu-topico-privado"

docker-compose up -d
```

### Systemd (Linux)

```bash
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders
go build -o fiscal-reminders .
sudo ./install.sh
```

### Manual

```bash
go build -o fiscal-reminders .
NTFY_TOPIC=meu-topico ./fiscal-reminders
```

## ⚙️ Configuração

```bash
cp config.example.json config.json
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

## 📱 Receber Notificações

### 1. Instalar App

- [Android (Play Store)](https://play.google.com/store/apps/details?id=io.heckel.ntfy)
- [iOS (App Store)](https://apps.apple.com/app/ntfy/id1625396347)
- [F-Droid](https://f-droid.org/packages/io.heckel.ntfy/)

### 2. Subscrever Tópico

Abre a app → "+" → Introduz o teu tópico (ex: `fiscal-reminders`)

### 3. Alternativas

```bash
# Browser
open https://ntfy.sh/SEU-TOPICO

# CLI
ntfy subscribe SEU-TOPICO

# cURL
curl -s ntfy.sh/SEU-TOPICO/json
```

## 🏠 Self-Hosted ntfy

```bash
export NTFY_SERVER=https://ntfy.meudominio.pt
export NTFY_TOPIC=fiscal-privado
docker-compose up -d
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

```bash
git clone https://github.com/komen205/fiscal-reminders.git
cd fiscal-reminders
go test ./...
go run main.go
```

## 📄 Licença

[MIT](LICENSE) - usa livremente!

## ⭐ Star History

Se este projeto te é útil, considera dar uma ⭐!

---

<p align="center">
  Feito com ❤️ em 🇵🇹 Portugal
</p>

