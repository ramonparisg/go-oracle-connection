#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.13.6-buster AS build_stage
RUN mkdir -p /go/src/
WORKDIR /go/src/
COPY . /go/src/
RUN go get github.com/godror/godror
RUN go get github.com/go-eden/slf4go
RUN go build -o goapp

#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM oracle/instantclient:19 AS final_stage
RUN mkdir -p /go/src/
COPY --from=build_stage /go/src/ /go/src/
WORKDIR /go/src/
ENTRYPOINT ./goapp