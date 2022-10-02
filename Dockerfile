FROM golang:latest

# Remote Containerがデフォルトで開くディレクトリは/workspace
WORKDIR /workspace

# install go tools
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/stamblerre/gocode@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
# RUN go install github.com/lib/pq@latest