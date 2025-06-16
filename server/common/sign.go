package common

import (
	stdpath "path"

	"alist/internal/conf"
	"alist/internal/model"
	"alist/internal/setting"
	"alist/internal/sign"
)

func Sign(obj model.Obj, parent string, encrypt bool) string {
	if obj.IsDir() || (!encrypt && !setting.GetBool(conf.SignAll)) {
		return ""
	}
	return sign.Sign(stdpath.Join(parent, obj.GetName()))
}
