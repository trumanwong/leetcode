package main

import (
	"fmt"
	"testing"
	"truman.com/leetcode/1181.BeforeandAfterPuzzle/beforeAndAfterPuzzles"
)

func TestBeforeAndAfterPuzzles(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{"writing code", "code rocks"}, []string{"writing code rocks"}},
		{[]string{"mission statement", "a quick bite to eat", "a chip off the old block", "chocolate bar", "mission impossible", "a man on a mission", "block party", "eat my words", "bar of soap"},
			[]string{"a chip off the old block party", "a man on a mission impossible", "a man on a mission statement", "a quick bite to eat my words", "chocolate bar of soap"}},
			{[]string{"a", "b", "a"}, []string{"a"}},
			{[]string{"ni kntqfmv thyqxe dh xhnbd thyqxe s","s oqefp kntqfmv l ts nalc dbnt","l zmb ts thyqxe kxi dh ni ovdqvb dbnt s l j nqirao j l","z ovdqvb qoqhnxt kntqfmv xhnbd j l","zjje l s ysvpvc ysvpvc oqefp rhgfzeaz rhgfzeaz zjje nalc jymavm xhnbd j l l nqirao j xhnbd","ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z","qtp ovdqvb dh jymavm kntqfmv zjje qtp qoqhnxt zjje zjje qtp j rhgfzeaz ts z j","ysvpvc kxi","nalc zaowxcc kxi","j nqirao vhyntf j j j jymavm rhgfzeaz zjje ts qtp lirhaxd j","ysvpvc dbnt zbvntzo qtp trijn","z nalc qtp qtp ni ts","dh","l lirhaxd dh ysvpvc kntqfmv j zjje ysvpvc ysvpvc trijn s dh dh dbnt dh","vhyntf thyqxe ysvpvc trijn","zbvntzo vhyntf zaowxcc vhyntf ni z qoqhnxt j zmb l ni","kntqfmv ni ovdqvb ni zmb oqefp kntqfmv j l zjje zbvntzo nqirao nqirao s","qtp lirhaxd zbvntzo zjje thyqxe nalc kxi zjje dbnt l z j","z thyqxe oqefp trijn lirhaxd jymavm zjje","zaowxcc zaowxcc zaowxcc thyqxe ni","qoqhnxt lirhaxd zjje vhyntf s","trijn xhnbd l ni kxi j s j trijn ysvpvc nqirao ts kntqfmv qtp vhyntf trijn j dbnt ts nqirao","qtp z qoqhnxt oqefp ts oqefp kntqfmv thyqxe s","zmb ni s nqirao rhgfzeaz nqirao vhyntf rhgfzeaz ovdqvb dbnt vhyntf l l ts","nqirao thyqxe kxi thyqxe l zmb j nqirao zmb ysvpvc zbvntzo l nalc j ni z ts xhnbd z kntqfmv kxi","dbnt qoqhnxt zjje","trijn s","xhnbd ni qtp j qoqhnxt zjje zaowxcc ni ts dh xhnbd zjje z dh ts","vhyntf s ovdqvb oqefp z zbvntzo trijn qoqhnxt oqefp trijn kntqfmv qoqhnxt l zaowxcc ni","l vhyntf trijn dh lirhaxd zmb dbnt vhyntf qoqhnxt ni l j s"},
				[]string{"dbnt qoqhnxt zjje l s ysvpvc ysvpvc oqefp rhgfzeaz rhgfzeaz zjje nalc jymavm xhnbd j l l nqirao j xhnbd","kntqfmv ni ovdqvb ni zmb oqefp kntqfmv j l zjje zbvntzo nqirao nqirao s oqefp kntqfmv l ts nalc dbnt","l lirhaxd dh ysvpvc kntqfmv j zjje ysvpvc ysvpvc trijn s dh dh dbnt dh","l vhyntf trijn dh lirhaxd zmb dbnt vhyntf qoqhnxt ni l j s oqefp kntqfmv l ts nalc dbnt","l zmb ts thyqxe kxi dh ni ovdqvb dbnt s l j nqirao j l lirhaxd dh ysvpvc kntqfmv j zjje ysvpvc ysvpvc trijn s dh dh dbnt dh","l zmb ts thyqxe kxi dh ni ovdqvb dbnt s l j nqirao j l vhyntf trijn dh lirhaxd zmb dbnt vhyntf qoqhnxt ni l j s","ni kntqfmv thyqxe dh xhnbd thyqxe s oqefp kntqfmv l ts nalc dbnt","qoqhnxt lirhaxd zjje vhyntf s oqefp kntqfmv l ts nalc dbnt","qtp lirhaxd zbvntzo zjje thyqxe nalc kxi zjje dbnt l z j nqirao vhyntf j j j jymavm rhgfzeaz zjje ts qtp lirhaxd j","qtp ovdqvb dh jymavm kntqfmv zjje qtp qoqhnxt zjje zjje qtp j rhgfzeaz ts z j nqirao vhyntf j j j jymavm rhgfzeaz zjje ts qtp lirhaxd j","qtp z qoqhnxt oqefp ts oqefp kntqfmv thyqxe s oqefp kntqfmv l ts nalc dbnt","s oqefp kntqfmv l ts nalc dbnt qoqhnxt zjje","trijn s oqefp kntqfmv l ts nalc dbnt","trijn xhnbd l ni kxi j s j trijn ysvpvc nqirao ts kntqfmv qtp vhyntf trijn j dbnt ts nqirao thyqxe kxi thyqxe l zmb j nqirao zmb ysvpvc zbvntzo l nalc j ni z ts xhnbd z kntqfmv kxi","ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z nalc qtp qtp ni ts","ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z ovdqvb qoqhnxt kntqfmv xhnbd j l","ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z thyqxe oqefp trijn lirhaxd jymavm zjje","vhyntf s ovdqvb oqefp z zbvntzo trijn qoqhnxt oqefp trijn kntqfmv qoqhnxt l zaowxcc ni kntqfmv thyqxe dh xhnbd thyqxe s","vhyntf thyqxe ysvpvc trijn s","vhyntf thyqxe ysvpvc trijn xhnbd l ni kxi j s j trijn ysvpvc nqirao ts kntqfmv qtp vhyntf trijn j dbnt ts nqirao","xhnbd ni qtp j qoqhnxt zjje zaowxcc ni ts dh xhnbd zjje z dh ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z","ysvpvc dbnt zbvntzo qtp trijn s","ysvpvc dbnt zbvntzo qtp trijn xhnbd l ni kxi j s j trijn ysvpvc nqirao ts kntqfmv qtp vhyntf trijn j dbnt ts nqirao","z nalc qtp qtp ni ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z","z ovdqvb qoqhnxt kntqfmv xhnbd j l lirhaxd dh ysvpvc kntqfmv j zjje ysvpvc ysvpvc trijn s dh dh dbnt dh","z ovdqvb qoqhnxt kntqfmv xhnbd j l vhyntf trijn dh lirhaxd zmb dbnt vhyntf qoqhnxt ni l j s","z ovdqvb qoqhnxt kntqfmv xhnbd j l zmb ts thyqxe kxi dh ni ovdqvb dbnt s l j nqirao j l","z thyqxe oqefp trijn lirhaxd jymavm zjje l s ysvpvc ysvpvc oqefp rhgfzeaz rhgfzeaz zjje nalc jymavm xhnbd j l l nqirao j xhnbd","zaowxcc zaowxcc zaowxcc thyqxe ni kntqfmv thyqxe dh xhnbd thyqxe s","zbvntzo vhyntf zaowxcc vhyntf ni z qoqhnxt j zmb l ni kntqfmv thyqxe dh xhnbd thyqxe s","zjje l s ysvpvc ysvpvc oqefp rhgfzeaz rhgfzeaz zjje nalc jymavm xhnbd j l l nqirao j xhnbd ni qtp j qoqhnxt zjje zaowxcc ni ts dh xhnbd zjje z dh ts","zmb ni s nqirao rhgfzeaz nqirao vhyntf rhgfzeaz ovdqvb dbnt vhyntf l l ts thyqxe nalc xhnbd ysvpvc l trijn qoqhnxt zaowxcc qtp lirhaxd j xhnbd zbvntzo lirhaxd vhyntf z"}},
	}

	for _, test := range tests {
		ret := beforeAndAfterPuzzles.BeforeAndAfterPuzzles(test.input)
		fmt.Println(len(ret) == len(test.output))
		for i, v := range ret {
			if v != test.output[i] {
				fmt.Println(i)
				fmt.Println(v)
				fmt.Println(test.output[i])
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				//break
			}
		}
	}
}