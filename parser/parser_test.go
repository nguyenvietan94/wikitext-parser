package parser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParser(t *testing.T) {
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
}

func TestCompanyTemplate(t *testing.T) {
	samples := map[string]map[string]string{
		`{{Hộp thông tin tóm tắt về công ty
			| name = Apple Inc.
			| logo = [[Tập tin:Apple logo black.svg||80px|Logo của Apple Inc.]]
			| image = Apple park cupertino 2019.jpg
			| image_size = 260px
			| image_caption = Trụ sở Apple ở [[Cupertino, California]]
			| er = [[Công ty đại chúng]]
			| traded_as = {{nasdaq|AAPL}}, {{lse|0HDZ}}, {{FWB|APC}}
			| predecesốsor = 
			| foundation = {{Start date and age|1976|04|01}} ([[Cupertino, California|Cupertino]], [[California]], [[Hoa Kỳ|Mỹ]])
			| founder = [[Steve Jobs]], [[Steve Wozniak]], [[Ronald Wayne]]<ref name=AppleConf>{{chú thích sách| last = Linzmayer| first = Ronald W.| title = Apple Confidential: The Real Story of Apple Computer, Inc.| publisher = No Starch Press| year = 1999| url = http://extras.denverpost.com/books/chap0411h.htm| access-date = ngày 1 tháng 6 năm 2018 | access-date = ngày 1 tháng 6 năm 2018}}</ref>
			| location_city = [[Cupertino, California]]
			| location_country = Mỹ
			| num_locations = 510 cửa hàng bán lẻ
			| num_locations_year = 2020
			| area_served = Toàn thế giới
			| key_people = [[Tim Cook]] ([[Tổng giám đốc điều hành|CEO]])
			[[Arthur D. Levinson]] ([[Chủ tịch hội đồng quản trị]])
			[[Jeff Williams]] ([[Giám đốc điều hành|COO]])
			| industry = [[Phần cứng|Phần cứng máy tính]]{{·}} [[Phần mềm|Phần mềm máy tính]], phụ kiện, thiết bị di động
			| products =
			{{Collapsible list
			 |framestyle=border:none; padding:0;
			 |title=
			 |1=<li>[[Macintosh|Mac]]
			 |2=<li>[[iPod]]
			 |3=<li>[[iPhone]]
			 |4=<li>[[iPad]]
			 |5=<li>[[Apple Watch]]
			 |6=<li>[[Apple TV]]
			 |7=<li>[[macOS]]
			 |8=<li>[[iOS(Apple)|iOS]] |9=<li>[[iPadOS]]
			 |10=<li>[[watchOS]]
			 |11=<li>[[tvOS]]
			}}
			| services =
			{{Collapsible list
			 |framestyle=border:none; padding:0;
			 |title=
			 |1=<li>[[Apple Arcade]]
			 |2=<li>[[Apple Card]]
			 |3=<li>[[Apple Music]]
			 |4=<li>[[Apple News+]]
			 |5=<li>[[Apple TV+]]
			 |6=<li>[[Apple Store (online)|Apple Store online]]
			 |7=<li>[[App Store (iOS)|App Store]]
			 |8=<li>[[iTunes|iTunes Store]]
			 |9=<li>[[Mac App Store]]
			 |10=<li>[[iBookstore|iBooks]]
			 |11=<li>[[iCloud]]
			 |12=<li>[[Apple Pay]]
			 |13=<li>[[iMessage]]
			 |14=<li>[[FaceTime]]
			}}
			| revenue = {{increase}} {{US$|274.515&nbsp;tỉ|link=yes}}<ref name="SEC filing">{{chú thích web|url=https://s2.q4cdn.com/470004039/files/doc_financials/2020/q4/FY20_Q4_Consolidated_Financial_Statements.pdf|date=30 tháng 10 năm 2020}}</ref>
			| revenue_year = 2020
			| operating_income = {{increase}} {{US$|66.288&nbsp;tỉ}}<ref name="SEC filing"/>
			| income_year = 2020
			| net_income = {{increase}} {{US$|57.411&nbsp;tỉ}}<ref name="SEC filing"/>
			| net_income_year = 2020
			| assets = {{decrease}} {{US$|323.888&nbsp;tỉ}}<ref name="SEC filing"/>
			| assets_year = 2020
			| equity = {{decrease}} {{US$|65.339&nbsp;tỉ}}<ref name="SEC filing"/>
			| equity_year = 2020
			| num_employees = 147,000<ref name="SEC filing"/>
			| num_employees_year = 2020
			| divisions = 
			| subsid = [[Shazam (company)|Shazam]], [[FileMaker Inc.]], [[Anobit]], [[Braeburn Capital]], [[Beats Electronics]]
			| website = {{URL|apple.com}}
			| intl = yes
			}}`: {
			"loại hình":   "Công ty đại chúng",
			"mã niêm yết": "NASDAQ:AAPL, LSE:0HDZ, FWB:APC",
			"ngành nghề": "	Phần cứng máy tính · Phần mềm máy tính, phụ kiện, thiết bị di động",
			"thành lập":            "1 tháng 4 năm 1976 (Cupertino, California, Mỹ)",
			"người sáng lập":       "Steve Jobs, Steve Wozniak, Ronald Wayne",
			"trụ sở chính":         "Cupertino, California, Mỹ",
			"số lượng trụ sở":      "510 cửa hàng bán lẻ (2020)",
			"khu vực hoạt động":    "Toàn thế giới",
			"thành viên chủ chốt":  "Tim Cook (CEO)·Arthur D. Levinson (Chủ tịch hội đồng quản trị)·Jeff Williams (COO)",
			"sản phẩm":             "Mac·iPod·iPhone·iPad·Apple Watch·Apple TV·macOS·iOS·iPadOS·watchOS·tvOS",
			"dịch vụ":              "Apple Arcade · Apple Card · Apple Music · Apple News+ · Apple TV+ · Apple Store online · App Store · iTunes Store · Mac App Store · iBooks · iCloud · Apple Pay · iMessag · FaceTime",
			"doanh thu":            "274.515 tỉ đô la Mỹ (2020)",
			"lợi nhuận kinh doanh": "66.288 tỉ đô la Mỹ (2020)",
			"lợi nhuận ròng":       "57.411 tỉ đô la Mỹ (2020)",
			"tổng tài sản":         "323.888 tỉ đô la Mỹ (2020)",
			"tổng vốn chủ sở hữu":  "65.339 tỉ đô la Mỹ (2020)",
			"số nhân viên":         "147,000 (2020)",
			"công ty con":          "Shazam·FileMaker Inc.·Anobit·Braeburn Capital·Beats Electronics",
			"website":              "apple.com",
		},
	}
	for data, expected := range samples {
		p := NewParser(data)
		err := p.parse()
		assert.Equal(t, err, nil)
		params := p.getParams()
		for key, val := range expected {
			assert.Equal(t, params[key], val)
		}
	}
}
