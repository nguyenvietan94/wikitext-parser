package parser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParser(t *testing.T) {
	// data := `{{Thông tin tiểu sử bóng đá
	// 	| caption = Messi trong màu áo [[đội tuyển bóng đá quốc gia Argentina|Argentina]] tại [[World Cup 2018]]
	// 	| fullname = Lionel Andrés Messi<ref>{{cite web |url=https://www.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |title=FIFA World Cup Russia 2018: List of Players: Argentina |work=FIFA |page=1 |date=15 July 2018 |archive-url=https://web.archive.org/web/20190611000407/https://www.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |archive-date=11 June 2019}}</ref>
	// 	| birth_date = {{birth date and age|1987|6|24|df=y}}<ref>{{cite web |url=https://tournament.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |title=2018 FIFA World Cup Russia: List of players: Argentina |work=FIFA |page=1 |date=15 July 2018 |access-date=13 October 2018 |archive-url=https://web.archive.org/web/20180619164139/https://tournament.fifadata.com/documents/FWC/2018/pdf/FWC_2018_SQUADLISTS.PDF |archive-date=19 June 2018 |url-status=dead}}</ref>
	// 	| height = {{height|m=1,69}}<ref>{{Cite web|title=Lionel Messi|url=https://en.psg.fr/teams/first-team/squad/lionel-messi|access-date=23 August 2021|website=PSG}}</ref>
	// 	| birth_place = [[Rosario]], [[Santa Fe]], Argentina
	// 	}}`

	// data := `{{Hộp thông tin tóm tắt về công ty|name=Binance|logo=Binance logo.svg|foundation=2017|founders={{unbulleted list|[[Changpeng Zhao]]|Yi He}}|products=[[Cryptocurrency exchange]]|location_country=[[Malta]]<ref>{{chú thích web |title=Why world leader crypto exchange Binance moved to Malta |url=https://www.maltatoday.com.mt/business/business_news/93170/why_world_leader_crypto_exchange_binance_moved_to_malta#.XTGugOhKhhE |website=[[Malta Today]] |access-date =ngày 19 tháng 7 năm 2019}}</ref>|key_people=Changpeng Zhao ([[CEO]])|homepage={{URL|www.binance.com}}}}`

	data := `{{Infobox person
			| name = Taylor Swift
			| image = 191125 Taylor Swift at the 2019 American Music Awards (cropped).png
			| caption = Swift tại lễ trao [[giải thưởng Âm nhạc Mỹ 2019]]
			| birth_name = Taylor Alison Swift
			| birth_date = {{Birth date and age|1989|12|13}}
			| birth_place = [[West Reading, Pennsylvania]], Hoa Kỳ
			| partner = [[Joe Alwyn]] (2016–nay)<ref>Sources on recognition of Alwyn as Swift's "partner"
			*{{Chú thích web |last1=Malone Kircher |first1=Madison |title=Thank You, Taylor Swift, for Making My Sad Yuletide Gay(er) |url=https://www.vulture.com/2020/12/taylor-swift-dorothea-evermore-queer-song-lyrics-analysis.html |website=Vulture |date=December 11, 2020 |access-date=15 September 2021 |quote=For "betty", it was "William Bowery" a.k.a. her partner Joe Alwyn who, according to Swift’s folklore concert film, plays the piano "just beautifully" and roams the house singing fully formed songs.}}
			*{{Chú thích web |title=Netflix's Miss Americana Unveils a New Taylor Swift |url=https://slate.com/culture/2020/01/taylor-swift-documentary-miss-americana-sundance-review-politics-trump.html |website=Slate |date=January 24, 2020 |access-date=15 September 2021 |quote=Although Swift credits much of her recent evolution to her long-term relationship with Joe Alwyn...}}
			*{{cite magazine |title=Paul McCartney & Taylor Swift |url=https://www.rollingstone.com/music/music-features/paul-mccartney-taylor-swift-musicians-on-musicians-1089058/ |magazine=Rolling Stone |date=November 13, 2020 |access-date=15 September 2021 |quote=McCartney: So how does that go? Does your partner sympathize with that and understand?  Swift: Oh, absolutely.}}</ref>
			| relatives = {{plainlist|
			* [[Austin Swift]] (em trai)
			* [[Marjorie Finlay]] (bà ngoại)
			}}
			| occupation = {{flatlist|
			* Ca sĩ kiêm sáng tác nhạc
			* {{nowrap|Nhà sản xuất thu âm}}
			* Diễn viên
			* Đạo diễn phim 
			* Doanh nhân
			}}
			| other_names = Nils Sjöberg<ref>{{cite magazine|first=Gil|last=Kaufman|url=https://www.billboard.com/articles/news/dance/7438158/taylor-swift-calvin-harris-co-wrote-this-is-what-you-came-for-pseudonym|title=Taylor Swift Co-Wrote Calvin Harris' 'This Is What You Came For' Under Swedish Pseudonym|magazine=[[Billboard (magazine)|Billboard]]|date=July 13, 2016|access-date=May 25, 2020}}</ref>
			| years_active = 2004–nay
			| awards = [[Danh sách giải thưởng và đề cử của Taylor Swift|Danh sách]]
			| website = {{URL|taylorswift.com}}
			| module = {{Infobox musical artist|embed=yes
			| origin             = [[Nashville, Tennessee]], Hoa Kỳ
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
			* ukulele
			}}
			| label              = {{flatlist|
			* [[Republic Records|Republic]]
			* [[Big Machine Records|Big Machine]]
			}}
			| associated_acts    = {{flatlist|
			* [[Ed Sheeran]]
			* [[Justin Vernon]]
			* [[Haim (ban nhạc)|Haim]]}}
			}}
			| signature = Taylor Swift Signature.png
			| học vị = [[Tiến sĩ danh dự]] [[ngành]] [[Mỹ thuật]] [[Đại học New York]]
			| chiều cao = 1,78 [[m]]
			}}`

	p := NewParser(data)
	err := p.parse()
	assert.Equal(t, err, nil)
}

func TestSimple(t *testing.T) {
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
	p := NewParser(data)
	err := p.parse()
	assert.Equal(t, err, nil)
	p.printTokens()
}
