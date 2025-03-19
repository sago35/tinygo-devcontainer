FROM docker.io/tinygo/tinygo-dev:0.36.0

RUN go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/sago35/tinygo-autocmpl@latest

RUN apt update && \
    apt install openocd -y

RUN echo 'eval "$(tinygo-autocmpl --completion-script-bash)"' >> ~/.bashrc
