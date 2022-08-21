# build stage #0
# define base image
FROM golang:1.19-alpine AS stage0
# create work directory
RUN mkdir /prog
# switch to work directory
WORKDIR /prog
# copy all files
ADD . .
# download dependencies
RUN go mod download
# build proglication
# CGO_ENABLED is disabled for cross system compilation
RUN go build -o prog .
# expose port
EXPOSE 8080
# run proglication
CMD ["./prog"]