package snapshot

// Namer translates a resource's logical id into the cache name it is stored
// under. Implementations live alongside this file:
//
//   - LegacyNamer (namer_legacy.go) returns the id verbatim — preserves
//     historical naming so `xds:///foo` keeps working unchanged.
//   - XDSTPNamer (namer_xdstp.go) prefixes the id with
//     `xdstp://<authority>/<type>/` so the resource is addressable via gRPC
//     A47 federation under that authority, e.g. `xds://<authority>/foo`.
//
// Resource emission code threads a Namer through every name and inter-resource
// reference so that one pass produces a self-consistent set of resources whose
// internal references stay within a single naming namespace.
type Namer interface {
	Listener(id string) string
	RouteConfig(id string) string
	Cluster(id string) string
	Endpoint(id string) string

	// CacheKey produces a stable, namer-unique key for caches that may hold
	// entries from multiple Namers for the same logical id (for example, the
	// per-endpoint resource cache emits both legacy and xdstp variants).
	CacheKey(id string) string
}
