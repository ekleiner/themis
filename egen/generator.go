package main

import (
	"fmt"
	"strings"
)

const (
	packageTemplate = `package %s

/* AUTOMATICALLY GENERATED FROM %s - DO NOT EDIT */
`

	singleImportTemplate = `
import %s
`

	multiImportTemplate = `
import (
	%s
)
`

	singleIDTemplate = `
const %sID = iota
`

	multiIDTemplate = `
const (
	%sID = iota
	%sID
)
`

	noFieldTypeTemplate = `
type %s struct {
	errorLink
}
`

	multiFieldTypeTemplate = `
type %s struct {
	errorLink
	%s
}
`

	noFieldConstructor = `
func new%s%s() *%s {
	return &%s{
		errorLink: errorLink{id: %sID}}
}
`

	multiFieldConstructor = `
func new%s%s(%s) *%s {
	return &%s{
		errorLink: errorLink{id: %sID},
		%s}
}
`

	noArgMethod = `
func (e *%s) Error() string {
	return e.errorf(%q)
}
`

	multiArgMethod = `
func (e *%s) Error() string {
	return e.errorf(%q, %s)
}
`

	multiArgAndSnippetMethod = `
func (e *%s) Error() string {
%s
	return e.errorf(%q, %s)
}
`
)

func (e *errors) generate() {
	fmt.Printf(packageTemplate, e.Pkg, e.src)

	if len(e.Imports) > 1 {
		fmt.Printf(multiImportTemplate, strings.Join(e.Imports, "\n\t"))
	} else if len(e.Imports) > 0 {
		fmt.Printf(singleImportTemplate, e.Imports[0])
	}

	IDs := make([]string, len(e.Errors))
	for i, item := range e.Errors {
		IDs[i] = item.ID
	}

	if len(IDs) > 1 {
		fmt.Printf(multiIDTemplate, IDs[0], strings.Join(IDs[1:], "ID\n\t"))
	} else {
		fmt.Printf(singleIDTemplate, IDs[0])
	}

	for _, item := range e.Errors {
		if len(item.Fields) > 0 {
			maxLen := 0
			for _, field := range item.Fields {
				if maxLen < len(field.ID) {
					maxLen = len(field.ID)
				}
			}

			fields := make([]string, len(item.Fields))
			for i, field := range item.Fields {
				fields[i] = fmt.Sprintf("%s%s %s", field.ID, strings.Repeat(" ", maxLen-len(field.ID)), field.Type)
			}
			fmt.Printf(multiFieldTypeTemplate, item.ID, strings.Join(fields, "\n\t"))

			if maxLen < len("errorLink") {
				maxLen = len("errorLink")
			}

			args := []string{}
			group := []string{}
			pType := ""
			for i, field := range item.Fields {
				if field.Type != pType {
					if len(group) > 0 {
						args = append(args, fmt.Sprintf("%s %s", strings.Join(group, ", "), pType))
					}

					pType = field.Type
					group = []string{field.ID}
				} else {
					group = append(group, field.ID)
				}

				fields[i] = fmt.Sprintf("%s:%s %s", field.ID, strings.Repeat(" ", maxLen-len(field.ID)), field.ID)
			}

			if len(group) > 0 {
				args = append(args, fmt.Sprintf("%s %s", strings.Join(group, ", "), pType))
			}

			fmt.Printf(multiFieldConstructor,
				strings.ToUpper(string(item.ID[0])),
				item.ID[1:],
				strings.Join(args, ", "),
				item.ID,
				item.ID,
				item.ID,
				strings.Join(fields, ",\n\t\t"))
		} else {
			fmt.Printf(noFieldTypeTemplate, item.ID)

			fmt.Printf(noFieldConstructor,
				strings.ToUpper(string(item.ID[0])),
				item.ID[1:],
				item.ID,
				item.ID,
				item.ID)
		}

		if len(item.Args) > 0 {
			args := make([]string, len(item.Args))
			snippets := []string{}
			for i, arg := range item.Args {
				if len(arg.Field) > 0 {
					args[i] = fmt.Sprintf("e.%s", arg.Field)
				} else if len(arg.Expr) > 0 {
					args[i] = arg.Expr
				} else if arg.Snippet != nil {
					args[i] = arg.Snippet.Result

					code := ""
					lines := strings.Split(strings.TrimSpace(arg.Snippet.Code), "\n")
					for _, line := range lines {
						if len(strings.TrimSpace(line)) > 0 {
							code += fmt.Sprintf("\t%s\n", line)
						} else {
							code += "\n"
						}
					}

					snippets = append(snippets, code)
				}
			}

			if len(snippets) > 0 {
				fmt.Printf(multiArgAndSnippetMethod,
					item.ID,
					strings.Join(snippets, "\n\n"),
					item.Msg,
					strings.Join(args, ", "))
			} else {
				fmt.Printf(multiArgMethod, item.ID, item.Msg, strings.Join(args, ", "))
			}
		} else {
			fmt.Printf(noArgMethod, item.ID, item.Msg)
		}
	}
}
