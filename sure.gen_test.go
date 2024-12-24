package simplejsonx

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_pkg_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

type ObjectType struct{}

func TestGen(t *testing.T) {
	pkgPath := syntaxgo_reflect.GetPkgPathV2[ObjectType]()
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.SOFT, pkgPath).WithNewPkgName("simplejson_soft"))
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.MUST, pkgPath).WithNewPkgName("simplejson_must"))
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.OMIT, pkgPath).WithNewPkgName("simplejson_omit"))
}
