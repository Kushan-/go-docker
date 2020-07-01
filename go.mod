module github.com/Kushan-/go-docker

go 1.13

require (
	github.com/Kushan-/go-docker/pkg/goodbye v0.0.0
	github.com/Kushan-/go-docker/pkg/handler v0.0.0
	github.com/Kushan-/go-docker/pkg/router v0.0.0
	// github.com/Kushan-/go-docker/middleware v0.0.0
	github.com/gofiber/fiber v1.12.1
	github.com/gofiber/jwt v0.1.1
	github.com/gofiber/logger v0.2.3 // indirect
)

replace (
	github.com/Kushan-/go-docker/pkg/goodbye => ./pkg/goodbye
	github.com/Kushan-/go-docker/pkg/handler => ./pkg/handler
	github.com/Kushan-/go-docker/pkg/router => ./pkg/router
	// github.com/Kushan-/go-docker/middleware => ./middleware
)
