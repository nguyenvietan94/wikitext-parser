package main

import (
	"fmt"
	"wikitext-parser/parser"
)

func main() {
	data := `{{Hộp thông tin tóm tắt về công ty
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

	parser := parser.NewParser(data)
	err := parser.Parse()
	if err != nil {
		panic(err)
	}
	// p.printTokens()
	params := parser.GetParams()
	for key, value := range params {
		fmt.Printf("%s: %s\n", key, value)
	}
}
