module github.com/Kushan-/go--docker

go 1.13

require (
	github.com/Kushan-/go--docker/goodbye v0.0.0
	github.com/gofiber/fiber v1.12.0
)

replace (
	github.com/Kushan-/go--docker/goodbye => ./goodbye
)
