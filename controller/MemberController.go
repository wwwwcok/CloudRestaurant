package controller

import (
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (r *MemberController) Router(Engine *gin.Engine) {
	//Engine.GET("/api/sendcode", r.sendSmsCode)
	Engine.OPTIONS("/api/login_sms", r.smsLogin)
	//Engine.POST("/api/login_sms", r.smsLogin)
	Engine.GET("/api/captcha", r.captcha)
	Engine.POST("/api/vertifycha", r.vertifyCaptcha)
	Engine.POST("/api/login_pwd", r.nameLogin)
	Engine.POST("/api/upload/avator", r.uploadAvator)
	Engine.GET("/api/userinfo", r.userinfo)

}

func (c *MemberController) userinfo(context *gin.Context) {
	cookie, err := tool.CookieAuth(context)
	if err != nil {
		//context.Abort()
		tool.Failed(context, "cookie失效")
		return
	}
	var MemberService service.MemberService
	member := MemberService.GetUserInfo(cookie.Value)
	if member == nil {
		tool.Failed(context, "没用查询到消息")
		return
	}

	tool.Success(context, member)
}

func (c *MemberController) uploadAvator(context *gin.Context) {
	//从上下文中解析上传的数据 file ,user_id
	userId := context.PostForm("user_id")
	fmt.Println(userId)
	file, err := context.FormFile("avatar")
	if err != nil {
		tool.Failed(context, "图像参数解析失败")
		return
	}

	//判断user_id是否登录
	sess := tool.Getsess(context, "user_id"+userId)
	if sess == nil {
		tool.Failed(context, "没有找到对应session,参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)
	//将文件保存到本地

	fileName := "./uploadfile" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "图像更新失败")
	}

	//将文件上传到fastDFS系统
	fileId := tool.UploadFile(fileName)
	if fileId != "" {
		os.Remove(fileName)
	}

	//将文件保存到数据库中
	memberSevice := service.MemberService{}
	path := memberSevice.UploadAvator(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://192.168.0.90:80"+"/"+path)
		return
	}

}

func (c *MemberController) nameLogin(context *gin.Context) {
	//解析参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确")
		return
	}
	//登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		sess, _ := json.Marshal(member)

		//给登录的用户设置session
		err := tool.Setsess(context, "user_"+strconv.Itoa(int(member.Id)), sess)
		if err != nil {
			tool.Failed(context, "登录失败:设置session失败")
		}
		tool.Success(context, member)
		return

	}
}

// http://localhost:8090/hello/sendcode?
func (c *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")

	if !exist {
		tool.Failed(context, "参数解析失败")
		return
	}
	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)

	if isSend {
		tool.Success(context, "参数解析成功")
		return
	}
	tool.Failed(context, "发送失败")
}

func (c *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
	}

	//手机验证登录的控制器

	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if err == nil {

		//登录成功时设置cookie
		context.SetCookie("cookie_user", string(rune(member.Id)), 10*60, "/", "localhost", true, true)

		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登录失败")

}

func (r *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

func (r *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败")
	}

	result := tool.VertifyCaptcha(captcha.Id, captcha.Vertifyvalue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
	
}
