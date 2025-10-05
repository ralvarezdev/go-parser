package json

type (
	// Writer is the interface for writing JSON tags
	Writer interface {
		OverwriteTags(
			filePath string,
			structJSONTagMapper StructsTagsMapper,
		) error
		HideStructsTags(
			filePath string,
			structsFieldsMapper StructsFieldsMapper,
		) error
		HideFilesTags(
			goFileStructFieldsMapper FilesStructsFieldsMapper,
		) error
	}
)
