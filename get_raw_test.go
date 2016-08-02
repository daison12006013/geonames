package geonames

import (
	"bufio"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRaw(t *testing.T) {
	Convey("Testing getRaw function", t, func() {
		var err error
		var s *bufio.Scanner

		txtName := "timeZones.txt"
		txtURL := geonamesURL + txtName

		s, err = getRaw(txtURL, txtName)
		So(err, ShouldBeNil)
		So(s, ShouldNotBeNil)

		zipName := "AD.zip"
		zipURL := fmt.Sprintf("http://download.geonames.org/export/zip/%s.zip", "AD")

		s, err = getRaw(zipURL, zipName)
		So(err, ShouldBeNil)
		So(s, ShouldNotBeNil)
	})
}
