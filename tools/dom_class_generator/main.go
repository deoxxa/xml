package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/segmentio/go-snakecase"
)

func up(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

var trep = map[string]string{
	"DOMString": "string",
	"short":     "int16",
	"long":      "int32",
	"boolean":   "bool",
	"void":      "",
}

var crep = map[string]string{
	"NamedNodeMap": "NamedNodeMap",
	"NodeList":     "NodeList",
}

func formatType(s string, i bool) string {
	if _, ok := trep[s]; ok {
		return trep[s]
	} else if i {
		return "I" + s
	} else if _, ok := crep[s]; ok {
		return crep[s]
	} else {
		return "*" + s
	}
}

var dval = map[string]string{
	"boolean":   "false",
	"DOMString": "\"\"",
	"void":      "",
	"short":     "0",
}

func defaultValue(s string) string {
	if _, ok := dval[s]; ok {
		return dval[s]
	} else {
		return "nil"
	}
}

func formatName(s string) string {
	if s == "type" {
		return "theType"
	} else {
		return s
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("data length: %d", len(data))

	p := &Parser{
		idl:    &IDL{},
		Buffer: string(data),
	}

	p.Init()

	if err := p.Parse(); err != nil {
		log.Fatal(err)
	}

	p.Execute()

	log.Printf("writing %d classes", len(p.idl.Interfaces))

	for _, i := range p.idl.Interfaces {
		code := ""

		code += fmt.Sprintf("package dom\n\n")

		code += fmt.Sprintf("type I%s interface {\n", i.Name)

		if i.Inherits != "" {
			code += fmt.Sprintf("  I%s\n\n", i.Inherits)
		}

		if len(i.Attributes) > 0 {
			code += fmt.Sprintf("  // accessors\n")

			for _, a := range i.Attributes {
				code += fmt.Sprintf("  Get%s() %s\n", up(a.Name), formatType(a.Type.Type, true))
				code += fmt.Sprintf("  Set%s(%s %s)\n", up(a.Name), formatName(a.Name), formatType(a.Type.Type, true))
			}
		} else {
			code += fmt.Sprintf("  // no accessors\n")
		}

		if len(i.Functions) > 0 {
			code += fmt.Sprintf("  // spec-defined functions\n")

			for _, f := range i.Functions {
				code += fmt.Sprintf("  %s(", up(f.Name))

				for n, a := range f.Arguments {
					if n != 0 {
						code += fmt.Sprintf(", ")
					}

					code += fmt.Sprintf("%s %s", a.Name, formatType(a.Type.Type, true))
				}

				code += fmt.Sprintf(") %s\n", formatType(f.Type.Type, true))
			}
		} else {
			code += fmt.Sprintf("  // no spec-defined functions\n")
		}

		code += fmt.Sprintf("}\n\n")

		if len(i.Constants) > 0 {
			code += fmt.Sprintf("const (\n")

			for _, c := range i.Constants {
				code += fmt.Sprintf("  %s = %s\n", c.Name, c.Value)
			}

			code += fmt.Sprintf(")\n\n")
		}

		code += fmt.Sprintf("type %s struct {\n", i.Name)

		if i.Inherits != "" {
			code += fmt.Sprintf("  %s\n", i.Inherits)
		}

		if len(i.Attributes) > 0 {
			for _, a := range i.Attributes {
				code += fmt.Sprintf("  %s %s\n", formatName(a.Name), formatType(a.Type.Type, true))
			}
		}

		code += fmt.Sprintf("}\n\n")

		if len(i.Attributes) > 0 {
			for _, a := range i.Attributes {
				code += fmt.Sprintf("func (%s %s) Get%s() %s {\n  return %s.%s\n}\n\n", strings.ToLower(i.Name[0:1]), i.Name, up(a.Name), formatType(a.Type.Type, true), strings.ToLower(i.Name[0:1]), formatName(a.Name))
				code += fmt.Sprintf("func (%s %s) Set%s(%s %s) {\n  %s.%s = %s\n}\n\n", strings.ToLower(i.Name[0:1]), i.Name, up(a.Name), formatName(a.Name), formatType(a.Type.Type, true), strings.ToLower(i.Name[0:1]), formatName(a.Name), formatName(a.Name))
			}
		}

		if len(i.Functions) > 0 {
			for _, f := range i.Functions {
				code += fmt.Sprintf("func (%s %s) %s(", strings.ToLower(i.Name[0:1]), i.Name, up(f.Name))

				for n, a := range f.Arguments {
					if n != 0 {
						code += fmt.Sprintf(", ")
					}

					code += fmt.Sprintf("%s %s", a.Name, formatType(a.Type.Type, true))
				}

				code += fmt.Sprintf(") %s {\n  panic(\"unimplemented\") // TODO\n\n  return %s\n}\n\n", formatType(f.Type.Type, true), defaultValue(f.Type.Type))

				fmt.Printf("%s::%s\n", i.Name, up(f.Name))
			}
		}

		if err := ioutil.WriteFile("./dom/"+snakecase.Snakecase(i.Name)+".go", []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
	}
}
