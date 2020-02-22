package schema

import (
	"regexp"
	"strings"
)

type Check struct {
	Name       string
	Constraint string // length(phone) >= 10
	*Field
}

// ParseCheckConstraints parse schema check constraints
func (schema *Schema) ParseCheckConstraints() map[string]Check {
	var checks = map[string]Check{}
	for _, field := range schema.FieldsByDBName {
		if chk := field.TagSettings["CHECK"]; chk != "" {
			names := strings.Split(chk, ",")
			if len(names) > 1 && regexp.MustCompile("^[A-Za-z]+$").MatchString(names[0]) {
				checks[names[0]] = Check{Name: names[0], Constraint: strings.Join(names[1:], ","), Field: field}
			} else {
				name := schema.namer.CheckerName(schema.Table, field.DBName)
				checks[name] = Check{Name: name, Constraint: chk, Field: field}
			}
		}
	}
	return checks
}