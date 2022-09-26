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
	AmazonInfobox string = `{{Infobox company
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
		| key_people = [[Andy Jassy]]<small>([[Giám đốc điều hành|CEO]])</small><br />[[Werner Vogels]]<small>([[Giám đốc công nghệ|CTO]])</small>
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
	MicrosoftInfobox string = `{{Tóm tắt về công ty
		|tên = Microsoft Corporation
		|logo = Microsoft logo and wordmark.svg
		|logo_caption = Logo Microsoft từ năm 2012{{ndash}}nay
		|image = Building92microsoft.jpg
		|image_caption = Tòa nhà 92 trong khuôn viên Microsoft Redmond ở [[Redmond, Washington]]
		|loại = [[Công ty cổ phần]] ({{nasdaq|MSFT}})<br />[[Chỉ số trung bình công nghiệp Dow Jones|Dow Jones Industrial Average Component]]<br />[[S&P 500|S&P 500 Component]]
		|foundation = [[Albuquerque, New Mexico]]<br />4 tháng 4 năm 1975
		|founder = [[Bill Gates]]<br />[[Paul Allen]]
		|location = [[Microsoft Redmond Campus]]<br />[[Redmond, Washington]], Hoa Kỳ
		|area_served = Toàn cầu
		|key_people = [[Satya Nadella]] <small>([[Tổng giám đốc điều hành|CEO]], chủ tịch)</small> <br />[[Bill Gates]] <small>(Nhà sáng lập)</small>
		|industry = [[Phần mềm|Phần mềm máy tính]]<br />[[Phần cứng|Phần cứng máy tính]]<br />[[Điện thoại di động]]<br />[[Thiết bị viễn thông]]<br />[[Phân phối kỹ thuật số]]<br />[[Điện tử tiêu dùng]]<br />[[Trò chơi điện tử]]<br />[[Tư vấn công nghệ thông tin]]<br />[[Quảng cáo trực tuyến]]<br />[[Bán lẻ]]<br />[[Phần mềm ô-tô]]
		|products = [[Microsoft#Các sản phẩm|Xem Các sản phẩm]]
		|services = [[Microsoft#Các sản phẩm|Xem Các dịch vụ]]
		|market cap = $169,72 tỉ <small>(ngày 13 tháng 2 năm 2009)</small><ref>{{chú thích web |url=http://www.google.com/finance?q=MSFT |title=Company Profile for Microsoft Corp (MSFT) |access-date=ngày 13 tháng 2 năm 2009}}</ref>
		| revenue = {{Increase}} {{US$|110,36&nbsp;tỉ|link=yes}}<ref name="ER-FY18">{{Chú thích web |url=https://www.microsoft.com/en-us/Investor/earnings/FY-2018-Q4/press-release-webcast |title=Earnings Release FY18 Q4 |date= 19 tháng 7, 2018 |website=Microsoft |access-date=27 tháng 2, 2019}}</ref>
		| revenue_year = 2018
		| operating_income = {{Increase}} {{US$|35,05&nbsp;tỉ}}<ref name="ER-FY18" />
		| income_year = 2018
		| net_income = {{Increase}} {{US$|16,57&nbsp;tỉ}}<ref name="ER-FY18" />
		| net_income_year = 2018
		| assets = {{Increase}} {{US$|254.84&nbsp;tỉ}}<ref name="ER-FY18" />
		| assets_year = 2018
		| equity = {{Increase}} {{US$|82,71&nbsp;tỉ}}<ref name="ER-FY18" />
		| equity_year = 2018
		| num_employees = 138,944<ref name="ER-FY18" />
		| num_employees_year = 2018
		| website = {{URL|microsoft.com}}
		|image_size=300}}`
	MetaInfobox string = `{{Infobox company
		| name = Meta Platforms, Inc.
		| logo = Meta Inc. logo.svg
		| type = [[Công ty đại chúng|Đại chúng]]
		| image = File:Meta Platforms Headquarters Menlo Park California.jpg
		| image_caption = Toàn cảnh trụ sở Meta tại Menlo Park
		| traded_as = {{ubl|Lớp A: {{NASDAQ|FB}}|[[NASDAQ-100|NASDAQ-100 component]]|[[S&P 100|S&P 100 component]]|[[S&P 500|S&P 500 component]]|Lớp B: Không niêm yết}}
		| industry = {{ubl|[[Dịch vụ mạng xã hội]]|[[Quảng cáo]]}}
		| founded = {{Start date and age|2004|02|04}} tại [[Cambridge, Massachusetts]]
		| founders = {{plainlist|
		* [[Mark Zuckerberg]]
		* [[Eduardo Saverin]]
		* [[Andrew McCollum]]
		* [[Dustin Moskovitz]]
		* [[Chris Hughes]]
		}}
		| hq_location_city = [[Menlo Park, California]]
		| hq_location_country = U.S.
		| area_served = Toàn thế giới (trừ các quốc gia kiểm duyệt)
		| key_people = {{plainlist|
		* [[Mark Zuckerberg]] ([[Chủ tịch]], [[Tổng giám đốc điều hành|CEO]], [[Controlling interest|Controlling Shareholder]])
		* [[Sheryl Sandberg]] ([[Giám đốc điều hành|COO]])
		* [[David Wehner]] ([[Giám đốc tài chính|CFO]])
		* [[Mike Schroepfer]] ([[Giám đốc công nghệ|CTO]])
		}}
		| products = {{Plainlist|
		* [[Facebook]]
		* [[Instagram]]
		* [[Facebook Messenger|Messenger]]
		* [[WhatsApp]]
		* [[Facebook Watch|Watch]]
		* [[Facebook Portal|Portal]]
		* [[Oculus VR|Oculus]]
		* [[Calibra (công ty)|Calibra]]
		}}
		| revenue = {{increase}} 55,838 tỷ [[Đô la Mỹ|Mỹ kim]]
		| revenue_year = 2018
		| operating_income = {{increase}} 24,913 tỷ Mỹ kim
		| income_year = 2018
		| net_income = {{increase}} 22,111 tỷ Mỹ kim
		| net_income_year = 2018
		| assets = {{increase}} 97,334 tỷ Mỹ kim
		| assets_year = 2018
		| equity = {{increase}} 84,127 tỷ Mỹ kim
		| equity_year = 2018
		| num_employees = 43.030<ref>https://s21.q4cdn.com/399680738/files/doc_financials/2019/q3/FB-Q3-2019-Earnings-Release.pdf</ref>
		| num_employees_year = ngày 30 tháng 9 năm 2019
		| website = {{URL|https://about.fb.com/}}
		| footnotes = <ref name="Our History">{{chú thích web|title=Our History|url=https://newsroom.fb.com/company-info/|website=Facebook|access-date =ngày 7 tháng 11 năm 2018}}</ref><ref>{{chú thích web|last=Shaban|first=Hamza|title=Digital advertising to surpass print and TV for the first time, report says|url=https://www.washingtonpost.com/technology/2019/02/20/digital-advertising-surpass-print-tv-first-time-report-says|website=The Washington Post|access-date=ngày 2 tháng 6 năm 2019|date=ngày 20 tháng 2 năm 2019}}</ref><ref name=SOI>{{chú thích web|title=FB Income Statement|url=https://www.nasdaq.com/symbol/fb/financials|website=NASDAQ.com}}</ref><ref name=DOI>{{chú thích web|title=FB Balance Sheet|url=https://www.nasdaq.com/symbol/fb/financials?query=balance-sheet|website=NASDAQ.com}}</ref><ref>{{chú thích web|url=https://newsroom.fb.com/company-info/|title=Stats|publisher=Facebook|access-date=ngày 25 tháng 7 năm 2019|date=ngày 30 tháng 6 năm 2019}}</ref>
		}}`
	NetflixInfobox string = `{{Infobox company
		| name = Netflix, Inc.
		| logo = [[Tập tin:Netflix 2015 logo.svg|frameless|upright]]
		| screenshot= [[Tập tin:netflixvn.jpeg]]
		| caption = [[Ảnh chụp màn hình]] trang chủ Netflix phiên bản tiếng Việt
		| company_type = [[Công ty đại chúng]]
		| traded_as = {{ubl|{{NASDAQ|NFLX}}|[[NASDAQ-100]]|[[S&P 100]]|[[S&P 500]]}}
		| area_served = Toàn cầu, ngoại trừ [[Trung Quốc đại lục]], [[Syria]], [[Bắc Triều Tiên]] và [[Bán đảo Krym]]<ref>{{chú thích web|title=Where is Netflix available?|url=https://help.netflix.com/en/node/14164|publisher=Netflix|access-date =ngày 8 tháng 8 năm 2017|url-status=live|archiveurl=https://web.archive.org/web/20170707184542/https://help.netflix.com/en/node/14164|archivedate=ngày 7 tháng 7 năm 2017|df=mdy-all}}</ref>
		| founder = {{plainlist|
		* [[Reed Hastings]]
		* [[Marc Randolph]]
		}}
		| industry = Giải trí, [[truyền thông đại chúng]]
		| products = {{flatlist|
		* Streaming media
		* video on demand
		}}
		| services = {{flatlist|
		* Nhà sản xuất phim
		* Nhà phân phối phim
		* Chương trình truyền hình
		}}
		| divisions = US Streaming<br />International Streaming<br />Domestic DVD<ref>{{chú thích web|url= http://revenuesandprofits.com/how-netflix-makes-money/|title= How Netflix Makes Money? – Revenues &|last= Miglani|date= ngày 18 tháng 6 năm 2015|url-status=live|archiveurl= https://web.archive.org/web/20170219172342/https://revenuesandprofits.com/how-netflix-makes-money/|archivedate= ngày 19 tháng 2 năm 2017|df= mdy-all}}</ref>
		| revenue = {{increase}}{{US$|15.794 tỉ|link=yes}} (2018)<ref name="ir.netflix.com">{{chú thích web|url=https://www.netflixinvestor.com/financials/financial-statements/default.aspx|title=Netflix Q4 2018 Results|access-date = ngày 18 tháng 1 năm 2019}}</ref>
		| operating_income = {{increase}}US$1.605 tỉ (2018)<ref name="ir.netflix.com"/>
		| net_income = {{increase}}US$1.211 tỉ (2018)<ref name="ir.netflix.com"/>
		| assets = {{increase}}US$25.974 tỉ (2018)<ref name="ir.netflix.com"/>
		| equity = {{increase}}US$5.289 tỉ (2018)<ref name="ir.netflix.com"/>
		| alexa = {{increase}} 24 {{small|{{nowrap|({{as of|2019|04|06|alt=Tháng tư 2019}})}}}}<ref>{{chú thích web |title=Netflix.com Traffic, Demographics and Competitors - Alexa |url=https://www.alexa.com/siteinfo/netflix.com |website=www.alexa.com |access-date=ngày 6 tháng 4 năm 2019 |archive-date=2017-10-29 |archive-url=https://web.archive.org/web/20171029142606/https://www.alexa.com/siteinfo/netflix.com |url-status=dead }}</ref>
		| registration = Cần thiết
		| foundation = {{Start date and age|1997|8|29}}<ref>{{chú thích web|title=Business Search – Business Entities – Business Programs {{!}} California Secretary of State|url=https://businesssearch.sos.ca.gov/CBS/SearchResults?SearchType=CORP&SearchCriteria=Netflix&SearchSubType=Keyword|website=businesssearch.sos.ca.gov|access-date =ngày 26 tháng 5 năm 2017|url-status=live|archiveurl=https://web.archive.org/web/20170813103404/https://businesssearch.sos.ca.gov/CBS/SearchResults?SearchType=CORP&SearchCriteria=Netflix&SearchSubType=Keyword|archivedate=ngày 13 tháng 8 năm 2017|df=mdy-all}}</ref> tại [[Scotts Valley, California]]
		| location = [[Los Gatos, California]], Hoa Kỳ
		| key_people = {{plainlist|
		* Reed Hastings {{small|(Chủ tịch, Giám đốc điều hành)}}
		* [[Ted Sarandos]] {{small|(Giám đốc thương mại)}}
		}}
		| subsid = {{Plainlist|
		* DVD.com
		* [[Millarworld]]<ref name="Archived copy">{{chú thích web |url=https://ir.netflix.com/node/29631/html |title=Netflix - Financials - SEC Filings |access-date = ngày 30 tháng 1 năm 2018 |url-status=live |archiveurl=https://web.archive.org/web/20180131024137/https://ir.netflix.com/node/29631/html |archivedate=ngày 31 tháng 1 năm 2018 |df=mdy-all }}</ref>
		* LT-LA<ref>{{chú thích web |last1=Hipes |first1=Patrick |title=Netflix Takes Top Awards Strategist Lisa Taback Off The Table |url=https://deadline.com/2018/07/netflix-lisa-taback-hollywood-awards-strategist-hire-1202428876/ |website=deadline.com |access-date =ngày 18 tháng 7 năm 2018|date = ngày 18 tháng 7 năm 2018}}</ref>
		*ABQ Studios
		*Netflix Animation
		*Netflix Pte. Ltd.
		*Netflix Services UK Limited
		*Netflix Streaming Services International B.V.
		*Netflix Streaming Services, Inc.
		*Netflix Global, LLC
		*Netflix Services Germany GmbH
		*NetflixCS, Inc.
		*Netflix Luxembourg S.a r.l.
		*Netflix Studios
		*Netflix Entretenimento Brasil LTDA.
		*[[StoryBots]]
		}}
		| num_users = 154 triệu người (toàn cầu),<br />
		148 triệu người (đã thanh toán)<ref name="NFLX Q1 2019" />
		| num_employees = {{increase}} 5,400 (2017)<ref name="Archived copy"/>
		| website = {{URL|https://www.netflix.com}}
		}}`
	TPBankInfobox string = `{{Tóm tắt về công ty
		| biểu trưng      = Tập tin:Logo TPBank.svg
		| loại            = Tài chính
		| thành lập       = 5/5/2008
		| thành viên chủ chốt = Đỗ Minh Phú (Chủ tịch HĐQT)
		| sản phẩm        = Dịch vụ tài chính, Ngân hàng
		| tài sản         = 164,593 tỷ VNĐ (30/09/2019)
		| khẩu hiệu       = Vì chúng tôi hiểu bạn
		| trang chủ       = [http://tpb.vn/ https://tpb.vn/]
		|Vốn điều lệ=8,565 tỷ VNĐ (30/09/2019)}}`
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
			"foundation":         "01-04-1976(Cupertino, California, Mỹ)",
			"founder":            "Steve Jobs, Steve Wozniak, Ronald Wayne",
			"location_city":      "Cupertino, California",
			"location_country":   "Mỹ",
			"num_locations":      "510 cửa hàng bán lẻ",
			"num_locations_year": "2020",
			"area_served":        "Toàn thế giới",
			"key_people":         "Tim Cook(CEO) Arthur D. Levinson(Chủ tịch hội đồng quản trị) Jeff Williams(COO)",
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
		AmazonInfobox: {
			"name":             "Amazon.com, Inc.",
			"logo":             "Amazon logo.svg",
			"logo_size":        "175px",
			"image":            "Seattle Spheres on May 10, 2018.jpg",
			"image_size":       "250px",
			"image_caption":    "Amazon Spheres, một phần của trụ sở chính Amazon ở Seattle",
			"type":             "Đại chúng",
			"traded_as":        "NASDAQ:AMZN, NASDAQ-100 Component, S&P 100 Component, S&P 500 Component",
			"founded":          "05-07-1994 tại Bellevue, Washington, Hoa Kỳ",
			"founder":          "Jeff Bezos",
			"hq_location":      "Seattle, Washington",
			"area_served":      "Toàn cầu",
			"key_people":       "Andy Jassy(CEO), Werner Vogels(CTO)",
			"num_employees":    "798,000 (2019)",
			"industry":         "Điện toán đám mây, Trí tuệ nhân tạo, Hệ thống mạng, Phân phối kỹ thuật số, Thương mại điện tử, Điện tử gia dụng",
			"revenue":          "280,522 tỷ USD(2019)",
			"operating_income": "14,541 tỷ USD (2019)",
			"net_income":       "11,588 tỷ USD (2019)",
			"assets":           "255,248 tỷ USD (2019)",
			"equity":           "62,060 tỷ USD (2019)",
			"products":         "Amazon Echo, Amazon Kindle, Amazon Fire, Amazon Fire Tv, Amazon Fire OS",
			"website_type":     "thương mại điện tử",
			"language":         "tiếng Anh, tiếng Nhật, tiếng Đức, tiếng Pháp, tiếng Trung",
			"advertising":      "web banner và video",
			"launch_date":      "1995",
			"alexa":            "5",
			"website":          "https://www.amazon.com",
			"dịch vụ":          "Amazon.com, Amazon Alexa, Amazon Appstore, Amazon Music, Amazon Prime, Amazon Prime Video, Amazon Web Services",
		},
		MicrosoftInfobox: {
			"tên":                "Microsoft Corporation",
			"logo":               "Microsoft logo and wordmark.svg",
			"logo_caption":       "Logo Microsoft từ năm 2012 - nay",
			"image":              "Building92microsoft.jpg",
			"image_caption":      "Tòa nhà 92 trong khuôn viên Microsoft Redmond ở Redmond, Washington",
			"loại":               "Công ty cổ phần(NASDAQ:MSFT), Dow Jones Industrial Average Component, S&P 500 Component",
			"foundation":         "Albuquerque, New Mexico, 4 tháng 4 năm 1975",
			"founder":            "Bill Gates, Paul Allen",
			"location":           "Microsoft Redmond Campus, Redmond, Washington, Hoa Kỳ",
			"area_served":        "Toàn cầu",
			"key_people":         "Satya Nadella(CEO, chủ tịch), Bill Gates(Nhà sáng lập)",
			"industry":           "Phần mềm máy tính, Phần cứng máy tính, Điện thoại di động, Thiết bị viễn thông, Phân phối kỹ thuật số, Điện tử tiêu dùng, Trò chơi điện tử, Tư vấn công nghệ thông tin, Quảng cáo trực tuyến, Bán lẻ, Phần mềm ô-tô",
			"products":           "Xem Các sản phẩm",
			"services":           "Xem Các dịch vụ",
			"market cap":         "$169,72 tỉ(ngày 13 tháng 2 năm 2009)",
			"revenue":            "110,36 tỉ đô la Mỹ",
			"revenue_year":       "2018",
			"operating_income":   "35,05 tỉ đô la Mỹ",
			"income_year":        "2018",
			"net_income":         "16,57 tỉ đô la Mỹ",
			"net_income_year":    "2018",
			"assets":             "254.84 tỉ đô la Mỹ",
			"assets_year":        "2018",
			"equity":             "82,71 tỉ đô la Mỹ",
			"equity_year":        "2018",
			"num_employees":      "138,944",
			"num_employees_year": "2018",
			"website":            "microsoft.com",
			"image_size":         "300",
		},
		MetaInfobox: {
			"name":                "Meta Platforms, Inc.",
			"logo":                "Meta Inc. logo.svg",
			"type":                "Đại chúng",
			"image":               "Meta Platforms Headquarters Menlo Park California.jpg",
			"image_caption":       "Toàn cảnh trụ sở Meta tại Menlo Park",
			"traded_as":           "Lớp A: NASDAQ:FB, NASDAQ-100 component, S&P 100 component, S&P 500 component, Lớp B: Không niêm yết",
			"industry":            "Dịch vụ mạng xã hội, Quảng cáo",
			"founded":             "04-02-2004 tại Cambridge, Massachusetts",
			"founders":            "Mark Zuckerberg, Eduardo Saverin, Andrew McCollum, Dustin Moskovitz, Chris Hughes",
			"hq_location_city":    "Menlo Park, California",
			"hq_location_country": "U.S.",
			"area_served":         "Toàn thế giới (trừ các quốc gia kiểm duyệt)",
			"key_people":          "Mark Zuckerberg(Chủ tịch, CEO, Controlling Shareholder), Sheryl Sandberg(COO), David Wehner(CFO), Mike Schroepfer(CTO)",
			"products":            "Facebook, Instagram, Messenger, WhatsApp, Watch, Portal, Oculus, Calibra",
			"revenue":             "55,838 tỷ Mỹ kim",
			"revenue_year":        "2018",
			"operating_income":    "24,913 tỷ Mỹ kim",
			"income_year":         "2018",
			"net_income":          "22,111 tỷ Mỹ kim",
			"net_income_year":     "2018",
			"assets":              "97,334 tỷ Mỹ kim",
			"assets_year":         "2018",
			"equity":              "84,127 tỷ Mỹ kim",
			"equity_year":         "2018",
			"num_employees":       "43.030",
			"num_employees_year":  "ngày 30 tháng 9 năm 2019",
			"website":             "https://about.fb.com/",
			"footnotes":           "",
		},
		NetflixInfobox: {
			"name":             "Netflix, Inc.",
			"logo":             "Netflix 2015 logo.svg",
			"screenshot":       "netflixvn.jpeg",
			"caption":          "Ảnh chụp màn hình trang chủ Netflix phiên bản tiếng Việt",
			"company_type":     "Công ty đại chúng",
			"traded_as":        "NASDAQ:NFLX, NASDAQ-100, S&P 100, S&P 500",
			"area_served":      "Toàn cầu, ngoại trừ Trung Quốc đại lục, Syria, Bắc Triều Tiên và Bán đảo Krym",
			"founder":          "Reed Hastings, Marc Randolph",
			"industry":         "Giải trí, truyền thông đại chúng",
			"products":         "Streaming media, video on demand",
			"services":         "Nhà sản xuất phim, Nhà phân phối phim, Chương trình truyền hình",
			"divisions":        "US Streaming, International Streaming, Domestic DVD",
			"revenue":          "15.794 tỉ đô la Mỹ(2018)",
			"operating_income": "US$1.605 tỉ (2018)",
			"net_income":       "US$1.211 tỉ (2018)",
			"assets":           "US$25.974 tỉ (2018)",
			"equity":           "US$5.289 tỉ (2018)",
			"alexa":            "24(06-04-2019)",
			"registration":     "Cần thiết",
			"foundation":       "29-8-1997 tại Scotts Valley, California",
			"location":         "Los Gatos, California, Hoa Kỳ",
			"key_people":       "Reed Hastings(Chủ tịch, Giám đốc điều hành), Ted Sarandos(Giám đốc thương mại)",
			"subsid":           "DVD.com, Millarworld, LT-LA, ABQ Studios, Netflix Animation, Netflix Pte. Ltd., Netflix Services UK Limited, Netflix Streaming Services International B.V., Netflix Streaming Services, Inc., Netflix Global, LLC, Netflix Services Germany GmbH, NetflixCS, Inc., Netflix Luxembourg S.a r.l., Netflix Studios, Netflix Entretenimento Brasil LTDA., StoryBots",
			"num_users":        "154 triệu người (toàn cầu), 148 triệu người (đã thanh toán)",
			"num_employees":    "5,400 (2017)",
			"website":          "https://www.netflix.com",
		},
		TPBankInfobox: {
			"biểu trưng":          "Logo TPBank.svg",
			"loại":                "Tài chính",
			"thành lập":           "5/5/2008",
			"thành viên chủ chốt": "Đỗ Minh Phú (Chủ tịch HĐQT)",
			"sản phẩm":            "Dịch vụ tài chính, Ngân hàng",
			"tài sản":             "164,593 tỷ VNĐ (30/09/2019)",
			"khẩu hiệu":           "Vì chúng tôi hiểu bạn",
			"trang chủ":           "http://tpb.vn/ https://tpb.vn/",
			"Vốn điều lệ":         "8,565 tỷ VNĐ (30/09/2019)",
		},
	}
	for data, expected := range samples {
		p := NewParser(data)
		err := p.parse()
		assert.Equal(t, err, nil)
		// p.printTokens()
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
			"thành lập":            "01-04-1976(Cupertino, California, Mỹ)",
			"người sáng lập":       "Steve Jobs, Steve Wozniak, Ronald Wayne",
			"trụ sở chính":         "Cupertino, California, Mỹ",
			"số lượng trụ sở":      "510 cửa hàng bán lẻ(2020)",
			"khu vực hoạt động":    "Toàn thế giới",
			"thành viên chủ chốt":  "Tim Cook(CEO) Arthur D. Levinson(Chủ tịch hội đồng quản trị) Jeff Williams(COO)",
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
		AmazonInfobox: {
			"tên công ty":          "Amazon.com, Inc.",
			"logo":                 "Amazon logo.svg",
			"image":                "Seattle Spheres on May 10, 2018.jpg",
			"loại hình":            "Đại chúng",
			"mã niêm yết":          "NASDAQ:AMZN, NASDAQ-100 Component, S&P 100 Component, S&P 500 Component",
			"thành lập":            "05-07-1994 tại Bellevue, Washington, Hoa Kỳ",
			"người sáng lập":       "Jeff Bezos",
			"trụ sở chính":         "Seattle, Washington",
			"khu vực hoạt động":    "Toàn cầu",
			"thành viên chủ chốt":  "Andy Jassy(CEO), Werner Vogels(CTO)",
			"số nhân viên":         "798,000 (2019)",
			"ngành nghề":           "Điện toán đám mây, Trí tuệ nhân tạo, Hệ thống mạng, Phân phối kỹ thuật số, Thương mại điện tử, Điện tử gia dụng",
			"doanh thu":            "280,522 tỷ USD(2019)",
			"lợi nhuận kinh doanh": "14,541 tỷ USD (2019)",
			"lợi nhuận ròng":       "11,588 tỷ USD (2019)",
			"tổng tài sản":         "255,248 tỷ USD (2019)",
			"tổng vốn chủ sở hữu":  "62,060 tỷ USD (2019)",
			"sản phẩm":             "Amazon Echo, Amazon Kindle, Amazon Fire, Amazon Fire Tv, Amazon Fire OS",
			"website":              "https://www.amazon.com",
			"dịch vụ":              "Amazon.com, Amazon Alexa, Amazon Appstore, Amazon Music, Amazon Prime, Amazon Prime Video, Amazon Web Services",
		},
		MicrosoftInfobox: {
			"tên công ty":          "Microsoft Corporation",
			"logo":                 "Microsoft logo and wordmark.svg",
			"image":                "Building92microsoft.jpg",
			"loại hình":            "Công ty cổ phần(NASDAQ:MSFT), Dow Jones Industrial Average Component, S&P 500 Component",
			"thành lập":            "Albuquerque, New Mexico, 4 tháng 4 năm 1975",
			"người sáng lập":       "Bill Gates, Paul Allen",
			"trụ sở":               "Microsoft Redmond Campus, Redmond, Washington, Hoa Kỳ",
			"khu vực hoạt động":    "Toàn cầu",
			"thành viên chủ chốt":  "Satya Nadella(CEO, chủ tịch), Bill Gates(Nhà sáng lập)",
			"ngành nghề":           "Phần mềm máy tính, Phần cứng máy tính, Điện thoại di động, Thiết bị viễn thông, Phân phối kỹ thuật số, Điện tử tiêu dùng, Trò chơi điện tử, Tư vấn công nghệ thông tin, Quảng cáo trực tuyến, Bán lẻ, Phần mềm ô-tô",
			"sản phẩm":             "Xem Các sản phẩm",
			"dịch vụ":              "Xem Các dịch vụ",
			"doanh thu":            "110,36 tỉ đô la Mỹ(2018)",
			"lợi nhuận kinh doanh": "35,05 tỉ đô la Mỹ(2018)",
			"lợi nhuận ròng":       "16,57 tỉ đô la Mỹ(2018)",
			"tổng tài sản":         "254.84 tỉ đô la Mỹ(2018)",
			"tổng vốn chủ sở hữu":  "82,71 tỉ đô la Mỹ(2018)",
			"số nhân viên":         "138,944(2018)",
			"website":              "microsoft.com",
		},
		MetaInfobox: {
			"tên công ty":          "Meta Platforms, Inc.",
			"logo":                 "Meta Inc. logo.svg",
			"loại hình":            "Đại chúng",
			"image":                "Meta Platforms Headquarters Menlo Park California.jpg",
			"mã niêm yết":          "Lớp A: NASDAQ:FB, NASDAQ-100 component, S&P 100 component, S&P 500 component, Lớp B: Không niêm yết",
			"ngành nghề":           "Dịch vụ mạng xã hội, Quảng cáo",
			"thành lập":            "04-02-2004 tại Cambridge, Massachusetts",
			"người sáng lập":       "Mark Zuckerberg, Eduardo Saverin, Andrew McCollum, Dustin Moskovitz, Chris Hughes",
			"trụ sở chính":         "Menlo Park, California, U.S.",
			"khu vực hoạt động":    "Toàn thế giới (trừ các quốc gia kiểm duyệt)",
			"thành viên chủ chốt":  "Mark Zuckerberg(Chủ tịch, CEO, Controlling Shareholder), Sheryl Sandberg(COO), David Wehner(CFO), Mike Schroepfer(CTO)",
			"sản phẩm":             "Facebook, Instagram, Messenger, WhatsApp, Watch, Portal, Oculus, Calibra",
			"doanh thu":            "55,838 tỷ Mỹ kim(2018)",
			"lợi nhuận kinh doanh": "24,913 tỷ Mỹ kim(2018)",
			"lợi nhuận ròng":       "22,111 tỷ Mỹ kim(2018)",
			"tổng tài sản":         "97,334 tỷ Mỹ kim(2018)",
			"tổng vốn chủ sở hữu":  "84,127 tỷ Mỹ kim(2018)",
			"số nhân viên":         "43.030(ngày 30 tháng 9 năm 2019)",
			"website":              "https://about.fb.com/",
		},
		NetflixInfobox: {
			"tên công ty":          "Netflix, Inc.",
			"logo":                 "Netflix 2015 logo.svg",
			"mã niêm yết":          "NASDAQ:NFLX, NASDAQ-100, S&P 100, S&P 500",
			"khu vực hoạt động":    "Toàn cầu, ngoại trừ Trung Quốc đại lục, Syria, Bắc Triều Tiên và Bán đảo Krym",
			"người sáng lập":       "Reed Hastings, Marc Randolph",
			"ngành nghề":           "Giải trí, truyền thông đại chúng",
			"sản phẩm":             "Streaming media, video on demand",
			"dịch vụ":              "Nhà sản xuất phim, Nhà phân phối phim, Chương trình truyền hình",
			"doanh thu":            "15.794 tỉ đô la Mỹ(2018)",
			"lợi nhuận kinh doanh": "US$1.605 tỉ (2018)",
			"lợi nhuận ròng":       "US$1.211 tỉ (2018)",
			"tổng tài sản":         "US$25.974 tỉ (2018)",
			"tổng vốn chủ sở hữu":  "US$5.289 tỉ (2018)",
			"thành lập":            "29-8-1997 tại Scotts Valley, California",
			"trụ sở":               "Los Gatos, California, Hoa Kỳ",
			"thành viên chủ chốt":  "Reed Hastings(Chủ tịch, Giám đốc điều hành), Ted Sarandos(Giám đốc thương mại)",
			"công ty con":          "DVD.com, Millarworld, LT-LA, ABQ Studios, Netflix Animation, Netflix Pte. Ltd., Netflix Services UK Limited, Netflix Streaming Services International B.V., Netflix Streaming Services, Inc., Netflix Global, LLC, Netflix Services Germany GmbH, NetflixCS, Inc., Netflix Luxembourg S.a r.l., Netflix Studios, Netflix Entretenimento Brasil LTDA., StoryBots",
			"số nhân viên":         "5,400 (2017)",
			"website":              "https://www.netflix.com",
		},
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
