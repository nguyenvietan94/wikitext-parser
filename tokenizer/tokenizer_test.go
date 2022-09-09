package tokenizer

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTokenizer(t *testing.T) {

	samples := map[string][]Token{
		`{{Thông tin tiểu sử bóng đá | caption = Messi trong màu áo [[đội tuyển bóng đá quốc gia Argentina|Argentina]] tại [[World Cup 2018]]}}`: {
			{"templateOpen", "{{"},
			{"text", "Thông tin tiểu sử bóng đá"},
			{"templateParamSeparator", "|"},
			{"text", "caption"},
			{"templateParamEquals", "="},
			{"text", "Messi trong màu áo"},
			{"wikilinkOpen", "[["},
			{"text", "đội tuyển bóng đá quốc gia Argentina"},
			{"templateParamSeparator", "|"},
			{"text", "Argentina"},
			{"wikilinkClose", "]]"},
			{"text", "tại"},
			{"wikilinkOpen", "[["},
			{"text", "World Cup 2018"},
			{"wikilinkClose", "]]"},
			{"templateClose", "}}"},
		},
		`{{Thông tin nhân vật
			| tên = Bradley Cooper
			| hình = Bradley Cooper avp 2014.jpg
			| ghi chú hình = Bradley tại buổi ra mắt ''[[American Hustle]]'' ở Pháp, tháng 2 năm 2014
			| cỡ hình = 200px
			| tên khai sinh = Bradley Charles Cooper
			| ngày sinh = {{birth date and age|1975|1|5}}
			| học vấn = [[Đại học Georgetown]]<br> [[The New School]]
			| nghề nghiệp = Diễn viên, nhà sản xuất
			| năm hoạt động = 1999 - nay
			| người hôn phối = {{marriage|[[Jennifer Esposito]]|2006|2007|reason=divorced}}
			| Bạn đời = 
			| nơi sinh = [[Philadelphia]], [[Pennsylvania]], [[Hoa Kỳ]]
			}}`: {
			{"templateOpen", "{{"},
			{"text", "Thông tin nhân vật"},
			{"templateParamSeparator", "|"},
			{"text", "tên"},
			{"templateParamEquals", "="},
			{"text", "Bradley Cooper"},
			{"templateParamSeparator", "|"},
			{"text", "hình"},
			{"templateParamEquals", "="},
			{"text", "Bradley Cooper avp 2014.jpg"},
			{"templateParamSeparator", "|"},
			{"text", "ghi chú hình"},
			{"templateParamEquals", "="},
			{"text", "Bradley tại buổi ra mắt"},
			{"italic", "''"},
			{"wikilinkOpen", "[["},
			{"text", "American Hustle"},
			{"wikilinkClose", "]]"},
			{"italic", "''"},
			{"text", "ở Pháp, tháng 2 năm 2014"},
			{"templateParamSeparator", "|"},
			{"text", "cỡ hình"},
			{"templateParamEquals", "="},
			{"text", "200px"},
			{"templateParamSeparator", "|"},
			{"text", "tên khai sinh"},
			{"templateParamEquals", "="},
			{"text", "Bradley Charles Cooper"},
			{"templateParamSeparator", "|"},
			{"text", "ngày sinh"},
			{"templateParamEquals", "="},
			{"templateOpen", "{{"},
			{"text", "birth date and age"},
			{"templateParamSeparator", "|"},
			{"text", "1975"},
			{"templateParamSeparator", "|"},
			{"text", "1"},
			{"templateParamSeparator", "|"},
			{"text", "5"},
			{"templateClose", "}}"},
			{"templateParamSeparator", "|"},
			{"text", "học vấn"},
			{"templateParamEquals", "="},
			{"wikilinkOpen", "[["},
			{"text", "Đại học Georgetown"},
			{"wikilinkClose", "]]"},
			{"break", "<br>"},
			{"wikilinkOpen", "[["},
			{"text", "The New School"},
			{"wikilinkClose", "]]"},
			{"templateParamSeparator", "|"},
			{"text", "nghề nghiệp"},
			{"templateParamEquals", "="},
			{"text", "Diễn viên, nhà sản xuất"},
			{"templateParamSeparator", "|"},
			{"text", "năm hoạt động"},
			{"templateParamEquals", "="},
			{"text", "1999 - nay"},
			{"templateParamSeparator", "|"},
			{"text", "người hôn phối"},
			{"templateParamEquals", "="},
			{"templateOpen", "{{"},
			{"text", "marriage"},
			{"templateParamSeparator", "|"},
			{"wikilinkOpen", "[["},
			{"text", "Jennifer Esposito"},
			{"wikilinkClose", "]]"},
			{"templateParamSeparator", "|"},
			{"text", "2006"},
			{"templateParamSeparator", "|"},
			{"text", "2007"},
			{"templateParamSeparator", "|"},
			{"text", "reason"},
			{"templateParamEquals", "="},
			{"text", "divorced"},
			{"templateClose", "}}"},
			{"templateParamSeparator", "|"},
			{"text", "Bạn đời"},
			{"templateParamEquals", "="},
			{"templateParamSeparator", "|"},
			{"text", "nơi sinh"},
			{"templateParamEquals", "="},
			{"wikilinkOpen", "[["},
			{"text", "Philadelphia"},
			{"wikilinkClose", "]]"},
			{"text", ","},
			{"wikilinkOpen", "[["},
			{"text", "Pennsylvania"},
			{"wikilinkClose", "]]"},
			{"text", ","},
			{"wikilinkOpen", "[["},
			{"text", "Hoa Kỳ"},
			{"wikilinkClose", "]]"},
			{"templateClose", "}}"},
		},
		`| height = {{height|m=1,69}}<ref>{{Cite web|title=Lionel Messi|url=https://en.psg.fr/teams/first-team/squad/lionel-messi|access-date=23 August 2021|website=PSG}}</ref>`: {
			{"templateParamSeparator", "|"},
			{"text", "height"},
			{"templateParamEquals", "="},
			{"templateOpen", "{{"},
			{"text", "height"},
			{"templateParamSeparator", "|"},
			{"text", "m"},
			{"templateParamEquals", "="},
			{"text", "1,69"},
			{"templateClose", "}}"},
			{"tagRefOpen", "<ref>"},
			{"templateOpen", "{{"},
			{"text", "Cite web"},
			{"templateParamSeparator", "|"},
			{"text", "title"},
			{"templateParamEquals", "="},
			{"text", "Lionel Messi"},
			{"templateParamSeparator", "|"},
			{"text", "url"},
			{"templateParamEquals", "="},
			{"text", "https://en.psg.fr/teams/first-team/squad/lionel-messi"},
			{"templateParamSeparator", "|"},
			{"text", "access-date"},
			{"templateParamEquals", "="},
			{"text", "23 August 2021"},
			{"templateParamSeparator", "|"},
			{"text", "website"},
			{"templateParamEquals", "="},
			{"text", "PSG"},
			{"templateClose", "}}"},
			{"tagRefClose", "</ref>"},
		},
	}

	for data, expected := range samples {
		tokens, err := Tokenize(data)
		assert.Equal(t, err, nil)
		var out []Token
		for _, token := range tokens {
			out = append(out, Token{token.Type, token.Token})
		}
		assert.Equal(t, out, expected)

		minLen := len(out)
		if len(expected) < minLen {
			minLen = len(expected)
		}
		for i := 0; i < minLen; i++ {
			assert.Equal(t, out[i], expected[i], fmt.Sprintf("line %d", i))
		}

	}
}

