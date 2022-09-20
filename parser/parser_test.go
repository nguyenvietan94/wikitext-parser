package parser

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

var (
	AppleInfobox string = `{{Hộp thông tin tóm tắt về công ty
		| name = Apple Inc.
		| logo = [[Tập tin:Apple logo black.svg||80px|Logo của Apple Inc.]]
		| image = Apple park cupertino 2019.jpg
		| image_size = 260px
		| image_caption = Trụ sở Apple ở [[Cupertino, California]]
		| type = [[Công ty đại chúng]]
		| traded_as = {{nasdaq|AAPL}}, {{lse|0HDZ}}, {{FWB|APC}}
		| predecessor = 
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
		}}`
	Amazon string = `{{Infobox company
			| name  = Amazon.com, Inc.
			| logo  = Amazon logo.svg
			| logo_size = 175px
			| image = Seattle Spheres on May 10, 2018.jpg
			| image_size = 250px
			| image_caption = [[Amazon Spheres]], một phần của trụ sở chính Amazon ở [[Seattle]]
			| type  = [[Công ty đại chúng|Đại chúng]]
			| traded_as = {{NASDAQ|AMZN}}<br/>[[NASDAQ-100|NASDAQ-100 Component]]<br/>[[S&P 100|S&P 100 Component]]<br/>[[S&P 500|S&P 500 Component]]
			| founded  = {{Start date and age|1994|07|05}} tại [[Bellevue, Washington]], [[Hoa Kỳ]]
			| founder  = [[Jeff Bezos]]
			| hq_location = [[Seattle|Seattle, Washington]] 
			| area_served = Toàn cầu
			| key_people = [[Andy Jassy]]<small>([[Giám đốc điều hành|CEO]])</small> <br />[[Werner Vogels]]<small>([[Giám đốc công nghệ|CTO]])<small>
			| num_employees = 798,000 (2019)
			| industry  = {{plainlist|
			* [[Điện toán đám mây]]
			*[[Trí tuệ nhân tạo]]
			*[[Hệ thống mạng]]
			*[[Phân phối kỹ thuật số]]
			* [[Thương mại điện tử]]
			* [[Điện tử gia dụng]]
			}}
			| revenue  = {{increase}} 280,522 tỷ [[Đô la Mỹ|USD]] (2019)
			| operating_income = {{increase}} 14,541 tỷ USD (2019)
			| net_income = {{increase}} 11,588 tỷ USD (2019)
			| assets = {{increase}} 255,248 tỷ USD (2019)
			| equity = {{increase}} 62,060 tỷ USD (2019)
			| products  = {{Hlist|[[Amazon Echo]]|[[Amazon Kindle]]|[[Amazon Fire]]|[[Amazon Fire Tv]]|[[Amazon Fire OS]]}}
			| website_type = [[thương mại điện tử]]
			| language  = tiếng Anh, tiếng Nhật, tiếng Đức, tiếng Pháp, tiếng Trung
			| advertising = [[web banner]] và [[video]]
			| launch_date = 1995
			| alexa = 5<ref>{{chú thích web| url=http://www.alexa.com/siteinfo/amazon.com| title=amazon.com - Traffic Details from Alexa| publisher=[[Alexa Internet]], Inc| access-date=ngày 24 tháng 8 năm 2010| archive-date=2016-04-03| archive-url=https://web.archive.org/web/20160403034253/http://www.alexa.com/siteinfo/amazon.com| url-status=dead}}</ref>
			| website  = {{URL|https://www.amazon.com}}
			|company_name=|công ty con=|production=|dịch vụ={{Hlist|[[Amazon (company)#Website|Amazon.com]]|[[Amazon Alexa]]|[[Amazon Appstore]]|[[Amazon Music]]|[[Amazon Prime]]|[[Amazon Prime Video]]|[[Amazon Web Services]]}}}}`
)

