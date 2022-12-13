<div align="center">
<img height="300" src="./res/logo-server-trasparent.webp" alt="SystemicDB Logo" />

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Twitter Handle](https://img.shields.io/twitter/follow/samuelbankstech)](https://twitter.com/samuelbankstech)
</div>

# SystemicDB Server

An extremely fast in-memory database server with a gRPC interface, and a core AVL tree at its core. 

## Overview

- ✅ Core storage is a custom implementation of an AVL Tree for scalability, speed, and stability
  - based on [SystemicDB Core](https://github.com/SamuelBanksTech/SystemicDB-Core)
- ✅ Interface is implemented with gRPC allowing ease of adoption and speed of communication
- ✅ Data stored using a MsgPack implementation for memory efficiency
- ✅ Automatic removal of data after user specified expiry time

