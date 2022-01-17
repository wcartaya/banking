Si falla la compilación por:
    main.go:3:8: cannot find package "banking/app" in any of:
        /usr/local/go/src/banking/app (from $GOROOT)
        /Users/{userid}/go/src/banking/app (from $GOPATH)
Ejecutar este comando:
export GO111MODULE=auto 