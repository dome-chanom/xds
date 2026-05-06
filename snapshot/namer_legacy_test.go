package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLegacyNamer_ReturnsIdVerbatim(t *testing.T) {
	n := LegacyNamer()
	assert.Equal(t, "foo.default:50051", n.Listener("foo.default:50051"))
	assert.Equal(t, "foo.default:50051", n.RouteConfig("foo.default:50051"))
	assert.Equal(t, "foo.default:grpc", n.Cluster("foo.default:grpc"))
	assert.Equal(t, "foo.default:grpc", n.Endpoint("foo.default:grpc"))
	assert.Equal(t, "foo.default:grpc", n.CacheKey("foo.default:grpc"))
}
