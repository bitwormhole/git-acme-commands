package data

// // Mix 把几个 VO 混合在一起
// func Mix(list ...*VO) *VO {
// 	dst := new(VO)
// 	for _, src := range list {
// 		if src.ACME != nil {
// 			dst.ACME = src.ACME
// 		}
// 		dst.Accounts = append(dst.Accounts, src.Accounts...)
// 		dst.Domains = append(dst.Domains, src.Domains...)
// 		dst.KeyPairs = append(dst.KeyPairs, src.KeyPairs...)
// 		dst.Dirs = append(dst.Dirs, src.Dirs...)
// 	}
// 	return dst
// }
