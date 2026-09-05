package deadline

import "time"

// Deadline represents a fiscal deadline
type Deadline struct {
	Name                    string
	Description             string
	Month                   int // 0 = monthly
	Day                     int
	AdjustToNextBusinessDay bool
	Priority                string // "urgent", "high", "default"
	Tags                    []string
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

	// Pagamentos por conta de IRS
	{
		Name:                    "💶 IRS - 1.º Pagamento por Conta",
		Description:             "Efetuar o 1.º pagamento por conta de IRS",
		Month:                   7,
		Day:                     20,
		AdjustToNextBusinessDay: true,
		Priority:                "high",
		Tags:                    []string{"irs", "pagamento-por-conta"},
	},
	{
		Name:                    "💶 IRS - 2.º Pagamento por Conta",
		Description:             "Efetuar o 2.º pagamento por conta de IRS",
		Month:                   9,
		Day:                     20,
		AdjustToNextBusinessDay: true,
		Priority:                "high",
		Tags:                    []string{"irs", "pagamento-por-conta"},
	},
	{
		Name:                    "💶 IRS - 3.º Pagamento por Conta",
		Description:             "Efetuar o 3.º pagamento por conta de IRS",
		Month:                   12,
		Day:                     20,
		AdjustToNextBusinessDay: true,
		Priority:                "high",
		Tags:                    []string{"irs", "pagamento-por-conta"},
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

// DateForYear returns the effective deadline date for a given year.
// Deadlines marked for business-day adjustment move from weekends to Monday.
func (d *Deadline) DateForYear(year int, location *time.Location) time.Time {
	date := time.Date(year, time.Month(d.Month), d.Day, 23, 59, 59, 0, location)

	if d.AdjustToNextBusinessDay {
		for date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			date = date.AddDate(0, 0, 1)
		}
	}

	return date
}
