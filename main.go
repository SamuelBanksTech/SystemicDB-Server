package main

import (
	"flag"
	"github.com/SamuelBanksTech/SystemicDB-Server/core"
)

func main() {
	var suf core.StartupFlags

	suf.Port = flag.Int("port", 5222, "The listening port for GRPC communication")
	suf.Mode = flag.String("mode", "master", "Use master or slave, if slave then a master ip must be specified using the -mip flag")
	suf.Mip = flag.String("mip", "0.0.0.0", "Master ip, only used if mode is set to \"slave\"")
	flag.Parse()

	core.Bootstrap(suf)
}
