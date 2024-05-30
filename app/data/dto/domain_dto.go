package dto

// Domain ...
type Domain struct {
	ID   string     // id of segment
	Name DomainName // 域名
	Ref  string     // 域配置文件('domain.config')的相对路径
}
