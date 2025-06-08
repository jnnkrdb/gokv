package messaging

// request type

type RequestType string

func (rt RequestType) String() string {
	return string(rt)
}

const (
	RT_Register    = "req-register"
	RT_KeepConn    = "req-keepc"
	RT_HealthCheck = "req-healthz"
	RT_Dropped     = "req-dropped"
)
