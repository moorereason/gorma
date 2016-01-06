package gorma

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/qor/inflection"
	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
)

const META_NAMESPACE = "github.com/bketelsen/gorma"

// timeFields stores a list of Time fields found in the model.  We update this
// in MakeModelDef and use it later to generate helper functions.
var timeFields []string

// titleCase converts a string to Title case.
func titleCase(s string) string {
	return strings.Title(s)
}

type Field struct {
	Column  string
	Coltype string
}

func GetAttributeColumns(att *design.AttributeDefinition) []Field {
	var columns []Field
	if o := att.Type.ToObject(); o != nil {
		o.IterateAttributes(func(n string, catt *design.AttributeDefinition) error {
			f := Field{
				Column:  n,
				Coltype: codegen.GoNativeType(catt.Type),
			}
			columns = append(columns, f)
			return nil
		})
	}

	return columns
}

// camelToSnake converts a given string to snake case.
func camelToSnake(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

// modelDir is the path to the directory where the schema controller is generated.
func modelDir() string {
	return filepath.Join(codegen.OutputDir, "models")
}

// deModel removes the word "Model" from the string.
func deModel(s string) string {
	return strings.Replace(s, "Model", "", -1)
}

// lower returns the string in lowercase.
func lower(s string) string {
	return strings.ToLower(s)
}

// upper returns the string in upper case.
func upper(s string) string {
	return strings.ToUpper(s)
}

// metaLookup is a helper function to lookup gorma-namespaced metadata keys in a
// case-insensitive way.
func metaLookup(md design.MetadataDefinition, hashtag string) (result string, ok bool) {
	needle := strings.ToLower(META_NAMESPACE + hashtag)
	for k, v := range md {
		k = strings.ToLower(k)
		if k == needle {
			return v, true
		}
	}

	return
}

// metaLookupTmpl is a helpful wrapper around metaLookup for use in templates.
func metaLookupTmpl(md design.MetadataDefinition, hashtag string) string {
	result, _ := metaLookup(md, hashtag)
	return result
}

// StorageDef creates the storage interface that will be used
// in place of a concrete type for testability.
func StorageDef(res *design.UserTypeDefinition) string {
	var associations string
	if assoc, ok := metaLookup(res.Metadata, "#many2many"); ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + "List" + pieces[0] + "(context.Context, int) []" + pieces[1] + "\n"
			associations = associations + "Add" + pieces[1] + "(context.Context, int, int) (error)\n"
			associations = associations + "Delete" + pieces[1] + "(context.Context, int, int) error \n"
		}
	}
	return associations
}

// includeForeignKey adds foreign key relations to the struct being
// generated.
func includeForeignKey(res *design.AttributeDefinition) string {
	var associations string
	if assoc, ok := metaLookup(res.Metadata, "#belongsto"); ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			associations = associations + child + "ID int\n"

		}
	}
	return associations
}

// plural returns the plural version of a word.
func plural(s string) string {
	return inflection.Plural(s)
}

// includeChildren adds the fields to a struct represented
// in a has-many relationship.
func includeChildren(res *design.AttributeDefinition) string {
	var associations string
	if assoc, ok := metaLookup(res.Metadata, "#hasmany"); ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			associations = associations + inflection.Plural(child) + " []" + child + "\n"
		}
	}
	if assoc, ok := metaLookup(res.Metadata, "#hasone"); ok {
		children := strings.Split(assoc, ",")
		for _, child := range children {
			associations = associations + child + " " + child + "\n"
			associations = associations + child + "ID " + "*sql.NullInt64\n"
		}
	}
	return associations
}

// includeMany2Many returns the appropriate struct tags
// for a m2m relationship in gorm.
func includeMany2Many(res *design.AttributeDefinition) string {
	var associations string
	if assoc, ok := metaLookup(res.Metadata, "#many2many"); ok {
		children := strings.Split(assoc, ",")

		for _, child := range children {
			pieces := strings.Split(child, ":")
			associations = associations + pieces[0] + "\t []" + pieces[1] + "\t" + "`gorm:\"many2many:" + pieces[2] + ";\"`\n"
		}
	}
	return associations
}

// includeAuthboss returns the tags required to implement authboss storage.
// Currently experimental and quite unfinished.
func includeAuthboss(res *design.AttributeDefinition) string {
	if _, ok := metaLookup(res.Metadata, "#authboss"); ok {
		fields := `	// Auth
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time
	`
		return fields
	}
	return ""
}

// split splits a string by separater `sep`.
func split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// includeTimeStamps returns the timestamp fields if "skipts" isn't set.
func includeTimeStamps(res *design.AttributeDefinition) string {
	var ts string
	if _, ok := metaLookup(res.Metadata, "#skipts"); ok {
		ts = ""
	} else {
		ts = "CreatedAt time.Time\nUpdatedAt time.Time\nDeletedAt *time.Time\n"
	}
	return ts
}

