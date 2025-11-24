package results

import (
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type EmailResult struct {
	StringResult
	Email *mail.Address
}

type UuidResult struct {
	StringResult
	UUID uuid.UUID
}

type UriResult struct {
	StringResult
	URI *url.URL
}

type IpResult struct {
	StringResult
	IP net.IP
}

type RegexResult struct {
	StringResult
	Regex *regexp.Regexp
}

type TimeResult struct {
	StringResult
	Time time.Time
}
