package core

import "github.com/SamuelBanksTech/SystemicDB-Server/systemicdb_service"

type StartupFlags struct {
	Port *int
	Mode *string
	Mip  *string
}

func Bootstrap(suf StartupFlags) {

	sdb := systemicdb_service.Init()

	InitGrpcServer(*suf.Port, sdb)
}
