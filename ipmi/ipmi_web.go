package ipmi

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"serverTemperature/model"
	"time"
)

func WebLogin(conf model.Config, c *cache.Cache) error {
	// base64编码
	ub := base64.StdEncoding.EncodeToString([]byte(conf.Ipmi.User))
	pd := base64.StdEncoding.EncodeToString([]byte(conf.Ipmi.Pwd))
	// 初始化http请求库
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// 登录Request
	post, err := client.R().SetQueryParam("username", ub).
		SetQueryParam("password", pd).
		Post(conf.Ipmi.Host + "/api/session")
	if err != nil {
		logrus.Warnf("login web BMC error: %v", err)
		return err
	}
	var resp model.UserSession
	err = json.Unmarshal(post.Body(), &resp)
	if err != nil {
		logrus.Warnf("login web BMC error: %v", err)
		return err
	}
	cookie := post.Cookies()
	_, isFound := c.Get("QSESSIONID")
	if isFound {
		c.Set("QSESSIONID", cookie[0].Value, 600*time.Second)
		c.Set("CSRFToken", resp.CSRFToken, 600*time.Second)
	} else {
		err = c.Add("QSESSIONID", cookie[0].Value, 600*time.Second)
		if err != nil {
			return err
		}
		err = c.Add("CSRFToken", resp.CSRFToken, 600*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}

func RefreshClock(conf model.Config, c *cache.Cache) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := WebLogin(conf, c)
			if err != nil {
				return
			}
			logrus.Infof("web BMC cookie successful refresh")
		}
	}
}

func ControlFansByWeb(speed int, conf model.Config, c *cache.Cache) error {
	session, _ := c.Get("QSESSIONID")
	token, _ := c.Get("CSRFToken")
	// 初始化http请求库
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	body := map[string]interface{}{
		"MODE":     0,
		"PWMid":    0,
		"PWMvalue": speed,
	}
	// 登录Request
	post, err := client.R().SetCookie(&http.Cookie{
		Name:  "QSESSIONID",
		Value: session.(string),
	}).SetHeader("x-csrftoken", token.(string)).
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetBody(body).
		Put(conf.Ipmi.Host + "/api/system_inventory/set_fan")
	if err != nil {
		logrus.Warnf("control fans speed error: %v", err)
		return err
	}
	if post.StatusCode() != 200 {
		logrus.Warnf("control fans speed request error: %v", err)
		return err
	}
	return nil
}

func SessionExit(conf model.Config, c *cache.Cache) {
	session, _ := c.Get("QSESSIONID")
	token, _ := c.Get("CSRFToken")
	// 初始化http请求库
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// 登录Request
	post, err := client.R().SetCookie(&http.Cookie{
		Name:  "QSESSIONID",
		Value: session.(string),
	}).SetHeader("x-csrftoken", token.(string)).
		SetHeader("X-Requested-With", "XMLHttpRequest").
		Delete(conf.Ipmi.Host + "/api/session")
	if err != nil {
		logrus.Warnf("clear session & exit error: %v", err)
		return
	}
	if post.StatusCode() != 200 {
		logrus.Warnf("clear session & exit error: %v", err)
		return
	}
	return
}
