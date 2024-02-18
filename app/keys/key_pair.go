package keys

import (
	"crypto"

	"github.com/starter-go/afs"
)

// KeyPair 密钥对
type KeyPair interface {
	Provider() KeyPairProvider

	Signer() crypto.Signer
}

// KeyPairLoader 密钥对加载器
type KeyPairLoader interface {
	Load(file afs.Path) (KeyPair, error)
}

// KeyPairSaver 密钥对保存器
type KeyPairSaver interface {
	Save(kp KeyPair, file afs.Path) error
}

// KeyPairGenerator 密钥对生成器
type KeyPairGenerator interface {
	Generate() (KeyPair, error)
}

// KeyPairProvider 密钥对提供者
type KeyPairProvider interface {
	Info() *KeyPairProviderInfo

	Generator() KeyPairGenerator
	Loader() KeyPairLoader
	Saver() KeyPairSaver
}

// KeyPairProviderInfo 密钥对提供者信息
type KeyPairProviderInfo struct {
	Algorithm string
	Provider  KeyPairProvider
}

// KeyPairProviderManager 密钥对提供者管理器
type KeyPairProviderManager interface {
	FindProvider(algorithm string) (KeyPairProvider, error)
}
