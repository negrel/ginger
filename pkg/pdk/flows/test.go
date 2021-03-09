package flows

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func applyBoxOffset(s styles.Style, offsets boxOffset, props [4]property.ID) {
	s.Set(property.MakeUnit(props[0], value.Unit{Value: offsets[0], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[1], value.Unit{Value: offsets[1], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[2], value.Unit{Value: offsets[2], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[3], value.Unit{Value: offsets[3], UnitID: value.CellUnit}))
}

func applyMargin(s styles.Style, margin boxOffset) {
	applyBoxOffset(s, margin, [4]property.ID{
		property.MarginLeft(), property.MarginTop(), property.MarginRight(), property.MarginBottom(),
	})
}

func applyBorder(s styles.Style, border boxOffset) {
	applyBoxOffset(s, border, [4]property.ID{
		property.BorderLeft(), property.BorderTop(), property.BorderRight(), property.BorderBottom(),
	})
}

func applyPadding(s styles.Style, padding boxOffset) {
	applyBoxOffset(s, padding, [4]property.ID{
		property.PaddingLeft(), property.PaddingTop(), property.PaddingRight(), property.PaddingBottom(),
	})
}

func makeConstraint(minWidth, minHeight, maxWidth, maxHeight int) Constraint {
	return Constraint{
		Min: geometry.Rect(0, 0, minWidth, minHeight),
		Max: geometry.Rect(0, 0, maxWidth, maxHeight),
	}
}
