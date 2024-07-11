package gongreqif

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-gongreqif/dist/ng-github.com-thomaspeugeot-gongreqif
var NgDistNg embed.FS
