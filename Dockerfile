FROM golang:1.22-alpine3.19 AS buildStage
WORKDIR /product-svc
COPY . ./
RUN go mod download
RUN go build -o ./product-svc ./cmd

FROM scratch AS release-stage 
WORKDIR /
COPY --from=buildStage /product-svc/product-svc /product-svc
COPY --from=buildStage /product-svc/dev.env /

EXPOSE 50052

ENTRYPOINT [ "/product-svc" ]