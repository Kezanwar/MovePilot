package form_repo

type FormData struct {
	Steps []Step `json:"steps"`
	Hero  *Hero  `json:"hero,omitempty"`
}

type Step struct {
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Fields      []Field    `json:"fields"`
	Condition   *Condition `json:"condition,omitempty"` // Optional condition object
}

type Field struct {
	UUID         string      `json:"uuid"`
	Type         string      `json:"type"` // "text", "email", "textarea", "select", "radio", "checkbox", "number", "date", etc.
	Name         string      `json:"name"`
	Label        string      `json:"label"`
	Placeholder  string      `json:"placeholder,omitempty"`
	Required     bool        `json:"required"`
	Validation   *Validation `json:"validation,omitempty"`
	Options      []Option    `json:"options,omitempty"`
	DefaultValue string      `json:"default_value,omitempty"`
	Condition    *Condition  `json:"condition,omitempty"` // Optional condition object
}

type Condition struct {
	Operator   string     `json:"operator"`   // "AND" or "OR"
	Conditions []CondRule `json:"conditions"` // Array of condition rules
}

type CondRule struct {
	Type     string      `json:"type"`     // "field" or "step"
	UUID     string      `json:"uuid"`     // UUID of the field or step to check
	Operator string      `json:"operator"` // "equals", "notEquals", "contains", "greaterThan", "lessThan", "isEmpty", "isNotEmpty"
	Value    interface{} `json:"value"`    // Value to compare against
}

type Validation struct {
	// String validations
	MinLength *int    `json:"min_length,omitempty"` // .min(n)
	MaxLength *int    `json:"max_length,omitempty"` // .max(n)
	Length    *int    `json:"length,omitempty"`     // .length(n)
	Matches   *string `json:"matches,omitempty"`    // .matches(regex)
	Email     *bool   `json:"email,omitempty"`      // .email()
	URL       *bool   `json:"url,omitempty"`        // .url()
	UUID      *bool   `json:"uuid,omitempty"`       // .uuid()

	// Number validations
	Min      *float64 `json:"min,omitempty"`       // .min(n)
	Max      *float64 `json:"max,omitempty"`       // .max(n)
	LessThan *float64 `json:"less_than,omitempty"` // .lessThan(n)
	MoreThan *float64 `json:"more_than,omitempty"` // .moreThan(n)
	Positive *bool    `json:"positive,omitempty"`  // .positive()
	Negative *bool    `json:"negative,omitempty"`  // .negative()
	Integer  *bool    `json:"integer,omitempty"`   // .integer()

	// Array validations (for multi-select, checkboxes)
	MinItems *int `json:"min_items,omitempty"` // .min(n)
	MaxItems *int `json:"max_items,omitempty"` // .max(n)

	// Date validations
	MinDate *string `json:"min_date,omitempty"` // .min(date) - ISO string
	MaxDate *string `json:"max_date,omitempty"` // .max(date) - ISO string
}

type Option struct {
	UUID  string `json:"uuid"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type Hero struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}
