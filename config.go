package torus

type Config struct {
	DataDir         string
	StorageSize     uint64
	MetadataAddress string //# "127.0.0.1:2379"
	ReadCacheSize   uint64
	ReadLevel       ReadLevel
	WriteLevel      WriteLevel
}
