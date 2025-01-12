// Generated by ego.
// DO NOT EDIT

//line model.ts.ego:1

package model

import "fmt"
import "html"
import "io"
import "context"

import (
	"github.com/boourns/scaffold/ast"
	"strings"
	//"github.com/boourns/scaffold/util"
	//"github.com/boourns/scaffold/sqlgen"
	//"fmt"
)

//arrays and slices aren't handled properly.
func tsTypeForField(f ast.Field) string {
	switch f.Type {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "time.Time":
		return "Date"
	default:
		return tsTypeForSQLType(f)
	}
}

func tsTypeForSQLType(f ast.Field) string {
	override := f.Override("sqlType", "")
	switch override {
	case "DATETIME":
		return "Date"
	case "TEXT":
		return "string"
	default:
		return "any"
	}
}

func fieldNameString(m *ast.Model) string {
	return strings.Join(m.FieldSlice(), ",")
}

func modelTemplateTS(w io.Writer, m *ast.Model) {

//line model.ts.ego:46
	_, _ = io.WriteString(w, "\nclass ")
//line model.ts.ego:46
	_, _ = fmt.Fprint(w, m.Name)
//line model.ts.ego:46
	_, _ = io.WriteString(w, " {\n    ")
//line model.ts.ego:47
	for _, field := range m.Fields {
//line model.ts.ego:48
		_, _ = io.WriteString(w, "\n    ")
//line model.ts.ego:48
		_, _ = fmt.Fprint(w, field.Name)
//line model.ts.ego:48
		_, _ = io.WriteString(w, ": ")
//line model.ts.ego:48
		_, _ = fmt.Fprint(w, tsTypeForField(field))
//line model.ts.ego:48
		_, _ = io.WriteString(w, " | undefined\n    ")
//line model.ts.ego:49
	}
//line model.ts.ego:50
	_, _ = io.WriteString(w, "\n    SelectAll: string = \"SELECT (")
//line model.ts.ego:50
	_, _ = fmt.Fprint(w, fieldNameString(m))
//line model.ts.ego:50
	_, _ = io.WriteString(w, ") FROM ")
//line model.ts.ego:50
	_, _ = fmt.Fprint(w, m.Name)
//line model.ts.ego:50
	_, _ = io.WriteString(w, "\"\n    SelectByID: string = \"SELECT (")
//line model.ts.ego:51
	_, _ = fmt.Fprint(w, fieldNameString(m))
//line model.ts.ego:51
	_, _ = io.WriteString(w, ") FROM ")
//line model.ts.ego:51
	_, _ = fmt.Fprint(w, m.Name)
//line model.ts.ego:51
	_, _ = io.WriteString(w, " WHERE ID=?\"\n}\nexport default ")
//line model.ts.ego:53
	_, _ = fmt.Fprint(w, m.Name)
//line model.ts.ego:54
	_, _ = io.WriteString(w, "\n")
//line model.ts.ego:54
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
