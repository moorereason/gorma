package gorma

import (
	"text/template"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

// ModelWriter generate code for a goa application media types.
// Media types are data structures used to render the response bodies.
type ModelWriter struct {
	*codegen.GoGenerator
	ModelTmpl *template.Template
}

// NewModelWriter returns a contexts code writer.
// Media types contain the data used to render response bodies.
func NewModelWriter(filename string) (*ModelWriter, error) {
	cw := codegen.NewGoGenerator(filename)
	funcMap := cw.FuncMap
	funcMap["gotypedef"] = codegen.GoTypeDef
	funcMap["gotyperef"] = codegen.GoTypeRef
	funcMap["goify"] = codegen.Goify
	funcMap["gotypename"] = codegen.GoTypeName
	funcMap["gonative"] = codegen.GoNativeType
	funcMap["typeUnmarshaler"] = codegen.TypeUnmarshaler
	funcMap["typeMarshaler"] = codegen.MediaTypeMarshaler
	funcMap["recursiveValidate"] = codegen.RecursiveChecker
	funcMap["tempvar"] = codegen.Tempvar
	funcMap["demodel"] = deModel
	funcMap["modeldef"] = ModelDef
	funcMap["snake"] = camelToSnake
	funcMap["split"] = split
	funcMap["storagedef"] = StorageDef
	funcMap["lower"] = lower
	funcMap["title"] = titleCase
	funcMap["plural"] = plural
	funcMap["metaLookup"] = metaLookupTmpl
	funcMap["columns"] = GetAttributeColumns
	funcMap["timefuncs"] = MakeTimeFuncs

	modelTmpl, err := template.New("models").Funcs(funcMap).Parse(modelTmpl)
	if err != nil {
		return nil, err
	}
	w := ModelWriter{
		GoGenerator: cw,
		ModelTmpl:   modelTmpl,
	}
	return &w, nil
}

// Execute writes the code for the context types to the writer.
func (w *ModelWriter) Execute(mt *design.UserTypeDefinition) error {
	return w.ModelTmpl.Execute(w, mt)
}
