package dao

import (
	"gorm.io/gorm"
	"online_judge/model"
)

func NewInfoDao() *InfoDaoImpl {
	return &InfoDaoImpl{
		db: DB,
	}
}

type InfoDaoImpl struct {
	db *gorm.DB
}

func (i *InfoDaoImpl) AddCorrect(userId int64) error {
	result := i.db.Model(&model.Info{}).Where(&model.Info{UserId: userId}).UpdateColumn("correct", gorm.Expr("correct + ?", 1))
	return result.Error
}

func (i *InfoDaoImpl) AddScore(userId int64, score int) error {
	result := i.db.Model(&model.Info{}).Where(&model.Info{UserId: userId}).UpdateColumn("score", gorm.Expr("score + ?", score))
	return result.Error
}

func (i *InfoDaoImpl) AddInfo(tx *gorm.DB, info *model.Info) error {
	result := tx.Create(info)
	return result.Error
}

func (i *InfoDaoImpl) GetInfo(uid int64) (model.Info, error) {
	var info model.Info
	result := i.db.Where(&model.Info{UserId: uid}).First(&info)
	return info, result.Error
}

func (i *InfoDaoImpl) UpdateInfo(info model.Info) (int64, error) {
	result := i.db.Model(&model.Info{}).Where(&model.Info{UserId: info.UserId}).Updates(model.Info{
		Name:     info.Name,
		UserId:   info.UserId,
		Nickname: info.Nickname,
		Gender:   info.Gender,
		Year:     info.Year,
		Month:    info.Month,
		Day:      info.Day,
		Email:    info.Email,
	})
	return result.RowsAffected, result.Error
}
