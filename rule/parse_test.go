package rule

import (
	"testing"

	"github.com/waflab/waflab/util"
)

func TestParseRule(t *testing.T) {
	rf := newRulefile(0, "REQUEST-920-PROTOCOL-ENFORCEMENT")
	text := util.ReadStringFromPath(util.CrsRuleDir + "REQUEST-920-PROTOCOL-ENFORCEMENT.conf")

	rf.loadRules(text)
	rf.syncParanoiaLevels()
	printRules(rf)
}

func TestParseRuleText(t *testing.T) {
	rf := newRulefile(0, "REQUEST-920-PROTOCOL-ENFORCEMENT")
	text := `
SecRule TX:EXECUTING_PARANOIA_LEVEL "@lt 1" "id:920011,phase:1,pass,nolog,skipAfter:END-REQUEST-920-PROTOCOL-ENFORCEMENT"
SecRule TX:EXECUTING_PARANOIA_LEVEL "@lt 1" "id:920012,phase:2,pass,nolog,skipAfter:END-REQUEST-920-PROTOCOL-ENFORCEMENT"
`

	rf.loadRules(text)
}

func TestParseAllRules(t *testing.T) {
	ReadRuleset("crs-3.2")
}
