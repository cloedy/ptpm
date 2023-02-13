package data

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "ptpm/internal/conf"

    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
    // TODO wrapped database client
    db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }

    db, err := NewDb(c)
    if err != nil {
        panic("database connect error")
    }

    dbAfterInit(db)

    return &Data{
        db: db,
    }, cleanup, nil
}

func NewDb(c *conf.Data) (*gorm.DB, error) {
    mdsn := c.Database.Mysqldsn
    db, err := gorm.Open(mysql.Open(mdsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    return db, err
}

func dbAfterInit(db *gorm.DB) {
    db.AutoMigrate(&Task{})
}