func TestExtractingInfoboxAsPlainText(t *testing.T) {
	samples := map[string]map[string]string{
		AppleInfobox: {
			"name":               "Apple Inc.",
			"logo":               "Apple logo black.svg",
			"image":              "Apple park cupertino 2019.jpg",
			"image_size":         "260px",
			"image_caption":      "Trụ sở Apple ở Cupertino, California",
			"type":               "Công ty đại chúng",
			"traded_as":          "NASDAQ:AAPL, LSE:0HDZ, FWB:APC",
			"foundation":         "01-04-1976 (Cupertino, California, Mỹ)",
			"founder":            "Steve Jobs, Steve Wozniak, Ronald Wayne",
			"location_city":      "Cupertino, California",
			"location_country":   "Mỹ",
			"num_locations":      "510 cửa hàng bán lẻ",
			"num_locations_year": "2020",
			"area_served":        "Toàn thế giới",
			"key_people":         "Tim Cook (CEO) Arthur D. Levinson (Chủ tịch hội đồng quản trị) Jeff Williams (COO)",
			"industry":           "Phần cứng máy tính · Phần mềm máy tính, phụ kiện, thiết bị di động",
			"products":           "Mac, iPod, iPhone, iPad, Apple Watch, Apple TV, macOS, iOS, iPadOS, watchOS, tvOS",
			"services":           "Apple Arcade, Apple Card, Apple Music, Apple News+, Apple TV+, Apple Store online, App Store, iTunes Store, Mac App Store, iBooks, iCloud, Apple Pay, iMessage, FaceTime",
			"revenue":            "274.515 tỉ đô la Mỹ",
			"revenue_year":       "2020",
			"operating_income":   "66.288 tỉ đô la Mỹ",
			"income_year":        "2020",
			"net_income":         "57.411 tỉ đô la Mỹ",
			"net_income_year":    "2020",
			"assets":             "323.888 tỉ đô la Mỹ",
			"assets_year":        "2020",
			"equity":             "65.339 tỉ đô la Mỹ",
			"equity_year":        "2020",
			"num_employees":      "147,000",
			"num_employees_year": "2020",
			"subsid":             "Shazam, FileMaker Inc., Anobit, Braeburn Capital, Beats Electronics",
			"website":            "apple.com",
			"intl":               "yes",
		},
		Amazon: {},
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

func TestExtractingRequiredFieldsInTemplate(t *testing.T) {
	samples := map[string]map[string]string{
		AppleInfobox: {
			"tên công ty":          "Apple Inc.",
			"logo":                 "Apple logo black.svg",
			"image":                "Apple park cupertino 2019.jpg",
			"loại hình":            "Công ty đại chúng",
			"mã niêm yết":          "NASDAQ:AAPL, LSE:0HDZ, FWB:APC",
			"ngành nghề":           "Phần cứng máy tính · Phần mềm máy tính, phụ kiện, thiết bị di động",
			"thành lập":            "01-04-1976 (Cupertino, California, Mỹ)",
			"người sáng lập":       "Steve Jobs, Steve Wozniak, Ronald Wayne",
			"trụ sở chính":         "Cupertino, California, Mỹ",
			"số lượng trụ sở":      "510 cửa hàng bán lẻ(2020)",
			"khu vực hoạt động":    "Toàn thế giới",
			"thành viên chủ chốt":  "Tim Cook (CEO) Arthur D. Levinson (Chủ tịch hội đồng quản trị) Jeff Williams (COO)",
			"sản phẩm":             "Mac, iPod, iPhone, iPad, Apple Watch, Apple TV, macOS, iOS, iPadOS, watchOS, tvOS",
			"dịch vụ":              "Apple Arcade, Apple Card, Apple Music, Apple News+, Apple TV+, Apple Store online, App Store, iTunes Store, Mac App Store, iBooks, iCloud, Apple Pay, iMessage, FaceTime",
			"doanh thu":            "274.515 tỉ đô la Mỹ(2020)",
			"lợi nhuận kinh doanh": "66.288 tỉ đô la Mỹ(2020)",
			"lợi nhuận ròng":       "57.411 tỉ đô la Mỹ(2020)",
			"tổng tài sản":         "323.888 tỉ đô la Mỹ(2020)",
			"tổng vốn chủ sở hữu":  "65.339 tỉ đô la Mỹ(2020)",
			"số nhân viên":         "147,000(2020)",
			"công ty con":          "Shazam, FileMaker Inc., Anobit, Braeburn Capital, Beats Electronics",
			"website":              "apple.com",
		},
		Amazon: {},
	}
	for data, expected := range samples {
		p := NewParser(data)
		err := p.parse()
		assert.Equal(t, err, nil)
		fields, err := p.getRequiredFields()
		assert.Equal(t, err, nil)
		for key, val := range expected {
			assert.Equal(t, fields[key] != nil, true, fmt.Sprintf("params[%s] is nil", key))
			if fields[key] != nil {
				assert.Equal(t, fields[key].value, val)
			}
		}
	}
}
