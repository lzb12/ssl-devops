package encrypt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//创建ssl证书
func CreatSsl(domain string) {

	cmd := fmt.Sprintf(`~/.acme.sh/acme.sh --issue --dns dns_ali -d %s -d *.%s`,domain,domain)
	//注意脚本要写绝对路径

	out,err := exec.Command("/bin/bash","-c",cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))

}


//判断证书目录是否存在
func CheckSslDir(path string) (bool,error) {
	_,err := os.Stat(path)
	if err == nil{
		return true,nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err){	//如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false,nil
	}
	return false,err//如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

//获取cassl证书内容
func GetCASsl(domain string) string {
	sslpath := fmt.Sprintf("/root/.acme.sh/%s/",domain)
	res,err	:= CheckSslDir(sslpath)
	if err != nil && res != false {
		return ""//目录不存在请检查ssl证书是否申请成功
	}

	capath := fmt.Sprintf("/root/.acme.sh/%s/ca.cer",domain)
	//读取证书ca内容
	f,err := ioutil.ReadFile(capath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(f))
	return string(f)
}

//获取key证书内容
func GetKeySsl(domain string) string {
	sslpath := fmt.Sprintf("/root/.acme.sh/%s/",domain)
	res,err	:= CheckSslDir(sslpath)
	if err != nil && res != false {
		return ""//目录不存在请检查ssl证书是否申请成功
	}

	keypath := fmt.Sprintf("/root/.acme.sh/%s/%s.key" ,domain,domain)
	//读取证书ca内容
	k,err := ioutil.ReadFile(keypath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(k))
	return string(k)

}