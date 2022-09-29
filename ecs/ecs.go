package ecs

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"ssl/encrypt"
)

//判断证书是否在本机服务器
func CheckSslPath(domain string) bool {
	iprecords,_ := net.LookupIP(domain)
	ipAddressToBoost := fmt.Sprintf("%d.%d.%d.%d",iprecords[0][12],iprecords[0][13],iprecords[0][14],iprecords[0][15])
	ecsip := "120.78.177.83"
	if ecsip == ipAddressToBoost {
		return true
	} else {
		return false
	}

}

//更新ssl证书
func UpdateSsl(domain string)  {
	sslpath := fmt.Sprintf("/etc/nginx/ssl/%s/",domain)

	res ,err := encrypt.CheckSslDir(sslpath)
	if err != nil {
		fmt.Println(err)
	}
	if res != true {
		os.Mkdir(sslpath,755)
	}

	cmd := fmt.Sprintf(`~/.acme.sh/acme.sh --installcert -d '%s' --key-file /etc/nginx/ssl/%s/'%s.key' --fullchain-file /etc/nginx/ssl/%s/'%s.pem' --reloadcmd "nginx -s reload"`,domain,domain,domain,domain)
	out,err := exec.Command("/bin/bash","-c",cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))

}


