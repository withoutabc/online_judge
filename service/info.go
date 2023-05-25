package service

import (
	"gorm.io/gorm"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
)

type InfoDao interface {
	AddInfo(tx *gorm.DB, info *model.Info) error
	GetInfo(uid int64) (model.Info, error)
	UpdateInfo(info model.Info) (int64, error)
}

func NewInfoServiceImpl() *InfoDaoImpl {
	return &InfoDaoImpl{
		InfoDao: dao.NewInfoDao(),
		UserDao: dao.NewUserDao(),
	}
}

type InfoDaoImpl struct {
	InfoDao
	UserDao
}

func (i *InfoDaoImpl) GetInfo(uid int64) (model.Info, int) {
	info, err := i.InfoDao.GetInfo(uid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Info{}, util.NoRecordErrCode
		}
		return model.Info{}, util.InternalServeErrCode
	}
	user, err := i.UserDao.SearchUserById(uid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Info{}, util.NoRecordErrCode
		}
		return model.Info{}, util.InternalServeErrCode
	}
	info.UserId = user.UserId
	return info, util.NoErrCode
}

func (i *InfoDaoImpl) UpdateInfo(info model.Info) int {
	count, err := i.InfoDao.UpdateInfo(info)
	if err != nil {
		return util.InternalServeErrCode
	}
	if count != 1 {
		return util.UpdateFailErrCode
	}
	return util.NoErrCode
}
