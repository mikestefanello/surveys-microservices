module github.com/mikestefanello/surveys-microservices/vote-service

go 1.15

require (
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mikestefanello/surveys-microservices/survey-service v0.0.0-20200922175831-2304c4cbfd84
	github.com/rs/zerolog v1.20.0
	github.com/satori/go.uuid v1.2.0
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.32.0
)
