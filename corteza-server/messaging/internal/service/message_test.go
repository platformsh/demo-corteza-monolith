package service

import (
	"strings"
	"testing"

	"github.com/cortezaproject/corteza-server/internal/test"
	"github.com/cortezaproject/corteza-server/messaging/types"
)

func TestMessageLength(t *testing.T) {
	svc := message{}
	e := func(out *types.Message, err error) error { return err }

	test.Assert(t, e(svc.Create(&types.Message{})) != nil, "Should not allow to create empty message")

	if settingsMessageBodyLength > 0 {
		longText := strings.Repeat("X", settingsMessageBodyLength+1)
		test.Assert(t, e(svc.Create(&types.Message{Message: longText})) != nil, "Should not allow to create message with really long text")
	}
}

func TestMentionsExtraction(t *testing.T) {
	var (
		svc   = message{}
		mm    types.MentionSet
		cases = []struct {
			text string
			ids  []uint64
		}{
			{"abcde",
				[]uint64{}},
			{"<@4095834095>",
				[]uint64{4095834095}},
			{"<@4095834095> <@4095834095>",
				[]uint64{4095834095}},
			{"<@4095834095> <@4095834097>",
				[]uint64{4095834095, 4095834097}},
			{"dfsf<@4095834095>dsfsd<@4095834097>sdfs",
				[]uint64{4095834095, 4095834097}},
			{"dfsf<@4095834095>dsfsd<@40958340dfsZ",
				[]uint64{4095834095}},
			{"<@4095834095 label> <@4095834097>",
				[]uint64{4095834095, 4095834097}},
		}
	)

	for _, c := range cases {
		mm = svc.extractMentions(&types.Message{Message: c.text})

		test.Assert(t, len(mm) == len(c.ids), "Number of extracted (%d) and expected (%d) user IDs do not match (%s)", len(mm), len(c.ids), c.text)

		for _, id := range c.ids {
			test.Assert(t, len(mm.FindByUserID(id)) == 1, "Owner ID (%d) was not extracted (%s)", id, c.text)
		}
	}
}
