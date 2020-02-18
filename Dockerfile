FROM alpine:3.7
WORKDIR /
COPY kustomize /kustomize
COPY .registry /.registry

# This container is meant to be used as CSI storage rather than a processing unit.
ENTRYPOINT ["find", "/kustomize"]