func TestSimple(t *testing.T) {
	// data := `{{Thông tin tiểu sử bóng đá
	// 	| fullname = Lionel Andrés Messi<ref>{{cite web |url=https://www.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |title=FIFA World Cup Russia 2018: List of Players: Argentina |work=FIFA |page=1 |date=15 July 2018 |archive-url=https://web.archive.org/web/20190611000407/https://www.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |archive-date=11 June 2019}}</ref>
	// 	| birth_date = {{birth date and age|1987|6|24|df=y}}<ref>{{cite web |url=https://tournament.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |title=2018 FIFA World Cup Russia: List of players: Argentina |work=FIFA |page=1 |date=15 July 2018 |access-date=13 October 2018 |archive-url=https://web.archive.org/web/20180619164139/https://tournament.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |archive-date=19 June 2018 |url-status=dead}}</ref>
	// 	}}`

	data := `{{Infobox person
				| genre              = {{flatlist|
				* [[Nhạc pop|Pop]]
				* [[Nhạc đồng quê|đồng quê]]
				* [[Alternative rock|alternative]]
				* [[Nhạc rock|rock]]
				* [[Dân gian đương đại|folk]]
				}}
				| instrument         = {{flatlist|
				* Giọng hát
				* guitar
				* banjo
				* piano
				* ukulele}}
			}}`
	tokens, err := Tokenize(data)
	assert.Equal(t, err, nil)
	for _, token := range tokens {
		fmt.Println(token.Token)
	}
}
