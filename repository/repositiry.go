package repository

import (
	"API_TRAINING/repository/sql"
	_ "embed"

	"github.com/midir99/sqload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	RepositoryModel struct {
		DB *gorm.DB
		Query *sql.Query
	}
)


func (r *RepositoryModel) SetDB(db *gorm.DB) {
	r.DB = db
}

//go:embed sql/query.sql
var sqlCode string

func (r *RepositoryModel)LoadSql() {
	r.Query = &sql.Query{}
	r.Query = sqload.MustLoadFromString[sql.Query](sqlCode)
}

var Repository = &RepositoryModel{}

func InitDB() error {
	// connStr := fmt.Sprintf(
	// 	"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// )
	dsn := "host=localhost user=postgres password=password dbname=tngp_training_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	Repository.SetDB(db)
	Repository.LoadSql()

    // fmt.Println("Loaded SQL:", Repository.query)
	return nil
}

// func MustNewPostgres() *Postgres {
//     start := time.Now()
//     db, err := sqlx.Connect("postgres", fmt.Sprintf(
//         "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//         cfg.Host,
//         cfg.Port,
//         cfg.User,
//         cfg.Pass,
//         cfg.DB,
//     ))
//     if err != nil {
//         zap.L().Fatal("cannot open postgres connection", zap.Error(err))
//     }
//     zap.L().Info(
//         "connected to postgres",
//         zap.String("label", "platformConnection"),
//         zap.String("platform", "postgres"),
//         zap.String("connectionTime", time.Since(start).String()),
//     )
//     return &Postgres{DB: db}
// }
