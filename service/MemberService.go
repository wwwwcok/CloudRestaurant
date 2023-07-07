package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/wonderivan/logger"
)

type MemberService struct {
}

func (ms *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {
	//获取手机号和验证码
	//验证手机号和验证码是否正确
	md := dao.MemberDao{Orm: tool.DbEngine}
	md.ValidateSmsCode(loginParam.Phone, loginParam.Code)
	//根据手机号查用户的数据
	md.QueryByPhone(loginParam.Phone)

	//新建member记录并保存
	user := model.Member{}
	user.Mobile = loginParam.Phone
	user.UserName = loginParam.Phone
	user.RegisterTime = time.Now().Unix()

	//返回的是插入数据的行数
	user.Id = md.InsertMember(user)

	return &user
}

func (ms *MemberService) Sendcode(phone string) bool {

	//产生验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	//调用阿里云sdk
	config := tool.GetConfig().Sms

	client, _ := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone

	par, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})

	request.TemplateParam = string(par)
	response, err := client.SendSms(request)

	fmt.Println("这是sms响应", response)

	if err != nil {
		logger.Error(err.Error())
		return false
	}

	//返回结果
	if response.Code == "OK" {
		//将验证码存储到数据库
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := dao.MemberDao{Orm: tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}

	return false
}

//用户登录
func (ms *MemberService) Login(name string, password string) *model.Member {
	//根据用户信息查询，存在直接返回
	md := dao.MemberDao{Orm: tool.DbEngine}
	member := md.Query(name, password)
	if member.UserName != "" {
		return member
	}

	//不存在新建
	user := model.Member{}
	user.Mobile = name
	user.Password = base64.StdEncoding.EncodeToString(sha256.New().Sum([]byte(password)))
	user.RegisterTime = time.Now().Unix()

	md.InsertMember(user)

	return &user
}

func (ms *MemberService) UploadAvator(userId int64, fileName string) string {
	MemberDao := dao.MemberDao{Orm: tool.DbEngine}
	result := MemberDao.UpdateMemberAvatar(userId, fileName)
	if result == 0 {
		return ""
	}

	return fileName
}

func (ms *MemberService) GetUserInfo(UserId string) *model.Member {
	id, _ := strconv.Atoi(UserId)
	dao := dao.MemberDao{Orm: tool.DbEngine}
	member := dao.GetUserById(id)
	return member
}
