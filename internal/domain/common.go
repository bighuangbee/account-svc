package domain

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Id64Info struct {
	Id *int64 `gorm:"type:bigint(11) not null auto_increment;primaryKey"`
}

type CreatedAtInfo struct {
	CreatedAt time.Time `json:"createdAt,omitempty" dbupdate:"created_at" gorm:"type:timestamp;not null;comment:创建时间"`
	CreatedBy uint64    `json:"createdById,omitempty" dbupdate:"created_by" gorm:"type:bigint(11) unsigned;not null;default:0;comment:创建者"`
}
type CreatedByNameInfo struct {
	CreatedByName string `json:"createdBy,omitempty" dbupdate:"created_by_name" gorm:"type:varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;comment:创建者姓名"`
}
type UpdatedAtInfo struct {
	UpdatedAt sql.NullTime `json:"updatedAt,omitempty" dbupdate:"updated_at" gorm:"type:timestamp;default:null;comment:更新时间"`
	UpdatedBy uint64       `json:"updatedById,omitempty" dbupdate:"updated_by" gorm:"comment:更新者"`
}
type UpdatedByNameInfo struct {
	UpdatedByName string `json:"updatedBy,omitempty" dbupdate:"updated_by_name" gorm:"comment:更新者姓名"`
}
type DeletedAtInfo struct {
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" dbupdate:"deleted_at" gorm:"type:timestamp;comment:删除时间"` //gorm 自动启动软删除
	DeletedBy uint64         `json:"deletedBy,omitempty" dbupdate:"deleted_by" gorm:"comment:删除者"`
}
type DeletedByNameInfo struct {
	DeletedByName string `json:"deletedByName,omitempty" dbupdate:"deleted_by_name" gorm:"comment:删除者姓名"`
}
