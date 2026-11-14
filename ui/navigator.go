package ui

import (
	"github.com/xmhrnx/giode"

	"gedanken/ui/theme"
)

// Navigator is a stack of screens that animates between them: a new
// screen slides in from the right while the old one slides out to the
// left, and going back reverses it. The system Back button pops.
type Navigator struct {
	stack   []giode.Element
	current *giode.TransitionElement
	leaving *giode.TransitionElement
}

// Directions of travel, as the transition kinds that bring a screen in.
var (
	fromRight = giode.Fade | giode.SlideLeft
	fromLeft  = giode.Fade | giode.SlideRight
)

// NewNavigator shows root without animating it.
func NewNavigator(root giode.Element) *Navigator {
	return &Navigator{
		stack:   []giode.Element{root},
		current: giode.Transition(opaque(root), giode.Fade, giode.TransitionNormal, giode.EaseOut).Appear(false),
	}
}

// Push shows a new screen on top of the stack.
func (n *Navigator) Push(screen giode.Element) {
	n.stack = append(n.stack, screen)
	n.swap(screen, fromRight, fromLeft)
}

// Replace swaps the top screen for another, as moving forward.
func (n *Navigator) Replace(screen giode.Element) {
	n.stack[len(n.stack)-1] = screen
	n.swap(screen, fromRight, fromLeft)
}

// Pop returns to the previous screen and reports whether there was one.
func (n *Navigator) Pop() bool {
	if len(n.stack) < 2 {
		return false
	}
	n.stack = n.stack[:len(n.stack)-1]
	n.swap(n.stack[len(n.stack)-1], fromLeft, fromRight)
	return true
}

// swap animates next in and the current screen out. exit is the kind
// whose reverse carries the old screen away; it leaves faster than the
// new screen arrives, so the two never read as one muddled page.
func (n *Navigator) swap(next giode.Element, enter, exit giode.TransitionKind) {
	old := n.current
	old.Kind(exit).Duration(giode.TransitionFast)
	old.OnDone(func(bool) {
		if n.leaving == old {
			n.leaving = nil
		}
	})
	old.Hide()
	n.leaving = old
	n.current = giode.Transition(opaque(next), enter, giode.TransitionNormal, giode.EaseOut)
}

// opaque backs a screen with the page color, so an arriving screen
// covers the one leaving beneath it.
func opaque(screen giode.Element) giode.Element {
	return giode.Box(giode.Styles{Width: -1, Height: -1, Background: theme.Background}, screen)
}

// Layout draws the leaving screen under the arriving one.
func (n *Navigator) Layout(gtx giode.Context) giode.Dimensions {
	if len(n.stack) > 1 {
		giode.HandleBack(func() { n.Pop() })
	}
	kids := []giode.Element{giode.Box(giode.Styles{Width: -1, Height: -1, ZIndex: 1}, n.current)}
	if n.leaving != nil {
		kids = append(kids, giode.Box(giode.Styles{
			Position: giode.PositionAbsolute, Top: giode.Flush, Left: giode.Flush,
			Right: giode.Flush, Bottom: giode.Flush,
		}, n.leaving))
	}
	return giode.Box(giode.Styles{Width: -1, Height: -1}, kids...).Layout(gtx)
}
