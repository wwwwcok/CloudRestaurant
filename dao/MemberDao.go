package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/wonderivan/logger"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) GetUserById(userId int) *model.Member {
	var member model.Member
	_, err := md.Where("id = ?", userId).Get(&member)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &member
}

func (md *MemberDao) UpdateMemberAvatar(userId int64, filename string) int64 {
	member := model.Member{Avatar: filename}
	result, err := md.Where("id = ? ", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result

}

func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member

	_, err := md.Where("name = ?and password = ?", name, base64.StdEncoding.EncodeToString(sha256.New().Sum([]byte(password)))).Get(&member)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &member
}

func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	_, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms)

	if err != nil {
		logger.Error(err.Error())
	}
	return &sms

}
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	_, err := md.Where("mobile = ?", phone).Get(&member)
	if err != nil {
		logger.Error(err.Error())
	}
	return &member
}

//新用户插入
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
