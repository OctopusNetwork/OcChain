package geoip

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"net"
	"github.com/stretchr/testify/assert"
)

// todo: Write meaningful test
func TestGeoIP(t *testing.T) {
	require.IsType(t, &cobra.Command{}, Command(log.Printf))
}

func TestIP(t *testing.T) {
	r, err := Reader()
	require.NoError(t, err)
	ip := net.ParseIP("2.60.177.126")
	c, _ := r.City(ip)
	assert.Equal(t, "Omsk", c.City.Names["en"])
	ip = net.ParseIP("195.19.132.64")
	c, _ = r.City(ip)
	assert.Equal(t, "Yekaterinburg", c.City.Names["en"])
	ip = net.ParseIP("77.45.128.13")
	c, _ = r.City(ip)
	assert.Equal(t, "Voronezh", c.City.Names["en"])
}