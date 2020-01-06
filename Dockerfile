# For build
FROM golang:alpine AS build
COPY . /app
RUN cd /app && go build -o pingpong

# For run
FROM alpine
WORKDIR /app
COPY --from=build /app/pingpong /app/

# Run the image as a non-root user
RUN adduser -D user
USER user

CMD ["./pingpong"]
