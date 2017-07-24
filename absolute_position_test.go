package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	eps = 0.000001
)

func assertFloatEqual(t *testing.T, got, exp float32) {
	assert.Equal(t, got, exp)
	diff := fabs(got - exp)
	if diff > eps {
		t.Fatalf("got: %.2f, exp: %.2f", got, exp)
	}
}

func TestAbsoluteLayoutWidthHeightStartTop(t *testing.T) {
	config := NewConfig()
	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeStart, 10)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutStartTopEndBottom(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeStart, 10)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetPosition(rootChild0, EdgeEnd, 10)
	NodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 80, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutWidthHeightStartTopEndBottom(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeStart, 10)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetPosition(rootChild0, EdgeEnd, 10)
	NodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestDoNotClampHeightOfAbsoluteNodeToHeightOfItsOverflowHiddenParent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetOverflow(root, OverflowHidden)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeStart, 0)
	NodeStyleSetPosition(rootChild0, EdgeTop, 0)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 100)
	NodeStyleSetHeight(rootChild0Child0, 100)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 50, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, -50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetHeight())
}

func TestAbsoluteLayoutWithinBorder(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeLeft, 10)
	NodeStyleSetMargin(root, EdgeTop, 10)
	NodeStyleSetMargin(root, EdgeRight, 10)
	NodeStyleSetMargin(root, EdgeBottom, 10)
	NodeStyleSetPadding(root, EdgeLeft, 10)
	NodeStyleSetPadding(root, EdgeTop, 10)
	NodeStyleSetPadding(root, EdgeRight, 10)
	NodeStyleSetPadding(root, EdgeBottom, 10)
	NodeStyleSetBorder(root, EdgeLeft, 10)
	NodeStyleSetBorder(root, EdgeTop, 10)
	NodeStyleSetBorder(root, EdgeRight, 10)
	NodeStyleSetBorder(root, EdgeBottom, 10)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 0)
	NodeStyleSetPosition(rootChild0, EdgeTop, 0)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild1, EdgeRight, 0)
	NodeStyleSetPosition(rootChild1, EdgeBottom, 0)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild2, EdgeLeft, 0)
	NodeStyleSetPosition(rootChild2, EdgeTop, 0)
	NodeStyleSetMargin(rootChild2, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild2, EdgeTop, 10)
	NodeStyleSetMargin(rootChild2, EdgeRight, 10)
	NodeStyleSetMargin(rootChild2, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild3, EdgeRight, 0)
	NodeStyleSetPosition(rootChild3, EdgeBottom, 0)
	NodeStyleSetMargin(rootChild3, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild3, EdgeTop, 10)
	NodeStyleSetMargin(rootChild3, EdgeRight, 10)
	NodeStyleSetMargin(rootChild3, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 10, root.LayoutGetLeft())
	assertFloatEqual(t, 10, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 10, root.LayoutGetLeft())
	assertFloatEqual(t, 10, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentFlexEnd(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyFlexEnd)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutJustifyContentCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsCenter(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsCenterOnChildOnly(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild0, AlignCenter)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndTopPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeTop, 10)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndBottomPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsoluteLayoutAlignItemsAndJustifyContentCenterAndLeftPosition(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeLeft, 5)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_align_items_and_justify_content_center_and_right_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexGrow(root, 1)
	NodeStyleSetWidth(root, 110)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPosition(rootChild0, EdgeRight, 5)
	NodeStyleSetWidth(rootChild0, 60)
	NodeStyleSetHeight(rootChild0, 40)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 110, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild0.LayoutGetHeight())
}

func TestPosition_root_with_rtl_should_position_withoutdirection(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetPosition(root, EdgeLeft, 72)
	NodeStyleSetWidth(root, 52)
	NodeStyleSetHeight(root, 52)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 72, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 72, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())
}

func TestAbsolute_layout_percentage_bottom_based_on_parent_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPositionPercent(rootChild0, EdgeTop, 50)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPositionPercent(rootChild1, EdgeBottom, 50)
	NodeStyleSetWidth(rootChild1, 10)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetPositionPercent(rootChild2, EdgeTop, 10)
	NodeStyleSetPositionPercent(rootChild2, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild2.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 160, rootChild2.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_column_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 20)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_row_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 20)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_column_container_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild0, AlignFlexEnd)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 20)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}

func TestAbsolute_layout_in_wrap_reverse_row_container_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetAlignSelf(rootChild0, AlignFlexEnd)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	NodeStyleSetWidth(rootChild0, 20)
	NodeStyleSetHeight(rootChild0, 20)
	NodeInsertChild(root, rootChild0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())
}