// isDateTime takes an attribute definition and returns whether there is a
// Format("date-time") validation associated with it.
func isDateTimeFormat(res *design.AttributeDefinition) bool {
	if res.Type.Kind() == design.StringKind && len(res.Validations) > 0 {
		for i := range res.Validations {
			switch v := res.Validations[i].(type) {
			case *design.FormatValidationDefinition:
				if v.Format == "date-time" {
					return true
				}
			}
		}
	}
	return false
}

// ModelDef is the main function to create a struct definition.
func ModelDef(res *design.UserTypeDefinition) string {
	var buffer bytes.Buffer
	timeFields = make([]string, 0)
	def := res.Definition()
	t := def.Type
	switch actual := t.(type) {
	case design.Object:
		actual = setupIDAttribute(actual, res)

		buffer.WriteString("struct {\n")
		keys := make([]string, len(actual))
		i := 0
		for n := range actual {
			keys[i] = n
			i++
		}
		sort.Strings(keys)
		for _, name := range keys {
			codegen.WriteTabs(&buffer, 1)

			var typedef string
			isTime := isDateTimeFormat(actual[name])
			if isTime {
				typedef = "*time.Time"
			} else {
				typedef = codegen.GoTypeDef(actual[name], 1, true, true)
			}

			fname := codegen.Goify(name, true)
			if isTime {
				timeFields = append(timeFields, fname)
				fname += "Time"
			}

			var tags string
			var omit string
			var gorm, sql string
			if !def.IsRequired(name) {
				omit = ",omitempty"
			}
			if val, ok := metaLookup(actual[name].Metadata, "#gormtag"); ok {
				gorm = fmt.Sprintf(" gorm:\"%s\"", val)
			} else if isTime {
				gorm = fmt.Sprintf(" gorm:\"column:%s\"", strings.ToLower(name))
			}
			if val, ok := metaLookup(actual[name].Metadata, "#sqltag"); ok {
				sql = fmt.Sprintf(" sql:\"%s\"", val)
			}
			tags = fmt.Sprintf(" `json:\"%s%s\"%s%s`", name, omit, gorm, sql)
			desc := actual[name].Description
			if desc != "" {
				desc = fmt.Sprintf("// %s\n", desc)
			}
			buffer.WriteString(fmt.Sprintf("%s%s %s%s\n", desc, fname, typedef, tags))
		}

		for k, v := range genfuncs {
			s := v(def)
			if s != "" {
				buffer.WriteString(fmt.Sprintf("%s%s", k, s))
			}
		}

		codegen.WriteTabs(&buffer, 0)
		buffer.WriteString("}")
		return buffer.String()
	default:
		panic("gorma bug: unexpected data structure type")
	}
}

// setupIDAttribute adds or updates the ID field of a user type definition.
func setupIDAttribute(obj design.Object, res *design.UserTypeDefinition) design.Object {
	idName := ""
	foundID := false
	for n := range obj {
		if n == "ID" || n == "Id" || n == "id" {
			idName = n
			foundID = true
		}
	}

	if foundID {
		// enforce lowercase key
		if idName != "id" {
			obj["id"] = obj[idName]
			delete(obj, idName)
		}
	} else {
		obj["id"] = &design.AttributeDefinition{
			Type:     design.Integer,
			Metadata: design.MetadataDefinition{},
		}
	}

	var gorm string
	if val, ok := metaLookup(res.Metadata, "#gormpktag"); ok {
		gorm = val
	} else {
		gorm = "primary_key"
	}

	// If the user already defined gormtag, leave it alone.
	if _, ok := metaLookup(obj["id"].Metadata, "#gormtag"); !ok {
		obj["id"].Metadata[META_NAMESPACE+"#gormtag"] = gorm
	}

	return obj
}

func MakeTimeFuncs(typeName string) string {
	if timeFields == nil || len(timeFields) == 0 {
		return ""
	}
	var buf bytes.Buffer
	type timefuncArgs struct {
		FieldName, TypeName string
	}

	t, err := template.New("timefuncs").Parse(timefuncTmpl)
	if err != nil {
		panic(err)
	}

	sort.Strings(timeFields)
	for i := range timeFields {
		args := timefuncArgs{
			FieldName: timeFields[i],
			TypeName:  typeName,
		}
		err := t.Execute(&buf, args)
		if err != nil {
			panic(err)
		}
	}

	return buf.String()
}

// isASCIILower returns whether c is an ASCII lower-case letter.
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// isASCIIDigit returns whether c is an ASCII digit.
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// unexport lowercases the first character of a string.
func unexport(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// genfuncs is a map of comments and functions that will be used by ModelDef
// to conditionally add fields to the model struct.  If the function returns
// content, the content will be preceded by the the map key, which should be a
// comment.
var genfuncs = map[string]func(*design.AttributeDefinition) string{
	"\n// Timestamps\n":   includeTimeStamps,
	"\n// Many2Many\n":    includeMany2Many,
	"\n// Foreign Keys\n": includeForeignKey,
	"\n// Children\n":     includeChildren,
	"\n// Authboss\n\n":   includeAuthboss,
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/3d26dc39376c307203d3a221bada26816b3073cf/lint.go#L482
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}
