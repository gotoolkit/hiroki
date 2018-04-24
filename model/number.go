package model

type Number struct {
    ID         uint   `gorm:"primary_key"`
    Reds       string `gorm:"type:char(17);not null"`
    Beginning  string `gorm:"type:char(2);not null"`
    Ending     string `gorm:"type:char(2);not null"`
    AC         uint8  `gorm:"type:tinyint;not null"`
    QuJian     string `gorm:"type:char(3);not null"`
    JiOu       string `gorm:"type:char(2);not null"`
    DaXiao     string `gorm:"type:char(2);not null"`
    Sum        uint   `gorm:"type:SMALLINT;not null"`
    TongWei    uint8  `gorm:"type:tinyint;not null"`
    LianHao    uint8  `gorm:"type:tinyint;not null"`
    MaxLianHao uint8  `gorm:"type:tinyint;not null"`
    KuaJu      uint8  `gorm:"type:tinyint;not null"`
}

func (Number) TableName() string {
    return "number"
}
