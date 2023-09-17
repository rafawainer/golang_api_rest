module github.com/rafawainer/golang_api_rest/api-go-rest

replace (
github.com/rafawainer/golang_api_rest/api-go-rest/models => ./models
github.com/rafawainer/golang_api_rest/api-go-rest/controllers => ./controllers
github.com/rafawainer/golang_api_rest/api-go-rest/routes => ./routes
)

go 1.21.0
