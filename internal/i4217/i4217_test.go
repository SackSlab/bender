package i4217

import "testing"

func Test(t *testing.T) {
	validIso := ISO4217{Name: "PYG", Code: "600", Units: 0}
	invalidIso := ISO4217{Name: "SWC", Code: "000", Units: 0} // Republic credits, from Stars Wars

	testCases := []struct {
		desc          string
		value         ISO4217
		expectedFound bool
		useCode       bool
	}{
		{
			desc:          "Should return true if the name found for given code",
			value:         validIso,
			expectedFound: true,
			useCode:       false,
		},
		{
			desc:          "Should return false if the name not found for given code",
			value:         invalidIso,
			expectedFound: false,
			useCode:       false,
		},
		{
			desc:          "Should return true if the code found for given name",
			value:         validIso,
			expectedFound: true,
			useCode:       true,
		},
		{
			desc:          "Should return false if the code not found for given name",
			value:         invalidIso,
			expectedFound: false,
			useCode:       true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var result ISO4217
			var found bool

			if tC.useCode {
				result, found = ByCode(tC.value.Code)
			} else {
				result, found = ByName(tC.value.Name)
			}

			if tC.expectedFound != found {
				t.Errorf("invalid resolution for iso value %v", result)
			}
		})
	}
}
