package service

import (
	"auroralab/models"
	"auroralab/repository"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type UserService struct {
	userDao *repository.UserDao
}

func (s *UserService) Apply(c *gin.Context) (models.BaseResponse, error) {
	resp := models.BaseResponse{
		StatusCode: 500,
		StatusMsg:  "fail",
	}

	name := c.PostForm("name") // 获取 name 参数
	if name == "" {
		resp.StatusMsg = "姓名不能为空"
		return resp, errors.New("参数不能为空值")
		// 如果 name 参数为空，则执行相关处理逻辑
	}

	email := c.PostForm("email") // 获取 email 参数
	if email == "" {
		resp.StatusMsg = "邮箱不能为空"
		return resp, errors.New("参数不能为空值")
		// 如果 email 参数为空，则执行相关处理逻辑
	}

	major := c.PostForm("major") // 获取 major 参数
	if major == "" {
		resp.StatusMsg = "专业不能为空"
		return resp, errors.New("参数不能为空值")
		// 如果 major 参数为空，则执行相关处理逻辑
	}

	grade := c.PostForm("grade") // 获取 grade 参数
	if grade == "" {
		resp.StatusMsg = "年级不能为空"
		return resp, errors.New("参数不能为空值")
		// 如果 grade 参数为空，则执行相关处理逻辑
	}

	department := c.PostForm("department") // 获取 department 参数
	if department == "" {
		resp.StatusMsg = "方向不能为空"
		return resp, errors.New("参数不能为空值")
		// 如果 department 参数为空，则执行相关处理逻辑
	}

	phone := c.PostForm("phone") // 获取 phone 参数

	sexStr := c.PostForm("sex") // 获取 sex 参数

	isBasicsStr := c.PostForm("is_basics") // 获取 is_basics 参数
	introduce := c.PostForm("introduce")   // 获取 introduce 参数

	user := repository.User{
		ID:         0,
		Name:       name,
		Sex:        sexStr,
		Email:      email,
		Major:      major,
		Introduce:  introduce,
		IsBasics:   isBasicsStr,
		Phone:      phone,
		Grade:      grade,
		Department: department,
	}

	userService := &UserService{userDao: &repository.UserDao{}}

	err := userService.userDao.AddUser(user)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	resp.StatusMsg = "success"
	resp.StatusCode = 200
	return resp, nil
}
func (s *UserService) Select(c *gin.Context) (models.SelectResponse, error) {
	var resp models.SelectResponse
	resp.StatusMsg = "fail"
	resp.StatusCode = 500
	userService := &UserService{
		userDao: &repository.UserDao{},
	}
	rest, err := userService.userDao.SelectDe()
	if err != nil {
		log.Println(err)
		return resp, err
	}
	resp.Dep = rest
	return resp, nil
}

func (s *UserService) Answer(c *gin.Context) (models.AnswerResponse, error) {
	client := &http.Client{}

	question := c.PostForm("question") // 获取 name 参数

	request := models.Answer{Prompt: question,
		UserID:         "#/chat/1697428187899",
		Network:        true,
		WithoutContext: false,
		System:         " ",
		Stream:         false}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	var data = bytes.NewReader(buf)

	req, err := http.NewRequest("POST", "https://api.binjie.fun/api/generateStream?refer__1360=n4mx0DBDnD2DgiDcDIo4Bu7ODkCY%2BRnvx2IhD", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.binjie.fun")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://chat17.aichatos.xyz")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://chat17.aichatos.xyz/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="118", "Microsoft Edge";v="118", "Not=A?Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Mobile Safari/537.36 Edg/118.0.2088.61")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bodyText)

	baseResp := models.AnswerResponse{
		Resp: models.BaseResponse{StatusCode: 200,
			StatusMsg: "success"},
		Answers: str,
	}

	//fmt.Printf("%s\n", bodyText)

	return baseResp, nil
}
