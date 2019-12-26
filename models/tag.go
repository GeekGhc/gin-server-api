package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	Status     int    `json:"status"`
}

//校验标签名是否重复
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? and delete_at = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

//创建标签
func AddTag(name string, status int, createdBy string) error {
	tag := Tag{
		Name:      name,
		Status:    status,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// 获取标签列表
func GetTags(pageNo int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageNo > 0 && pageSize > 0 {
		err = db.Where(maps).Find(&tags).Offset((pageNo - 1) * pageSize).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}
