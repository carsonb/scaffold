// Generated by ego.
// DO NOT EDIT

//line model.ego:1

package model

import "fmt"
import "html"
import "io"
import "context"

import (
	"github.com/boourns/scaffold/ast"
	"github.com/boourns/scaffold/sqlgen"
	"github.com/boourns/scaffold/util"
	"strings"
)

func modelTemplate(w io.Writer, m *ast.Model) {

//line model.ego:13
	_, _ = io.WriteString(w, "\n\npackage ")
//line model.ego:14
	_, _ = fmt.Fprint(w, strings.ToLower(m.Package))
//line model.ego:15
	_, _ = io.WriteString(w, "\n\nimport (\n\t\"github.com/boourns/dbutil\"\n  \"database/sql\"\n  \"fmt\"\n)\n\nfunc sqlFieldsFor")
//line model.ego:22
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:22
	_, _ = io.WriteString(w, "() string {\n  return \"")
//line model.ego:23
	_, _ = fmt.Fprint(w, util.StringJoin(m.FieldSlice(), ", "))
//line model.ego:23
	_, _ = io.WriteString(w, "\"\n}\n\nfunc load")
//line model.ego:26
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:26
	_, _ = io.WriteString(w, "(rows *sql.Rows) (*")
//line model.ego:26
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:26
	_, _ = io.WriteString(w, ", error) {\n\tret := ")
//line model.ego:27
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:27
	_, _ = io.WriteString(w, "{}\n\n\terr := rows.Scan(")
//line model.ego:29
	_, _ = fmt.Fprint(w, fieldString("&ret.", m.FieldSlice(), ""))
//line model.ego:29
	_, _ = io.WriteString(w, ")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &ret, nil\n}\n\nfunc Select")
//line model.ego:36
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:36
	_, _ = io.WriteString(w, "(tx dbutil.DBLike, cond string, condFields ...interface{}) ([]*")
//line model.ego:36
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:36
	_, _ = io.WriteString(w, ", error) {\n  ret := []*")
//line model.ego:37
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:37
	_, _ = io.WriteString(w, "{}\n  sql := fmt.Sprintf(\"SELECT %s from ")
//line model.ego:38
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:38
	_, _ = io.WriteString(w, " %s\", sqlFieldsFor")
//line model.ego:38
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:38
	_, _ = io.WriteString(w, "(), cond)\n\trows, err := tx.Query(sql, condFields...)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n    item, err := load")
//line model.ego:44
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:44
	_, _ = io.WriteString(w, "(rows)\n    if err != nil {\n      return nil, err\n    }\n    ret = append(ret, item)\n  }\n  rows.Close()\n  return ret, nil\n}\n\nfunc (s *")
//line model.ego:54
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:54
	_, _ = io.WriteString(w, ") Update(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(fmt.Sprintf(\"UPDATE ")
//line model.ego:55
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:55
	_, _ = io.WriteString(w, " SET ")
//line model.ego:55
	_, _ = fmt.Fprint(w, fieldString("", m.FieldSlice(), "=?"))
//line model.ego:55
	_, _ = io.WriteString(w, " WHERE ")
//line model.ego:55
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:55
	_, _ = io.WriteString(w, ".ID = ?\", ))\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n    params := []interface{}{")
//line model.ego:61
	_, _ = fmt.Fprint(w, fieldString("s.", m.FieldSlice(), ""))
//line model.ego:61
	_, _ = io.WriteString(w, "}\n    params = append(params, s.ID)\n\n\t\t_, err = stmt.Exec(params...)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n    return nil\n}\n\nfunc (s *")
//line model.ego:72
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:72
	_, _ = io.WriteString(w, ") Insert(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(\"INSERT INTO ")
//line model.ego:73
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:73
	_, _ = io.WriteString(w, "(")
//line model.ego:73
	_, _ = fmt.Fprint(w, fieldString("", m.FieldSliceWithoutID(), ""))
//line model.ego:73
	_, _ = io.WriteString(w, ") VALUES(")
//line model.ego:73
	_, _ = fmt.Fprint(w, util.QuestionMarks(len(m.FieldSliceWithoutID())))
//line model.ego:73
	_, _ = io.WriteString(w, ")\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tresult, err := stmt.Exec(")
//line model.ego:78
	_, _ = fmt.Fprint(w, fieldString("s.", m.FieldSliceWithoutID(), ""))
//line model.ego:78
	_, _ = io.WriteString(w, ")\n\t\tif err != nil {\n\t\t\treturn err\n    }\n\n    s.ID, err = result.LastInsertId()\n    if err != nil {\n      return err\n    }\n\t  return nil\n}\n\nfunc (s *")
//line model.ego:90
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:90
	_, _ = io.WriteString(w, ") Delete(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(\"DELETE FROM ")
//line model.ego:91
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:91
	_, _ = io.WriteString(w, " WHERE ID = ?\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\t_, err = stmt.Exec(s.ID)\n\t\tif err != nil {\n\t\t\treturn err\n    }\n\n\t  return nil\n}\n\nfunc Create")
//line model.ego:104
	_, _ = fmt.Fprint(w, m.Name)
//line model.ego:104
	_, _ = io.WriteString(w, "Table(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(`")
//line model.ego:105
	_, _ = fmt.Fprint(w, sqlgen.CreateTable(m))
//line model.ego:105
	_, _ = io.WriteString(w, "`)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\t_, err = stmt.Exec()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t  return nil\n}\n")
//line model.ego:116
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
