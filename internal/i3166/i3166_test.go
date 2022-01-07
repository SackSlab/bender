package i3166

import "testing"

func TestMap_Init(t *testing.T) {
	name, found := codes["PRY"]
	if !found {
		t.Error("the map of codes its not initializated")
	}

	code, found := names[name]
	if !found {
		t.Errorf("the map of codes its not initializated, code: %s", code)
	}
}

func Test(t *testing.T) {
	validIso := ISO3166{Name: "Paraguay", Code: "PRY"}
	invalidIso := ISO3166{Name: "Wonderland", Code: "WDL"}

	testCases := []struct {
		desc          string
		value         ISO3166
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
			var result ISO3166
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
