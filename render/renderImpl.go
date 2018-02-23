package render

///
// map type def
type TypeAliasPrefix string
type TypeOriginRef string
type TypeMap map[TypeOriginRef]TypeAliasPrefix

///
type ImportedType struct {
	Package   string
	AliasType string
	OrigType  string
}
type RenderImpl struct {
	BuiltInTypes []string
	ImportTypes  []ImportedType
}

func NewRenderImpl(BuiltIn []string, Import []ImportedType) *RenderImpl {
	return &RenderImpl{
		BuiltInTypes: BuiltIn,
		ImportTypes:  Import,
	}
}

func (r *RenderImpl) AllTypes() []*ImportedType {
	return []*ImportedType{
		&ImportedType{
			AliasType: "String",
			OrigType:  "string",
		},
		&ImportedType{
			AliasType: "Int",
			OrigType:  "int",
		},
	}
}

func (r *RenderImpl) ImportPackages() []string {
	return nil
}
