# git-acme-commands
git-acme-commands 提供一组扩展的git命令，用来申请、管理 HTTPS 证书，并存储在git仓库中。


## Usage
    git-acme [subcmd] [options...]


## sub-commands

| 命令       | 功能         | 用法                      |
| ---------- | ------------ | ------------------------- |
| add-domain | 添加域名     | git-acme  add-domain [dn] |
| gen-key    | 生成密钥对   |                           |
| request    | 申请证书     |                           |
| list-certs | 显示证书列表 |                           |
