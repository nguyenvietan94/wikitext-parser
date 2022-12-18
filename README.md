# wikitext-parser
A Wikitext parser that extracts info from wikipedia infoboxes.

### Features
This parser can extract info from many Wikipedia infobox templates (see [List of Wikipedia Infobox templates](https://en.wikipedia.org/wiki/Wikipedia:List_of_infoboxes)), which are written in Wikitext. A variety of Wikitext types are supported: [template](https://en.wikipedia.org/wiki/Help:Template), [list](https://en.wikipedia.org/wiki/Help:List), [wiki links](https://en.wikipedia.org/wiki/Help:Link), and so on.
### Get the source code
```
$ git clone https://github.com/nguyenvietan94/wikitext-parser.git
```
### Download all the required packages
```
$ go get ./...
```
### Run the code:
Example input: A [wikitext infobox](https://en.wikipedia.org/w/index.php?title=Apple_Inc.&action=edit) extracted from [Apple Inc. wikipedia article](https://en.wikipedia.org/wiki/Apple_Inc.).
```
{{Hộp thông tin tóm tắt về công ty
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
}}
```
Let's run the code to get the output, which is a set of fields and their string values extracted from the wikitext.
```
$ go run main.go
name: Apple Inc.
foundation: 01-04-1976(Cupertino, California, Mỹ)
divisions: 
type: Công ty đại chúng
assets: 323.888 tỉ USD
founder: Steve Jobs, Steve Wozniak, Ronald Wayne
subsid: Shazam, FileMaker Inc., Anobit, Braeburn Capital, Beats Electronics
location_country: Mỹ
num_locations_year: 2020
revenue: 274.515 tỉ USD
income_year: 2020
products: Mac, iPod, iPhone, iPad, Apple Watch, Apple TV, macOS, iOS, iPadOS, watchOS, tvOS
assets_year: 2020
industry: Phần cứng máy tính · Phần mềm máy tính, phụ kiện, thiết bị di động
revenue_year: 2020
net_income_year: 2020
num_employees_year: 2020
equity: 65.339 tỉ USD
operating_income: 66.288 tỉ USD
equity_year: 2020
key_people: Tim Cook(CEO) Arthur D. Levinson(Chủ tịch hội đồng quản trị) Jeff Williams(COO)
image_caption: Trụ sở Apple ở Cupertino, California
predecessor: 
area_served: Toàn thế giới
net_income: 57.411 tỉ USD
traded_as: NASDAQ:AAPL, LSE:0HDZ, FWB:APC
image_size: 260px
services: Apple Arcade, Apple Card, Apple Music, Apple News+, Apple TV+, Apple Store online, App Store, iTunes Store, Mac App Store, iBooks, iCloud, Apple Pay, iMessage, FaceTime
image: Apple park cupertino 2019.jpg
num_employees: 147,000
intl: yes
logo: Apple logo black.svg
location_city: Cupertino, California
num_locations: 510 cửa hàng bán lẻ
website: apple.com
```
### Applications
This project is part of the Coc Coc Knowledge Graph Project, which processes millions of Wikipedia articles in Wikitext. The extracted info then is displayed as an infobox.
Examples: Let's see some infoboxes via Coc Coc search engine: [Coc Coc](https://coccoc.com/search?query=coc+coc), [Apple Inc.](https://coccoc.com/search?query=apple+inc.).
