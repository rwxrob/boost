FROM ubuntu

LABEL maintainer="rob@rwx.gg"
LABEL io.k8s.description="Linux Beginner Boost (made with ❤️ by rwxrob)"
LABEL org.opencontainers.image.source="https://github.com/rwxrob/boost"

ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked apt-get update -q
RUN apt-get install -yq --no-install-recommends \
  git ed vim neovim emacs nano sudo man ca-certificates jq \
  shfmt shellcheck nodejs npm && \
  apt clean && rm -rf /var/lib/apt/lists/*

COPY . /

RUN echo "ubuntu:password" | chpasswd
RUN chown -R ubuntu:ubuntu /home/ubuntu

USER ubuntu
WORKDIR /home/ubuntu

RUN nvim --headless +PlugInstall +qall
RUN nvim --headless +'CocInstall -sync coc-json coc-tsserver' +qa
