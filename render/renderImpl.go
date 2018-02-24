package render

import (
	"strings"
)

///
// map type def
type TypeAliasPrefix string
type TypeOriginRef string
type TypeMap map[TypeOriginRef]TypeAliasPrefix

///
type RenderType struct {
	AliasType string
	OrigType  string
}

type RenderInputType struct {
	Package string
	Type    string
}

type RenderImpl struct {
	RenderTypes   []RenderType
	ImportPackage map[string]struct{}
}

func NewRenderImpl(input []RenderInputType) *RenderImpl {
	ret := &RenderImpl{
		RenderTypes:   make([]RenderType, 0, len(input)),
		ImportPackage: make(map[string]struct{}, 0),
	}

	for _, t := range input {
		// built in type
		if t.Package == "" || t.Package == "." {
			ret.RenderTypes = append(ret.RenderTypes, processBuiltInType(t.Type))
		} else {
			ret.RenderTypes = append(ret.RenderTypes, processImportType(&t))
			ret.ImportPackage[t.Package] = struct{}{}
		}
	}

	return ret

}

/*
// pkg.Fuck PkgFuck

		Package: "github.com/xxx/pkg",
		Type:    "Fuck",

	// *pkg.Fuck PkgFuckPtr

		Package: "github.com/xxx/pkg",
		Type:    "*Fuck",

		// **pkg.Fuck PkgFuckPtrPtr

		Package: "github.com/xxx/pkg",
		Type:    "**Fuck",
*/
func processImportType(rit *RenderInputType) RenderType {

	if strings.HasPrefix(rit.Type, "*") {
		return processImportPointerType(rit)
	}

	pkgName := processPkgName(rit.Package)
	upperType := stringToUpperHead(rit.Type)
	return RenderType{
		AliasType: stringToUpperHead(pkgName) + upperType,
		OrigType:  pkgName + "." + upperType,
	}
}

func processImportPointerType(rit *RenderInputType) RenderType {
	pkgName := processPkgName(rit.Package)
	lastStar := strings.LastIndex(rit.Type, "*")
	return RenderType{
		AliasType: stringToUpperHead(pkgName + processPointerType(rit.Type)),
		OrigType:  string(rit.Type[0:lastStar+1]) + pkgName + "." + string(rit.Type[lastStar+1:]),
	}

}

func processPkgName(pkgImport string) string {
	splits := strings.Split(pkgImport, "/")
	return splits[len(splits)-1]
}

// could be better
func processBuiltInType(t string) RenderType {

	ptrCase := processPointerType(t)

	return RenderType{
		AliasType: stringToUpperHead(ptrCase),
		OrigType:  t,
	}
}

func processPointerType(t string) string {
	if !strings.HasPrefix(t, "*") {
		return t
	}
	splits := strings.Split(t, "*")
	refType := splits[len(splits)-1]
	refType = stringToUpperHead(refType)
	for i := 0; i < len(splits)-1; i++ {
		refType += "Ptr"
	}
	return refType
}

func stringToUpperHead(t string) string {
	return strings.ToUpper(string(t[0])) + string(t[1:])
}

func (r *RenderImpl) AllTypes() []RenderType {
	return r.RenderTypes
}

func (r *RenderImpl) ImportPackages() map[string]struct{} {
	return r.ImportPackage
}
