package gospdk

type ConstructTargetParams struct {
	Name           string
	AliasName      string
	LunNameIdPairs string
	PGIGMappings   string

	QueueDepth    int
	ChapDisabled  int
	ChapEnabled   int
	ChapRequired  int
	ChapMutual    int
	ChapAuthGroup int
}
