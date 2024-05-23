package models

// PostalCode represents a single countries postal code
type PostalCode struct {
	CountryIso2Code string  `json:"country_iso2_code"` // country code      : iso country code, 2 characters
	PostalCode      string  `json:"postal_code"`       // postal code       : varchar(20)
	PlaceName       string  `json:"place_name"`        // place name        : varchar(180)
	AdminName1      string  `json:"admin_name1"`       // admin name1       : 1. order subdivision (state) varchar(100)
	AdminCode1      string  `json:"admin_code1"`       // admin code1       : 1. order subdivision (state) varchar(20)
	AdminName2      string  `json:"admin_name2"`       // admin name2       : 2. order subdivision (county/province) varchar(100)
	AdminCode2      string  `json:"admin_code2"`       // admin code2       : 2. order subdivision (county/province) varchar(20)
	AdminName3      string  `json:"admin_name3"`       // admin name3       : 3. order subdivision (community) varchar(100)
	AdminCode3      string  `json:"admin_code3"`       // admin code3       : 3. order subdivision (community) varchar(20)
	Latitude        float64 `json:"latitude"`          // latitude          : estimated latitude (wgs84)
	Longitude       float64 `json:"longitude"`         // longitude         : estimated longitude (wgs84)
	Accuracy        int     `json:"accuracy"`          // accuracy          : accuracy of lat/lng from 1=estimated to 6=centroi}
}
