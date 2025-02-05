package testdata

import _ "embed"

//go:embed berry-1.json
var Berry1 []byte

//go:embed berry-32.json
var Berry32 []byte

//go:embed berry-firmness-2.json
var BerryFirmness2 []byte

//go:embed berry-firmness-l1o4.json
var BerryFirmnessL1O4 []byte

//go:embed berry-flavor-1.json
var BerryFlavor1 []byte

//go:embed berry-flavor-l3o2.json
var BerryFlavorL3O2 []byte

//go:embed berry-l1o1.json
var BerryL1O1 []byte

//go:embed berry-l3o0.json
var BerryL3O0 []byte

//go:embed berry-l1o63.json
var BerryL1O63 []byte

//go:embed berry-l1o64.json
var BerryL1O64 []byte
