// Code generated by ent, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// FieldSlugName holds the string denoting the slug_name field in the database.
	FieldSlugName = "slug_name"
	// FieldPostCount holds the string denoting the post_count field in the database.
	FieldPostCount = "post_count"
	// Table holds the table name of the tag in the database.
	Table = "tag"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldSlug,
	FieldColor,
	FieldThumbnail,
	FieldSlugName,
	FieldPostCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() int64
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() int64
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)
