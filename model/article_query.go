// Code generated by go-queryset. DO NOT EDIT.
package model

import (
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ArticleQuerySet

// ArticleQuerySet is an queryset type for Article
type ArticleQuerySet struct {
	db *gorm.DB
}

// NewArticleQuerySet constructs new ArticleQuerySet
func NewArticleQuerySet(db *gorm.DB) ArticleQuerySet {
	return ArticleQuerySet{
		db: db.Model(&Article{}),
	}
}

func (qs ArticleQuerySet) w(db *gorm.DB) ArticleQuerySet {
	return NewArticleQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *Article) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Article) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// AddDateEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateEq(addDate string) ArticleQuerySet {
	return qs.w(qs.db.Where("add_date = ?", addDate))
}

// AddDateIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateIn(addDate ...string) ArticleQuerySet {
	if len(addDate) == 0 {
		qs.db.AddError(errors.New("must at least pass one addDate in AddDateIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("add_date IN (?)", addDate))
}

// AddDateLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateLike(addDate string) ArticleQuerySet {
	return qs.w(qs.db.Where("add_date LIKE ?", addDate))
}

// AddDateNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateNe(addDate string) ArticleQuerySet {
	return qs.w(qs.db.Where("add_date != ?", addDate))
}

// AddDateNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateNotIn(addDate ...string) ArticleQuerySet {
	if len(addDate) == 0 {
		qs.db.AddError(errors.New("must at least pass one addDate in AddDateNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("add_date NOT IN (?)", addDate))
}

// AddDateNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) AddDateNotlike(addDate string) ArticleQuerySet {
	return qs.w(qs.db.Where("add_date NOT LIKE ?", addDate))
}

// All is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) All(ret *[]Article) error {
	return qs.db.Find(ret).Error
}

// BodyEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyEq(body template.HTML) ArticleQuerySet {
	return qs.w(qs.db.Where("body = ?", body))
}

// BodyIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyIn(body ...template.HTML) ArticleQuerySet {
	if len(body) == 0 {
		qs.db.AddError(errors.New("must at least pass one body in BodyIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("body IN (?)", body))
}

// BodyLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyLike(body template.HTML) ArticleQuerySet {
	return qs.w(qs.db.Where("body LIKE ?", body))
}

// BodyNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyNe(body template.HTML) ArticleQuerySet {
	return qs.w(qs.db.Where("body != ?", body))
}

// BodyNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyNotIn(body ...template.HTML) ArticleQuerySet {
	if len(body) == 0 {
		qs.db.AddError(errors.New("must at least pass one body in BodyNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("body NOT IN (?)", body))
}

// BodyNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) BodyNotlike(body template.HTML) ArticleQuerySet {
	return qs.w(qs.db.Where("body NOT LIKE ?", body))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtEq(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtGt(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtGte(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtLt(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtLte(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) CreatedAtNe(createdAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Delete() error {
	return qs.db.Delete(Article{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Article{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Article{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtEq(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtGt(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtGte(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtIsNotNull() ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtIsNull() ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtLt(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtLte(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DeletedAtNe(deletedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionEq(description string) ArticleQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionIn(description ...string) ArticleQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description IN (?)", description))
}

// DescriptionLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionLike(description string) ArticleQuerySet {
	return qs.w(qs.db.Where("description LIKE ?", description))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionNe(description string) ArticleQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionNotIn(description ...string) ArticleQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", description))
}

// DescriptionNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) DescriptionNotlike(description string) ArticleQuerySet {
	return qs.w(qs.db.Where("description NOT LIKE ?", description))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) GetUpdater() ArticleUpdater {
	return NewArticleUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDEq(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDGt(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDGte(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDIn(ID ...uint) ArticleQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDLt(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDLte(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDNe(ID uint) ArticleQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) IDNotIn(ID ...uint) ArticleQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Limit(limit int) ArticleQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NickNameEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameEq(nickName string) ArticleQuerySet {
	return qs.w(qs.db.Where("nick_name = ?", nickName))
}

// NickNameIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameIn(nickName ...string) ArticleQuerySet {
	if len(nickName) == 0 {
		qs.db.AddError(errors.New("must at least pass one nickName in NickNameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("nick_name IN (?)", nickName))
}

// NickNameLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameLike(nickName string) ArticleQuerySet {
	return qs.w(qs.db.Where("nick_name LIKE ?", nickName))
}

// NickNameNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameNe(nickName string) ArticleQuerySet {
	return qs.w(qs.db.Where("nick_name != ?", nickName))
}

// NickNameNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameNotIn(nickName ...string) ArticleQuerySet {
	if len(nickName) == 0 {
		qs.db.AddError(errors.New("must at least pass one nickName in NickNameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("nick_name NOT IN (?)", nickName))
}

// NickNameNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) NickNameNotlike(nickName string) ArticleQuerySet {
	return qs.w(qs.db.Where("nick_name NOT LIKE ?", nickName))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) Offset(offset int) ArticleQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ArticleQuerySet) One(ret *Article) error {
	return qs.db.First(ret).Error
}

// OrderAscByAddDate is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByAddDate() ArticleQuerySet {
	return qs.w(qs.db.Order("add_date ASC"))
}

// OrderAscByBody is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByBody() ArticleQuerySet {
	return qs.w(qs.db.Order("body ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByCreatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByDeletedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByDescription is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByDescription() ArticleQuerySet {
	return qs.w(qs.db.Order("description ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByID() ArticleQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByNickName is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByNickName() ArticleQuerySet {
	return qs.w(qs.db.Order("nick_name ASC"))
}

// OrderAscByOrigin is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByOrigin() ArticleQuerySet {
	return qs.w(qs.db.Order("origin ASC"))
}

// OrderAscByPicture is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByPicture() ArticleQuerySet {
	return qs.w(qs.db.Order("picture ASC"))
}

// OrderAscByPostType is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByPostType() ArticleQuerySet {
	return qs.w(qs.db.Order("post_type ASC"))
}

// OrderAscByStatus is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByStatus() ArticleQuerySet {
	return qs.w(qs.db.Order("status ASC"))
}

// OrderAscBySubject is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscBySubject() ArticleQuerySet {
	return qs.w(qs.db.Order("subject ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByUpdatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByUserID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByUserID() ArticleQuerySet {
	return qs.w(qs.db.Order("user_id ASC"))
}

// OrderAscByVisited is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderAscByVisited() ArticleQuerySet {
	return qs.w(qs.db.Order("visited ASC"))
}

// OrderDescByAddDate is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByAddDate() ArticleQuerySet {
	return qs.w(qs.db.Order("add_date DESC"))
}

// OrderDescByBody is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByBody() ArticleQuerySet {
	return qs.w(qs.db.Order("body DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByCreatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByDeletedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByDescription is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByDescription() ArticleQuerySet {
	return qs.w(qs.db.Order("description DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByID() ArticleQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByNickName is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByNickName() ArticleQuerySet {
	return qs.w(qs.db.Order("nick_name DESC"))
}

// OrderDescByOrigin is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByOrigin() ArticleQuerySet {
	return qs.w(qs.db.Order("origin DESC"))
}

// OrderDescByPicture is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByPicture() ArticleQuerySet {
	return qs.w(qs.db.Order("picture DESC"))
}

// OrderDescByPostType is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByPostType() ArticleQuerySet {
	return qs.w(qs.db.Order("post_type DESC"))
}

// OrderDescByStatus is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByStatus() ArticleQuerySet {
	return qs.w(qs.db.Order("status DESC"))
}

// OrderDescBySubject is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescBySubject() ArticleQuerySet {
	return qs.w(qs.db.Order("subject DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByUpdatedAt() ArticleQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByUserID is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByUserID() ArticleQuerySet {
	return qs.w(qs.db.Order("user_id DESC"))
}

// OrderDescByVisited is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OrderDescByVisited() ArticleQuerySet {
	return qs.w(qs.db.Order("visited DESC"))
}

// OriginEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginEq(origin string) ArticleQuerySet {
	return qs.w(qs.db.Where("origin = ?", origin))
}

// OriginIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginIn(origin ...string) ArticleQuerySet {
	if len(origin) == 0 {
		qs.db.AddError(errors.New("must at least pass one origin in OriginIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("origin IN (?)", origin))
}

// OriginLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginLike(origin string) ArticleQuerySet {
	return qs.w(qs.db.Where("origin LIKE ?", origin))
}

// OriginNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginNe(origin string) ArticleQuerySet {
	return qs.w(qs.db.Where("origin != ?", origin))
}

// OriginNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginNotIn(origin ...string) ArticleQuerySet {
	if len(origin) == 0 {
		qs.db.AddError(errors.New("must at least pass one origin in OriginNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("origin NOT IN (?)", origin))
}

// OriginNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) OriginNotlike(origin string) ArticleQuerySet {
	return qs.w(qs.db.Where("origin NOT LIKE ?", origin))
}

// PictureEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureEq(picture string) ArticleQuerySet {
	return qs.w(qs.db.Where("picture = ?", picture))
}

// PictureIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureIn(picture ...string) ArticleQuerySet {
	if len(picture) == 0 {
		qs.db.AddError(errors.New("must at least pass one picture in PictureIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("picture IN (?)", picture))
}

// PictureLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureLike(picture string) ArticleQuerySet {
	return qs.w(qs.db.Where("picture LIKE ?", picture))
}

// PictureNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureNe(picture string) ArticleQuerySet {
	return qs.w(qs.db.Where("picture != ?", picture))
}

// PictureNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureNotIn(picture ...string) ArticleQuerySet {
	if len(picture) == 0 {
		qs.db.AddError(errors.New("must at least pass one picture in PictureNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("picture NOT IN (?)", picture))
}

// PictureNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PictureNotlike(picture string) ArticleQuerySet {
	return qs.w(qs.db.Where("picture NOT LIKE ?", picture))
}

// PostTypeEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeEq(postType string) ArticleQuerySet {
	return qs.w(qs.db.Where("post_type = ?", postType))
}

// PostTypeIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeIn(postType ...string) ArticleQuerySet {
	if len(postType) == 0 {
		qs.db.AddError(errors.New("must at least pass one postType in PostTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("post_type IN (?)", postType))
}

// PostTypeLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeLike(postType string) ArticleQuerySet {
	return qs.w(qs.db.Where("post_type LIKE ?", postType))
}

// PostTypeNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeNe(postType string) ArticleQuerySet {
	return qs.w(qs.db.Where("post_type != ?", postType))
}

// PostTypeNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeNotIn(postType ...string) ArticleQuerySet {
	if len(postType) == 0 {
		qs.db.AddError(errors.New("must at least pass one postType in PostTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("post_type NOT IN (?)", postType))
}

// PostTypeNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) PostTypeNotlike(postType string) ArticleQuerySet {
	return qs.w(qs.db.Where("post_type NOT LIKE ?", postType))
}

// StatusEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusEq(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status = ?", status))
}

// StatusGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusGt(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status > ?", status))
}

// StatusGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusGte(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status >= ?", status))
}

// StatusIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusIn(status ...int) ArticleQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status IN (?)", status))
}

// StatusLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusLt(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status < ?", status))
}

// StatusLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusLte(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status <= ?", status))
}

// StatusNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusNe(status int) ArticleQuerySet {
	return qs.w(qs.db.Where("status != ?", status))
}

// StatusNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) StatusNotIn(status ...int) ArticleQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status NOT IN (?)", status))
}

// SubjectEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectEq(subject string) ArticleQuerySet {
	return qs.w(qs.db.Where("subject = ?", subject))
}

// SubjectIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectIn(subject ...string) ArticleQuerySet {
	if len(subject) == 0 {
		qs.db.AddError(errors.New("must at least pass one subject in SubjectIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("subject IN (?)", subject))
}

// SubjectLike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectLike(subject string) ArticleQuerySet {
	return qs.w(qs.db.Where("subject LIKE ?", subject))
}

// SubjectNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectNe(subject string) ArticleQuerySet {
	return qs.w(qs.db.Where("subject != ?", subject))
}

// SubjectNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectNotIn(subject ...string) ArticleQuerySet {
	if len(subject) == 0 {
		qs.db.AddError(errors.New("must at least pass one subject in SubjectNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("subject NOT IN (?)", subject))
}

// SubjectNotlike is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) SubjectNotlike(subject string) ArticleQuerySet {
	return qs.w(qs.db.Where("subject NOT LIKE ?", subject))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtEq(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtGt(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtGte(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtLt(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtLte(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UpdatedAtNe(updatedAt time.Time) ArticleQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UserIDEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDEq(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id = ?", userID))
}

// UserIDGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDGt(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id > ?", userID))
}

// UserIDGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDGte(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id >= ?", userID))
}

// UserIDIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDIn(userID ...int) ArticleQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id IN (?)", userID))
}

// UserIDLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDLt(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id < ?", userID))
}

// UserIDLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDLte(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id <= ?", userID))
}

// UserIDNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDNe(userID int) ArticleQuerySet {
	return qs.w(qs.db.Where("user_id != ?", userID))
}

// UserIDNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) UserIDNotIn(userID ...int) ArticleQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id NOT IN (?)", userID))
}

// VisitedEq is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedEq(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited = ?", visited))
}

// VisitedGt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedGt(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited > ?", visited))
}

// VisitedGte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedGte(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited >= ?", visited))
}

// VisitedIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedIn(visited ...int) ArticleQuerySet {
	if len(visited) == 0 {
		qs.db.AddError(errors.New("must at least pass one visited in VisitedIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("visited IN (?)", visited))
}

// VisitedLt is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedLt(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited < ?", visited))
}

// VisitedLte is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedLte(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited <= ?", visited))
}

// VisitedNe is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedNe(visited int) ArticleQuerySet {
	return qs.w(qs.db.Where("visited != ?", visited))
}

// VisitedNotIn is an autogenerated method
// nolint: dupl
func (qs ArticleQuerySet) VisitedNotIn(visited ...int) ArticleQuerySet {
	if len(visited) == 0 {
		qs.db.AddError(errors.New("must at least pass one visited in VisitedNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("visited NOT IN (?)", visited))
}

// SetAddDate is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetAddDate(addDate string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.AddDate)] = addDate
	return u
}

// SetBody is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetBody(body template.HTML) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Body)] = body
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetCreatedAt(createdAt time.Time) ArticleUpdater {
	u.fields[string(ArticleDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetDeletedAt(deletedAt *time.Time) ArticleUpdater {
	u.fields[string(ArticleDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetDescription(description string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetID(ID uint) ArticleUpdater {
	u.fields[string(ArticleDBSchema.ID)] = ID
	return u
}

// SetNickName is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetNickName(nickName string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.NickName)] = nickName
	return u
}

// SetOrigin is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetOrigin(origin string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Origin)] = origin
	return u
}

// SetPicture is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetPicture(picture string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Picture)] = picture
	return u
}

// SetPostType is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetPostType(postType string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.PostType)] = postType
	return u
}

// SetStatus is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetStatus(status int) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Status)] = status
	return u
}

// SetSubject is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetSubject(subject string) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Subject)] = subject
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetUpdatedAt(updatedAt time.Time) ArticleUpdater {
	u.fields[string(ArticleDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUserID is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetUserID(userID int) ArticleUpdater {
	u.fields[string(ArticleDBSchema.UserID)] = userID
	return u
}

// SetVisited is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) SetVisited(visited int) ArticleUpdater {
	u.fields[string(ArticleDBSchema.Visited)] = visited
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ArticleUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ArticleQuerySet

// ===== BEGIN of Article modifiers

// ArticleDBSchemaField describes database schema field. It requires for method 'Update'
type ArticleDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ArticleDBSchemaField) String() string {
	return string(f)
}

// ArticleDBSchema stores db field names of Article
var ArticleDBSchema = struct {
	ID          ArticleDBSchemaField
	CreatedAt   ArticleDBSchemaField
	UpdatedAt   ArticleDBSchemaField
	DeletedAt   ArticleDBSchemaField
	PostType    ArticleDBSchemaField
	Subject     ArticleDBSchemaField
	Picture     ArticleDBSchemaField
	Description ArticleDBSchemaField
	Body        ArticleDBSchemaField
	UserID      ArticleDBSchemaField
	NickName    ArticleDBSchemaField
	Origin      ArticleDBSchemaField
	AddDate     ArticleDBSchemaField
	Visited     ArticleDBSchemaField
	Status      ArticleDBSchemaField
}{

	ID:          ArticleDBSchemaField("id"),
	CreatedAt:   ArticleDBSchemaField("created_at"),
	UpdatedAt:   ArticleDBSchemaField("updated_at"),
	DeletedAt:   ArticleDBSchemaField("deleted_at"),
	PostType:    ArticleDBSchemaField("post_type"),
	Subject:     ArticleDBSchemaField("subject"),
	Picture:     ArticleDBSchemaField("picture"),
	Description: ArticleDBSchemaField("description"),
	Body:        ArticleDBSchemaField("body"),
	UserID:      ArticleDBSchemaField("user_id"),
	NickName:    ArticleDBSchemaField("nick_name"),
	Origin:      ArticleDBSchemaField("origin"),
	AddDate:     ArticleDBSchemaField("add_date"),
	Visited:     ArticleDBSchemaField("visited"),
	Status:      ArticleDBSchemaField("status"),
}

// Update updates Article fields by primary key
// nolint: dupl
func (o *Article) Update(db *gorm.DB, fields ...ArticleDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"deleted_at":  o.DeletedAt,
		"post_type":   o.PostType,
		"subject":     o.Subject,
		"picture":     o.Picture,
		"description": o.Description,
		"body":        o.Body,
		"user_id":     o.UserID,
		"nick_name":   o.NickName,
		"origin":      o.Origin,
		"add_date":    o.AddDate,
		"visited":     o.Visited,
		"status":      o.Status,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Article %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ArticleUpdater is an Article updates manager
type ArticleUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewArticleUpdater creates new Article updater
// nolint: dupl
func NewArticleUpdater(db *gorm.DB) ArticleUpdater {
	return ArticleUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Article{}),
	}
}

// ===== END of Article modifiers

// ===== END of all query sets
