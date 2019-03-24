package user

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/humyPro/golang/logindemo/database"
	"github.com/humyPro/golang/logindemo/model"
	"github.com/humyPro/golang/logindemo/util"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var strTest, _ = regexp.Compile(`.?\s.?`)

var result model.Result
var db = database.DBInstance()
var redisCli = *database.RedisCli()

const TimeOut = int64(time.Minute * 60 * 24 * 7) //过期时间一周

type UserController struct {
}

func (u *UserController) Register(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		bytes, _ := json.Marshal(err)
		c.String(http.StatusOK, string(bytes))
		return
	}

	inStr := []string{user.Nickname, user.Username, user.Tel, user.Password}

	for _, str := range inStr {
		if strTest.MatchString(str) || str == "" {
			bytes, _ := json.Marshal(result.Err("请重新新完善注册信息"))
			c.String(http.StatusOK, string(bytes))
			return
		}
	}
	var temps []model.User
	db.Where("tel=?", user.Tel).Find(&temps)
	if len(temps) != 0 {
		bytes, _ := json.Marshal(result.Err("手机号码已被注册"))
		c.String(http.StatusOK, string(bytes))
		return
	}

	//密码md5加密
	user.Password = *util.Md5(&user.Password)
	db.Create(&user)
	bytes, _ := json.Marshal(result.Suc("注册成功"))
	c.String(http.StatusOK, string(bytes))
}

func (u *UserController) Login(c *gin.Context) {
	tel := c.PostForm("tel")
	pw := c.PostForm("password")
	if tel == "" || pw == "" {
		bytes, _ := json.Marshal(result.Err("帐号或密码错误"))
		c.String(http.StatusOK, string(bytes))
		return
	}
	var key = "userToken_" + tel
	//reply, _ := redisCli.Do("GET", key)
	var user model.User

	// redis中没有，重新登陆验证并写入redis
	//if reply == nil {
	db.Where("tel = ?", tel).First(&user)
	if user.Tel == "" {
		bytes, _ := json.Marshal(result.Err("用户不存在，请注册"))
		c.String(http.StatusOK, string(bytes))
		return
	}
	if *util.Md5(&pw) != user.Password {
		bytes, _ := json.Marshal(result.Err("帐号或密码错误"))
		c.String(http.StatusOK, string(bytes))
		return
	}

	str := strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(user.ID)) + user.Tel + user.Password
	token := *util.Md5(&str)

	reply2, _ := redisCli.Do("SET", key, token, "EX", TimeOut)
	if reply2.(string) == "OK" {
		//resData := make(map[string]string)
		//resData["token"]=token
		//resData["te"]=user.Nickname
		res := result.Get(1, "登录成功", token)
		bytes, _ := json.Marshal(res)
		c.String(http.StatusOK, string(bytes))
		return
	}
	bytes, _ := json.Marshal(result.Err("登录失败，请稍后重试"))
	c.String(http.StatusOK, string(bytes))
}

func (u *UserController) Authorize(c *gin.Context) {
	token := c.Query("token")
	tel := c.Query("tel")
	if token == "" || tel == "" {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	reply, _ := redis.String(redisCli.Do("GET", "userToken_"+tel))
	if reply == "" || reply != token {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.Next()
}

func (u *UserController) Logout(c *gin.Context) {
	//token := c.Query("token")
	tel := c.Query("tel")
	var key = "userToken_" + tel
	redis.Int(redisCli.Do("DEL", key))
	bytes, _ := json.Marshal(result.Suc("退出成功"))
	c.String(http.StatusOK, string(bytes))
}
