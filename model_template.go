package gorma

const modelTmpl = `// {{if .TypeDef.Description}}{{.TypeDef.Description}}{{else}}app.{{ .TypeName}} storage type{{end}}
// Identifier: {{ .TypeName}}
type {{.TypeName}} {{ modeldef .TypeDef }}
{{ $typedef := .TypeDef  }}
{{ $dynamictable := .DoDynamicTableName }}
{{ $typename  := .TypeName }}
{{ $cached := .DoCache }}
{{ $pks := .PrimaryKeys }}
{{ if .DoCustomTableName }}
func (m {{$typename}}) TableName() string {
	return "{{ .CustomTableName}}"
}
{{ end }}
{{ if .DoRoler }}
func (m {{$typename}}) GetRole() string {
	return *m.Role
}
{{end}}

type {{$typename}}Storage interface {
	DB() interface{}
	List(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}) []{{$typename}}
	One(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{ pkattributes $pks  }}) ({{$typename}}, error)
	Add(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, o {{$typename}}) ({{$typename}}, error)
	Update(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, o {{$typename}}) (error)
	Delete(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{ pkattributes $pks }}) (error)
{{ range $idx, $bt := .BelongsTo}}
	ListBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid int) []{{$typename}}
	OneBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid, id int) ({{$typename}}, error)
{{end}}
	{{ storagedef $typedef }}
}
type {{$typename}}DB struct {
	Db gorm.DB
	{{ if .DoCache }}cache *cache.Cache{{end}}
}
{{ range $idx, $bt := .BelongsTo}}
func {{$typename}}FilterBy{{$bt.Parent}}(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("{{ $bt.DatabaseField }}_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func (m *{{$typename}}DB) ListBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid int) []{{$typename}} {

	var objs []{{$typename}}
	m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Scopes({{$typename}}FilterBy{{$bt.Parent}}(parentid, &m.Db)).Find(&objs)
	return objs
}

func (m *{{$typename}}DB) OneBy{{$bt.Parent}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, parentid, {{ pkattributes $pks }}) ({{$typename}}, error) {
	{{ if $cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typename}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typename}}

	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Scopes({{$typename}}FilterBy{{$bt.Parent}}(parentid, &m.Db)).Find(&obj, id).Error
	{{ if $cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}
{{end}}

func New{{$typename}}DB(db gorm.DB) *{{$typename}}DB {
	{{ if $cached }}
	return &{{$typename}}DB{
		Db: db,
		cache: cache.New(5*time.Minute, 30*time.Second),
	}
	{{ else  }}
	return &{{$typename}}DB{Db: db}

	{{ end  }}
}

func (m *{{$typename}}DB) DB() interface{} {
	return &m.Db
}

func (m *{{$typename}}DB) List(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}) []{{$typename}} {

	var objs []{{$typename}}
	m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}


{{ range $idx, $col := columns .TypeDef.AttributeDefinition }}
func (m *{{$typename}}DB) ListBy{{title $col.Column}}Equal(ctx context.Context, {{lower $col.Column}} {{$col.Coltype}}{{ if $dynamictable }}, tableName string{{ end }}) []{{$typename}} {

	var objs []{{$typename}}
	m.Db.Where("{{lower $col.Column}} = ?",  {{lower $col.Column}}){{ if $dynamictable }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}
func (m *{{$typename}}DB) ListBy{{title $col.Column}}Like(ctx context.Context, {{lower $col.Column}} {{$col.Coltype}}{{ if $dynamictable }}, tableName string{{ end }}) []{{$typename}} {

	var objs []{{$typename}}
	m.Db.Where("{{lower $col.Column}} like ?",  {{lower $col.Column}}){{ if $dynamictable }}.Table(tableName){{ end }}.Find(&objs)
	return objs
}
{{ end  }}


func (m *{{$typename}}DB) One(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{pkattributes $pks}}) ({{$typename}}, error) {
	{{ if $cached }}//first attempt to retrieve from cache
	o,found := m.cache.Get(strconv.Itoa(id))
	if found {
		return o.({{$typename}}), nil
	}
	// fallback to database if not found{{ end }}
	var obj {{$typename}}
	{{ $l := len $pks }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Find(&obj, id).Error
	{{ else  }}
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Find(&obj).Where("{{pkwhere $pks}}", {{pkwherefields $pks}}).Error
	{{ end }}
	{{ if $cached }} go m.cache.Set(strconv.Itoa(id), obj, cache.DefaultExpiration) {{ end }}
	return obj, err
}

func (m *{{$typename}}DB) Add(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, model {{$typename}}) ({{$typename}}, error) {
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Create(&model).Error
	{{ if $cached }} go m.cache.Set(strconv.Itoa(model.ID), model, cache.DefaultExpiration) {{ end }}
	return model, err
}

func (m *{{$typename}}DB) Update(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, model {{$typename}}) error {
	obj, err := m.One(ctx{{ if $dynamictable }}, tableName{{ end }}, {{pkupdatefields $pks}})
	if err != nil {
		return  err
	}
	err = m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Model(&obj).Updates(model).Error
	{{ if $cached }}
	go func(){
	obj, err := m.One(ctx, model.ID)
	if err == nil {
		m.cache.Set(strconv.Itoa(model.ID), obj, cache.DefaultExpiration)
	}
	}()
	{{ end }}

	return err
}


func (m *{{$typename}}DB) Delete(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{pkattributes $pks}})  error {
	var obj {{$typename}}
	{{ $l := len $pks }}
	{{ if eq $l 1 }}
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Delete(&obj, id).Error
	{{ else  }}
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Delete(&obj).Where("{{pkwhere $pks}}", {{pkwherefields $pks}}).Error
	{{ end }}
	if err != nil {
		return  err
	}
	{{ if $cached }} go m.cache.Delete(strconv.Itoa(id)) {{ end }}
	return  nil
}

{{ range $idx, $bt := .M2M}}
func (m *{{$typename}}DB) Delete{{$bt.Relation}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{lower $typename}}ID,  {{$bt.LowerRelation}}ID int)  error {
	var obj {{$typename}}
	obj.ID = {{lower $typename}}ID
	var assoc {{$bt.LowerRelation}}.{{$bt.Relation}}
	var err error
	assoc.ID = {{$bt.LowerRelation}}ID
	if err != nil {
		return err
	}
	err = m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.PluralRelation}}").Delete(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typename}}DB) Add{{$bt.Relation}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{lower $typename}}ID, {{$bt.LowerRelation}}ID int) error {
	var {{lower $typename}} {{$typename}}
	{{lower $typename}}.ID = {{lower $typename}}ID
	var assoc {{$bt.LowerRelation}}.{{$bt.Relation}}
	assoc.ID = {{$bt.LowerRelation}}ID
	err := m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Model(&{{lower $typename}}).Association("{{$bt.PluralRelation}}").Append(assoc).Error
	if err != nil {
		return  err
	}
	return  nil
}
func (m *{{$typename}}DB) List{{$bt.PluralRelation}}(ctx context.Context{{ if $dynamictable }}, tableName string{{ end }}, {{lower $typename}}ID int)  []{{$bt.LowerRelation}}.{{$bt.Relation}} {
	var list []{{$bt.LowerRelation}}.{{$bt.Relation}}
	var obj {{$typename}}
	obj.ID = {{lower $typename}}ID
	m.Db{{ if $dynamictable }}.Table(tableName){{ end }}.Model(&obj).Association("{{$bt.PluralRelation}}").Find(&list)
	return  list
}
{{end}}
{{ range $idx, $bt := .BelongsTo}}
func Filter{{$typename}}By{{$bt.Parent}}(parent *int, list []{{$typename}}) []{{$typename}} {
	var filtered []{{$typename}}
	for _,o := range list {
		if o.{{$bt.Parent}}ID == int(*parent) {
			filtered = append(filtered,o)
		}
	}
	return filtered
}
{{end}}
`
