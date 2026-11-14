package curriculum

import (
	"fmt"
	"strings"
)

// latexCommands are the commands Giode's formula package typesets.
var latexCommands = toSet(`frac dfrac tfrac cfrac binom sqrt text textrm textnormal mbox textup
mathrm operatorname mathup rm mathbf boldsymbol bm textbf mathit textit left right
big Big bigg Bigg bigl Bigl biggl Biggl bigr Bigr biggr Biggr bigm Bigm
vec hat widehat bar overline dot ddot tilde widetilde overrightarrow underline check breve acute grave
begin end displaystyle textstyle nonumber notag label limits nolimits not lbrace rbrace Vert vert cr newline
alpha beta gamma delta epsilon varepsilon zeta eta theta vartheta iota kappa lambda mu nu xi omicron pi
varpi rho varrho sigma varsigma tau upsilon phi varphi chi psi omega Gamma Delta Theta Lambda Xi Pi Sigma
Upsilon Phi Psi Omega nabla partial hbar ell infty emptyset varnothing angle degree prime forall exists neg
Re Im aleph ldots dots cdots vdots ddots triangle square star dagger top bot backslash langle rangle lfloor
rfloor lceil rceil lvert rvert lVert rVert pm mp times cdot div ast circ bullet cup cap wedge land vee lor
oplus otimes setminus leq le geq ge neq ne approx equiv propto sim simeq cong ll gg lesssim gtrsim in notin
ni subset supset subseteq supseteq perp parallel mid doteq to rightarrow leftarrow gets leftrightarrow
Rightarrow Leftarrow Leftrightarrow implies impliedby iff mapsto longrightarrow longleftarrow uparrow
downarrow rightleftharpoons coloneqq triangleq colon sum prod coprod bigcup bigcap bigoplus bigotimes int
iint iiint oint oiint sin cos tan cot sec csc arcsin arccos arctan sinh cosh tanh coth exp ln log lg det
dim ker deg arg gcd Pr hom lim limsup liminf max min sup inf thinspace medspace thickspace negthinspace
quad qquad enspace`)

// latexSingles are the one-character commands: \\ \, \{ \% ...
const latexSingles = "\\ ,:>;!{}%$#&_|"

func toSet(words string) map[string]bool {
	set := map[string]bool{}
	for _, w := range strings.Fields(words) {
		set[w] = true
	}
	return set
}

// CheckLatex reports unknown commands and unbalanced braces in a
// formula, which the typesetter would otherwise show as raw text.
func CheckLatex(src string) error {
	depth := 0
	for i := 0; i < len(src); i++ {
		switch c := src[i]; {
		case c == '%':
			return fmt.Errorf("bare %% starts a comment; write \\%% in %q", src)
		case c == '{':
			depth++
		case c == '}':
			if depth--; depth < 0 {
				return fmt.Errorf("unmatched } in %q", src)
			}
		case c == '\\' && i+1 < len(src):
			j := i + 1
			for j < len(src) && isASCIILetter(src[j]) {
				j++
			}
			if j == i+1 {
				if !strings.ContainsRune(latexSingles, rune(src[j])) {
					return fmt.Errorf("unknown command \\%c in %q", src[j], src)
				}
				i++
				continue
			}
			if name := src[i+1 : j]; !latexCommands[name] {
				return fmt.Errorf("unknown command \\%s in %q", name, src)
			}
			i = j - 1
		}
	}
	if depth != 0 {
		return fmt.Errorf("unbalanced braces in %q", src)
	}
	return nil
}

func isASCIILetter(c byte) bool { return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }
