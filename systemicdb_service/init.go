package systemicdb_service

import "github.com/SamuelBanksTech/SystemicDB-Core/systemicdb"

type SystemicDBService struct {
	tagMap   map[string][][]byte
	keyStore *systemicdb.SystemicDB
}

func Init() *SystemicDBService {

	tm := make(map[string][][]byte)
	ks := systemicdb.NewSystemicDB()

	return &SystemicDBService{
		tagMap:   tm,
		keyStore: ks,
	}
}
