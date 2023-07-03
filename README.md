        使用说明
        config.yml下面添加阿里云key(注意key权限，最好只能操作cdn和云解析权限，以免不法分子利用)
        
        本项目主要开发ssl半自动化，写的代码很粗糙，很烂，（勿喷）本人吸取所有意见，欢迎大佬带我做项目
        
        在此之前linux上需要安装acme.sh地址：https://github.com/acmesh-official/acme.sh/wiki/%E8%AF%B4%E6%98%8E
        
        主要功能是为了实现公司ssl过多手工替换太麻烦，域名直接加载domain函数上即可，
        然后调用了let’s encrypt免费证书申请，使用acme.sh管理，实现了判断域名是否在cdn里面或者在ecs里面
        如果是刚好在cdn里面则调用了阿里云的sdk进行更新证书，如果刚好在ecs里面则使用acme.sh更新nginx证书。
        其他情况则手动更新。
