FROM golang:latest
RUN apt-get update && apt-get install -y software-properties-common && apt-get update && add-apt-repository ppa:niko2040/e19

RUN add-apt-repository ppa:niko2040/e19 && apt-get update && apt-get install -y --no-install-recommends \
		mingw-w64 \
		enlightenment \
		libeina-dev \
		libeina1 \
		libefl \
		libefl-dev

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]