package snapshot

// LegacyNamer returns the identity Namer — every method returns the id
// unchanged. Use this when emitting resources under their historical names
// for `xds:///foo` URLs.
func LegacyNamer() Namer { return identityNamer{} }

type identityNamer struct{}

func (identityNamer) Listener(id string) string    { return id }
func (identityNamer) RouteConfig(id string) string { return id }
func (identityNamer) Cluster(id string) string     { return id }
func (identityNamer) Endpoint(id string) string    { return id }
func (identityNamer) CacheKey(id string) string    { return id }
