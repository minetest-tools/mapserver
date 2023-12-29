package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type DigilineLcdBlock struct{}

func (this *DigilineLcdBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "digilinelcd")
	o.Attributes["text"] = md["text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
