# git-acme-commands
git-acme-commands 提供一组扩展的git命令，用来申请、管理 HTTPS 证书，并存储在git仓库中。


## Usage
    git-acme [subcmd] [options...]


## sub-commands

| 命令    | 功能                            |
| ------- | ------------------------------- |
| add     | 添加域名                        |
| prepare | 生成密钥对                      |
| request | 申请证书                        |
| certs   | 显示证书列表                    |
| domains | 显示域名列表                    |
| fetch   | 从 web 服务器下载当前使用的证书 |
