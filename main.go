package main

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"ssl/aliyun"
	"ssl/ding"
	"ssl/ecs"
	"ssl/encrypt"
	"time"
)

func main() {
	InitConfig()
	cheSSl()
}

func domain() []string {
	domainList := []string{
		"https://baidu.com",
	}
	return domainList
}

func cheSSl()  {
	keyid := viper.GetString("ali.keyid")
	secret := viper.GetString("ali.secret")
	domainList := domain()
	for _,v := range domainList {
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 10 * time.Second,
		}
		rep,err := client.Get(v)
		if err != nil {
			panic(err)
		}
		cert :=  rep.TLS.PeerCertificates[0]

		if cert.NotAfter.Sub(time.Now()) <= 15 * 24 * 60 * 60 *  time.Second && cert.NotAfter.Sub(time.Now()) >= 14 * 24 *60 * 60 * time.Second {
			d := ding.Webhook{
				AccessToken: "钉钉token", // 上面获取的 access_token
				Secret:      "",                                                                 // 上面获取的加签的值
			}

			certTime, _ := time.Parse("2006-01-02", cert.NotAfter.Local().Format("2006-01-02"))
			nowTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
			remaining := certTime.Sub(nowTime)

			certMessage := fmt.Sprintf("域名：%v\nssl证书到期时间为：%v\n剩余%v天请注意替换", v, cert.NotAfter.Local().Format("2006-01-02"), remaining.Hours()/24)
			fmt.Println(certMessage)
			_ = d.SendMessage(certMessage)


			err,res :=aliyun.CheckCdn(keyid,secret,v)
			if err != nil   {
				fmt.Println(err)
			}
			if res != true {
					encrypt.CreatSsl(v) //创建免费证书
					pem := encrypt.GetCASsl(v)	//获取证书pem
					key := encrypt.GetKeySsl(v)	//获取证书key
					err := aliyun.ModifyCert(keyid,secret,v,pem,key)
					if err != nil {
						fmt.Println(err)
					}
					_ = d.SendMessage(fmt.Sprintf("更新%s证书成功,证书在cdn里面"))
				} else {
					if ecs.CheckSslPath(v) != true {
						fmt.Println("ssl证书不在本机，请手动更新")
						return
					}
					encrypt.CreatSsl(v)	//创建免费证书
					ecs.UpdateSsl(v)	//更新ecs证书
					_ = d.SendMessage(fmt.Sprintf("更新%s证书成功,证书在ecs里面"))
			}
		}
	}
}

func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/config")

}

