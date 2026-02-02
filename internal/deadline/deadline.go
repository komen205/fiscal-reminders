package deadline

// Deadline represents a fiscal deadline
type Deadline struct {
	Name        string
	Description string
	Month       int // 0 = monthly
	Day         int
	Priority    string // "urgent", "high", "default"
	Tags        []string
}

// All contains all Portuguese fiscal deadlines
var All = []Deadline{
	// Declaração Trimestral Segurança Social
	{
		Name:        "📋 Declaração Trimestral SegSoc",
		Description: "Declarar rendimentos Out-Dez à Segurança Social",
		Month:       1, Day: 31,
		Priority: "high",
		Tags:     []string{"seguranca-social", "trimestral"},
	},
	{
		Name:        "📋 Declaração Trimestral SegSoc",
		Description: "Declarar rendimentos Jan-Mar à Segurança Social",
		Month:       4, Day: 30,
		Priority: "high",
		Tags:     []string{"seguranca-social", "trimestral"},
	},
	{
		Name:        "📋 Declaração Trimestral SegSoc",
		Description: "Declarar rendimentos Abr-Jun à Segurança Social",
		Month:       7, Day: 31,
		Priority: "high",
		Tags:     []string{"seguranca-social", "trimestral"},
	},
	{
		Name:        "📋 Declaração Trimestral SegSoc",
		Description: "Declarar rendimentos Jul-Set à Segurança Social",
		Month:       10, Day: 31,
		Priority: "high",
		Tags:     []string{"seguranca-social", "trimestral"},
	},

	// IVA Trimestral
	{
		Name:        "💶 Declaração IVA Trimestral",
		Description: "Entregar declaração IVA do 1º trimestre",
		Month:       5, Day: 20,
		Priority: "high",
		Tags:     []string{"iva", "trimestral"},
	},
	{
		Name:        "💶 Declaração IVA Trimestral",
		Description: "Entregar declaração IVA do 2º trimestre",
		Month:       8, Day: 20,
		Priority: "high",
		Tags:     []string{"iva", "trimestral"},
	},
	{
		Name:        "💶 Declaração IVA Trimestral",
		Description: "Entregar declaração IVA do 3º trimestre",
		Month:       11, Day: 20,
		Priority: "high",
		Tags:     []string{"iva", "trimestral"},
	},
	{
		Name:        "💶 Declaração IVA Trimestral",
		Description: "Entregar declaração IVA do 4º trimestre",
		Month:       2, Day: 20,
		Priority: "high",
		Tags:     []string{"iva", "trimestral"},
	},

	// Pagamento contribuições SegSoc (mensal)
	{
		Name:        "💳 Pagamento SegSoc",
		Description: "Pagar contribuições Segurança Social",
		Month:       0, Day: 20, // Month 0 = every month
		Priority: "default",
		Tags:     []string{"seguranca-social", "pagamento"},
	},

	// IRS Anual
	{
		Name:        "📝 IRS Anual - Início",
		Description: "Período de entrega do IRS começa",
		Month:       4, Day: 1,
		Priority: "default",
		Tags:     []string{"irs", "anual"},
	},
	{
		Name:        "📝 IRS Anual - Fim",
		Description: "Último dia para entregar IRS",
		Month:       6, Day: 30,
		Priority: "urgent",
		Tags:     []string{"irs", "anual"},
	},
}

// IsMonthly returns true if this is a monthly recurring deadline
func (d *Deadline) IsMonthly() bool {
	return d.Month == 0
}

// HasTag checks if deadline has a specific tag
func (d *Deadline) HasTag(tag string) bool {
	for _, t := range d.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

