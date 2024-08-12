templateFile=$GOPATH/src/github.com/go-easygen/easygen/test/commandlineFlag-Name
[ -s $templateFile.tmpl ] || templateFile=/usr/share/gocode/src/github.com/go-easygen/easygen/test/commandlineFlag-Name
[ -s $templateFile.tmpl ] || templateFile=/usr/share/doc/easygen/examples/commandlineFlag-Name
[ -s $templateFile.tmpl ] || {
  echo No template file found
  exit 1
}
#echo Using $templateFile.tmpl

gpkg=gistpost
easygen $templateFile ${gpkg}_cli | goimports > ${gpkg}_cliGen.go
echo ${gpkg}_cliGen.go generated/updated.
