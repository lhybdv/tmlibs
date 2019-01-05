module github.com/lhybdv/tmlibs

go 1.12

require (
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20181128100959-b001fa50d6b2
	github.com/tendermint/go-wire v0.14.1
	github.com/trias-lab/filestore v0.0.2
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc
	gopkg.in/go-playground/validator.v9 v9.24.0
)

replace github.com/tendermint/tmlibs => ./

replace github.com/tendermint/go-wire => github.com/lhybdv/go-wire v0.7.2
