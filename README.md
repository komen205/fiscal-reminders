# 🔔 Fiscal Reminders

Notificações automáticas para prazos fiscais portugueses via [ntfy.sh](https://ntfy.sh).

## Prazos Monitorizados

### Declaração Trimestral Segurança Social
| Trimestre | Prazo |
|-----------|-------|
| Out-Dez | 31 Janeiro |
| Jan-Mar | 30 Abril |
| Abr-Jun | 31 Julho |
| Jul-Set | 31 Outubro |

### IVA Trimestral
| Trimestre | Prazo |
|-----------|-------|
| 1º (Jan-Mar) | 20 Maio |
| 2º (Abr-Jun) | 20 Agosto |
| 3º (Jul-Set) | 20 Novembro |
| 4º (Out-Dez) | 20 Fevereiro |

### Outros
- **Pagamento SegSoc**: Dia 20 de cada mês
- **IRS Anual**: 1 Abril - 30 Junho

## Instalação

### Opção 1: Docker (recomendado)

```bash
# Clonar repositório
git clone https://github.com/seu-user/fiscal-reminders.git
cd fiscal-reminders

# Configurar tópico ntfy (opcional)
export NTFY_TOPIC="meu-topico-privado"

# Executar
docker-compose up -d
```

### Opção 2: Systemd

```bash
# Clonar e construir
git clone https://github.com/seu-user/fiscal-reminders.git
cd fiscal-reminders
go build -o fiscal-reminders .

# Instalar (requer root)
sudo ./install.sh
```

### Opção 3: Manual

```bash
# Construir
go build -o fiscal-reminders .

# Executar
NTFY_TOPIC=meu-topico ./fiscal-reminders
```

## Configuração

Copiar o exemplo e editar:

```bash
cp config.example.json config.json
nano config.json
```

Exemplo `config.json`:

```json
{
  "ntfy_topic": "fiscal-reminders",
  "ntfy_server": "https://ntfy.sh",
  "check_interval_hours": 12,
  "days_before_alert": [7, 3, 1, 0]
}
```

| Campo | Descrição |
|-------|-----------|
| `ntfy_topic` | Nome do tópico ntfy |
| `ntfy_server` | Servidor ntfy (self-hosted ou ntfy.sh) |
| `check_interval_hours` | Frequência de verificação |
| `days_before_alert` | Dias antes do prazo para alertar |

## Receber Notificações

### App Móvel
1. Instalar app ntfy ([Android](https://play.google.com/store/apps/details?id=io.heckel.ntfy) / [iOS](https://apps.apple.com/app/ntfy/id1625396347))
2. Subscrever ao tópico: `fiscal-reminders`

### Browser
Abrir: https://ntfy.sh/fiscal-reminders

### CLI
```bash
ntfy subscribe fiscal-reminders
```

## Self-Hosted ntfy

Se tiveres o teu próprio servidor ntfy:

```bash
export NTFY_SERVER=https://ntfy.meudominio.pt
export NTFY_TOPIC=fiscal
docker-compose up -d
```

## Licença

MIT

