package data

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    TaskId    string `gorm:"size:32;comment:任务ID"  json:"task_id"`
    UserId    string `gorm:"size:32;comment:用户ID"  json:"user_id"`
    Img       string `gorm:"size:200;comment:封面图"  json:"img"`
    Tags      string `gorm:"size:255;comment:多个标签"  json:"tags"`
    ShopPrice uint32 `gorm:"comment:市场价，标价" json:"shop_price"`
    Score     uint8  `gorm:"comment:得分" json:"score"`
}
