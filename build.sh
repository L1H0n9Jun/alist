appName="alist"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor="L1H0n9Jun"
gitCommit=$(git log --pretty=format:"%h" -1)

version=v4.0.0
webVersion=v4.0.0

echo "backend version: $version"
echo "frontend version: $webVersion"

ldflags="\
-w -s \
-X 'alist/internal/conf.BuiltAt=$builtAt' \
-X 'alist/internal/conf.GoVersion=$goVersion' \
-X 'alist/internal/conf.GitAuthor=$gitAuthor' \
-X 'alist/internal/conf.GitCommit=$gitCommit' \
-X 'alist/internal/conf.Version=$version' \
-X 'alist/internal/conf.WebVersion=$webVersion' \
"

mkdir -p build

flags="$ldflags --extldflags '-fPIE'"
GOOS=linux GOARCH=amd64 GOAMD64=v3 CGO_ENABLED=1 \
go build -trimpath -ldflags="$flags" -tags=jsoniter \
-o build/alist .

upx -9 build/alist
