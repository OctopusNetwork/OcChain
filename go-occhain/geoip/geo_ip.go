package geoip

import (
	"compress/gzip"
	"github.com/oschwald/geoip2-golang"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	reader   *geoip2.Reader
	Filename string
)

type LogFunc func(format string, args ...interface{})

func init() {
	Filename, _ = filepath.Abs("geoip.mmdb")
}

// Command returns cobra command for updating geo ip database
func Command(lf LogFunc) *cobra.Command {
	return &cobra.Command{
		Use:   "geoip",
		Short: "Updates geoip database",
		Long:  `Downloads last version of geoIP database created by MaxMind https://www.maxmind.com/en/geoip2-services-and-databases`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := Download(); err != nil {
				lf("Error downloading geoip '%v'", err)
			}
		},
	}
}

// Reader returns Geo2Ip reader instance
func Reader() (*geoip2.Reader, error) {
	if reader == nil {
		var err error
		_, err = os.Stat(Filename)
		if err != nil {
			err = Download()
			if err != nil {
				return nil, err
			}
		}
		reader, err = geoip2.Open(Filename)
		if err != nil {
			return nil, err
		}
	}
	return reader, nil
}

// Download file
func Download() error {
	resp, err := http.Get("http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gz.Close()

	tmpFilename := Filename + ".download"
	tmp, err := os.Create(tmpFilename)
	if err != nil {
		return err
	}
	defer tmp.Close()

	_, err = io.Copy(tmp, gz)
	if err != nil {
		return err
	}
	err = os.Rename(tmpFilename, Filename)
	if err != nil {
		return err
	}
	return nil
}