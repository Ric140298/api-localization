##################################
# STEP 1 build executable binary #
##################################
FROM golang:1.14.2 as builder
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/api-localization
# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# set variables for environment
ENV ENVIRONMENT "dev"
ENV LOG_LEVEL 0
# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get  -u github.com/stretchr/testify/assert && \
    go get -d -v ./... && \
    go install -v ./...
    
    #cd controllers && go test
# Craete binary

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o app .

##############################
# STEP 2 get ca-certificates #
##############################

FROM alpine:latest as certs
RUN apk --update add ca-certificates


###############################################################
# STEP 3 build a small image with ca-certificates from alpine #
###############################################################
FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy our static executable and required files.
WORKDIR /go/src/api-localization

# Copy binary and json files
COPY --from=builder /go/src/api-localization/app /go/src/api-localization/app
# Run the app binary.
CMD ["./app"]
EXPOSE $PORT
