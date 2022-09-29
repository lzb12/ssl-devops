// This file is auto-generated, don't edit it. Thanks.
package aliyun

import (
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}

//判断域名有没有加cdn
func CheckCdn (accessKeyId,accessKeySecret string,domain string) (_err error,bool2 bool) {
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err,false
	}

	describeUserDomainsRequest := &cdn20180510.DescribeUserDomainsRequest{
		DomainName: tea.String(domain),
	}
	runtime := &util.RuntimeOptions{}
		// 复制代码运行请自行打印 API 的返回值
		restult, _err := client.DescribeUserDomainsWithOptions(describeUserDomainsRequest, runtime)
		if _err != nil {
			return _err,false
		}
		if len(restult.Body.Domains.PageData) > 0 {
			return nil,true
		}
		return nil,false
}

//cdn修改证书
func ModifyCert (accessKeyId,accessKeySecret string,domain string,pem,cert string) (_err error) {
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	setDomainServerCertificateRequest := &cdn20180510.SetDomainServerCertificateRequest{
		DomainName: tea.String(domain),
		ServerCertificateStatus: tea.String("on"),
		ServerCertificate: tea.String(pem),
		PrivateKey: tea.String(cert),
		ForceSet: tea.String("1"),
		CertName: tea.String(domain),
		CertType: tea.String("upload"),
	}
	runtime := &util.RuntimeOptions{}
		_, _err = client.SetDomainServerCertificateWithOptions(setDomainServerCertificateRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
}




