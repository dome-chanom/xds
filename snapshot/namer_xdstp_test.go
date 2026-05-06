package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXDSTPNamer_FormatsPerResourceType(t *testing.T) {
	n := XDSTPNamer("alpha")

	assert.Equal(t, "xdstp://alpha/envoy.config.listener.v3.Listener/foo.default:50051", n.Listener("foo.default:50051"))
	assert.Equal(t, "xdstp://alpha/envoy.config.route.v3.RouteConfiguration/foo.default:50051", n.RouteConfig("foo.default:50051"))
	assert.Equal(t, "xdstp://alpha/envoy.config.cluster.v3.Cluster/foo.default:grpc", n.Cluster("foo.default:grpc"))
	assert.Equal(t, "xdstp://alpha/envoy.config.endpoint.v3.ClusterLoadAssignment/foo.default:grpc", n.Endpoint("foo.default:grpc"))
}

func TestXDSTPNamer_CacheKeyDistinguishesAuthorities(t *testing.T) {
	a := XDSTPNamer("alpha")
	b := XDSTPNamer("beta")
	legacy := LegacyNamer()

	id := "foo.default:grpc"
	assert.NotEqual(t, legacy.CacheKey(id), a.CacheKey(id))
	assert.NotEqual(t, a.CacheKey(id), b.CacheKey(id))
}
