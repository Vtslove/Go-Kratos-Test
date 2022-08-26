// Code generated by ent, DO NOT EDIT.

package tag

import (
	"kratos-blog/app/blog/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// SlugName applies equality check predicate on the "slug_name" field. It's identical to SlugNameEQ.
func SlugName(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlugName), v))
	})
}

// PostCount applies equality check predicate on the "post_count" field. It's identical to PostCountEQ.
func PostCount(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostCount), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	})
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSlug), v...))
	})
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	})
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	})
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	})
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	})
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	})
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	})
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	})
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	})
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSlug)))
	})
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSlug)))
	})
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	})
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	})
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldColor), v))
	})
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldColor), v...))
	})
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldColor), v...))
	})
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldColor), v))
	})
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldColor), v))
	})
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldColor), v))
	})
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldColor), v))
	})
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldColor), v))
	})
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldColor), v))
	})
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldColor), v))
	})
}

// ColorIsNil applies the IsNil predicate on the "color" field.
func ColorIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldColor)))
	})
}

// ColorNotNil applies the NotNil predicate on the "color" field.
func ColorNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldColor)))
	})
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldColor), v))
	})
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldColor), v))
	})
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThumbnail)))
	})
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThumbnail)))
	})
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThumbnail), v))
	})
}

// SlugNameEQ applies the EQ predicate on the "slug_name" field.
func SlugNameEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlugName), v))
	})
}

// SlugNameNEQ applies the NEQ predicate on the "slug_name" field.
func SlugNameNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlugName), v))
	})
}

// SlugNameIn applies the In predicate on the "slug_name" field.
func SlugNameIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSlugName), v...))
	})
}

// SlugNameNotIn applies the NotIn predicate on the "slug_name" field.
func SlugNameNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSlugName), v...))
	})
}

// SlugNameGT applies the GT predicate on the "slug_name" field.
func SlugNameGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlugName), v))
	})
}

// SlugNameGTE applies the GTE predicate on the "slug_name" field.
func SlugNameGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlugName), v))
	})
}

// SlugNameLT applies the LT predicate on the "slug_name" field.
func SlugNameLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlugName), v))
	})
}

// SlugNameLTE applies the LTE predicate on the "slug_name" field.
func SlugNameLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlugName), v))
	})
}

// SlugNameContains applies the Contains predicate on the "slug_name" field.
func SlugNameContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlugName), v))
	})
}

// SlugNameHasPrefix applies the HasPrefix predicate on the "slug_name" field.
func SlugNameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlugName), v))
	})
}

// SlugNameHasSuffix applies the HasSuffix predicate on the "slug_name" field.
func SlugNameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlugName), v))
	})
}

// SlugNameIsNil applies the IsNil predicate on the "slug_name" field.
func SlugNameIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSlugName)))
	})
}

// SlugNameNotNil applies the NotNil predicate on the "slug_name" field.
func SlugNameNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSlugName)))
	})
}

// SlugNameEqualFold applies the EqualFold predicate on the "slug_name" field.
func SlugNameEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlugName), v))
	})
}

// SlugNameContainsFold applies the ContainsFold predicate on the "slug_name" field.
func SlugNameContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlugName), v))
	})
}

// PostCountEQ applies the EQ predicate on the "post_count" field.
func PostCountEQ(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostCount), v))
	})
}

// PostCountNEQ applies the NEQ predicate on the "post_count" field.
func PostCountNEQ(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPostCount), v))
	})
}

// PostCountIn applies the In predicate on the "post_count" field.
func PostCountIn(vs ...uint32) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPostCount), v...))
	})
}

// PostCountNotIn applies the NotIn predicate on the "post_count" field.
func PostCountNotIn(vs ...uint32) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPostCount), v...))
	})
}

// PostCountGT applies the GT predicate on the "post_count" field.
func PostCountGT(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPostCount), v))
	})
}

// PostCountGTE applies the GTE predicate on the "post_count" field.
func PostCountGTE(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPostCount), v))
	})
}

// PostCountLT applies the LT predicate on the "post_count" field.
func PostCountLT(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPostCount), v))
	})
}

// PostCountLTE applies the LTE predicate on the "post_count" field.
func PostCountLTE(v uint32) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPostCount), v))
	})
}

// PostCountIsNil applies the IsNil predicate on the "post_count" field.
func PostCountIsNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPostCount)))
	})
}

// PostCountNotNil applies the NotNil predicate on the "post_count" field.
func PostCountNotNil() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPostCount)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		p(s.Not())
	})
}
