module github.com/Clever/sphinx

go 1.22.1

require (
	github.com/Clever/leakybucket v1.1.0
	github.com/aws/aws-sdk-go v1.55.5
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/stretchr/testify v1.9.0
	gopkg.in/Clever/kayvee-go.v6 v6.27.0
	gopkg.in/tylerb/graceful.v1 v1.2.15
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.1-0.20200118195451-b537c054d4b4 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace gopkg.in/Clever/kayvee-go.v6 => github.com/caido/dependency-kayvee-go/v6 v6.30.0
