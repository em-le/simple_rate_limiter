package limiter

type Rate interface {
	Allow() (bool, error)
}
