package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPreprocessTemplateName(t *testing.T) {
	samples := [][]string{
		{"hộp thông tin quốc gia", "quốc gia"},
		{"infobox country", "country"},
		{"thông tin diễn viên", "diễn viên"},
	}
	for _, sample := range samples {
		assert.Equal(t, len(sample), 2)
		templateName := sample[0]
		expected := sample[1]
		assert.Equal(t, PreprocessTemplateName(templateName), expected)
	}
}

func TestGetWikiInfoboxFromText(t *testing.T) {
	samples := map[string]string{
		`{{chú thích trong bài}}
		{{Tóm tắt về công ty
		| tên = Ngân hàng TMCP Ngoại thương Việt Nam<br/>(Vietcombank)
		[[Tập tin:Vietcombank Logo.png|250px]]
		| chú thích biểu trưng = Logo ngân hàng Vietcombank được sử dụng từ ngày 1/4/2013.
		| loại = [[Công ty cổ phần|Doanh nghiệp cổ phần]]
		| thành lập = 01/04/1963
		| thành viên chủ chốt = [[Phạm Quang Dũng]] - Chủ tịch Hội đồng quản trị|Chủ tịch Hội đồng quản trị<ref>[https://vnexpress.net/vietcombank-co-chu-tich-moi-4348794.html|title= Vietcombank có chủ tịch mới]</ref>
		| lĩnh vực = [[Tài chính]]
		| ngành = [[Ngân hàng]]
		| sản phẩm = [[Khu vực tài chính|Dịch vụ tài chính]]
		| tài sản = >1.300.000 tỷ [[đồng (tiền)|đồng]] (2021)
		| thu nhập =880.000 nghìn tỷ đồng 
		| thu nhập trước lãi vay và thuế =188 nghìn tỷ đồng 
		| số nhân viên = 20.115 (2020)
		| công ty mẹ =[[Ngân hàng Nhà nước Việt Nam]] 
		| trang chủ = https://vietcombank.com.vn
		| ghi chú = [[Vietcombank]]
		|logo=Vietcombank Logo.png|slogan=Chung niềm tin, vững tương lai}}
		'''Ngân hàng TMCP Ngoại thương Việt Nam'''
		(tên giao dịch quốc tế: ''Joint Stock Commercial Bank for Foreign Trade of Vietnam'')
		tên viết tắt: "'''Vietcombank'''", là công ty lớn nhất trên [[thị trường chứng khoán]]
		[[Việt Nam]] tính theo vốn hóa.`: `{{Tóm tắt về công ty
		| tên = Ngân hàng TMCP Ngoại thương Việt Nam<br/>(Vietcombank)
		[[Tập tin:Vietcombank Logo.png|250px]]
		| chú thích biểu trưng = Logo ngân hàng Vietcombank được sử dụng từ ngày 1/4/2013.
		| loại = [[Công ty cổ phần|Doanh nghiệp cổ phần]]
		| thành lập = 01/04/1963
		| thành viên chủ chốt = [[Phạm Quang Dũng]] - Chủ tịch Hội đồng quản trị|Chủ tịch Hội đồng quản trị<ref>[https://vnexpress.net/vietcombank-co-chu-tich-moi-4348794.html|title= Vietcombank có chủ tịch mới]</ref>
		| lĩnh vực = [[Tài chính]]
		| ngành = [[Ngân hàng]]
		| sản phẩm = [[Khu vực tài chính|Dịch vụ tài chính]]
		| tài sản = >1.300.000 tỷ [[đồng (tiền)|đồng]] (2021)
		| thu nhập =880.000 nghìn tỷ đồng 
		| thu nhập trước lãi vay và thuế =188 nghìn tỷ đồng 
		| số nhân viên = 20.115 (2020)
		| công ty mẹ =[[Ngân hàng Nhà nước Việt Nam]] 
		| trang chủ = https://vietcombank.com.vn
		| ghi chú = [[Vietcombank]]
		|logo=Vietcombank Logo.png|slogan=Chung niềm tin, vững tương lai}}`,
		`{{Underlinked}}
			{{prose}}
			{{Nhiều vấn đề
			|collapsed=no
			|
			{{Mở đầu quá dài}}
			
			}}
			{{Tóm tắt về công ty
			| biểu trưng      = Tập tin:Logo TPBank.svg
			| loại            = Tài chính
			| thành lập       = 5/5/2008
			| thành viên chủ chốt = Đỗ Minh Phú (Chủ tịch HĐQT)
			| sản phẩm        = Dịch vụ tài chính, Ngân hàng
			| tài sản         = 164,593 tỷ VNĐ (30/09/2019)
			| khẩu hiệu       = Vì chúng tôi hiểu bạn
			| trang chủ       = [http://tpb.vn/ https://tpb.vn/]
			|Vốn điều lệ=8,565 tỷ VNĐ (30/09/2019)}}'''Ngân hàng Thương mại Cổ phần Tiên Phong''' (hay '''TPBank''') là một [[ngân hàng thương mại cổ phần]]
			[[Danh sách ngân hàng tại Việt Nam|Việt Nam]] được thành lập ngày 05/05/2008 bởi các cổ đông chủ chốt gồm Công ty Cổ phần Tập đoàn Vàng bạc Đá quý DOJI,
			[[Công ty cổ phần FPT]], Công ty Tài chính quốc tế (IFC), Tổng công ty Tái bảo hiểm Việt Nam (Vinare) và SBI Ven Holding Pte. Ltd.,Singapore.`: `{{Tóm tắt về công ty
			| biểu trưng      = Tập tin:Logo TPBank.svg
			| loại            = Tài chính
			| thành lập       = 5/5/2008
			| thành viên chủ chốt = Đỗ Minh Phú (Chủ tịch HĐQT)
			| sản phẩm        = Dịch vụ tài chính, Ngân hàng
			| tài sản         = 164,593 tỷ VNĐ (30/09/2019)
			| khẩu hiệu       = Vì chúng tôi hiểu bạn
			| trang chủ       = [http://tpb.vn/ https://tpb.vn/]
			|Vốn điều lệ=8,565 tỷ VNĐ (30/09/2019)}}`,
		`{{Redirect2|Joseph Biden|Biden|người con trai quá cố của ông là Joseph Biden III|Beau Biden|các định nghĩa|Biden (định hướng)}}
			{{Thông tin viên chức 1
			| name          = Joe Biden
			| image         = Joe Biden presidential portrait.jpg
			| imagesize     = 220px
			| caption       = Biden vào năm 2021
			| office        = Tổng thống Hoa Kỳ
			|1namedata      = [[Kamala Harris]]
			|1blankname     = Phó Tổng thống
			| term_start          = 20 tháng 1 năm 2021
			| predecessor         = [[Donald Trump]]
			| order              = thứ 46
			| order2             = thứ 47
			| office2             = Phó Tổng thống Hoa Kỳ
			| president2          = [[Barack Obama]]
			| term_start2         = 20 tháng 1 năm 2009
			| term_end2           = 20 tháng 1 năm 2017
			| predecessor2        = [[Dick Cheney]]
			| successor2          = [[Mike Pence]]
			| jr/sr3              = [[Thượng viện Hoa Kỳ|Thượng nghị sĩ Hoa Kỳ]]
			| state3              = [[Delaware]]
			| term_start3         = 3 tháng 1 năm 1973
			| term_end3           = 15 tháng 1 năm 2009
			| predecessor3        = [[J. Caleb Boggs]]
			| successor3          = [[Ted Kaufman]]
			{{collapsed infobox section begin|last=yes|Các chức vụ khác
			|titlestyle = border:1px dashed lightgrey}}{{Infobox officeholder1|embed = yes
			| office4             = Chủ tịch [[Hội nghị của thượng viện Hoa Kỳ về kiểm soát ma túy quốc tế|Hội nghị về kiểm soát ma túy quốc tế]]
			| term_start4        = 2007
			| term_end4          = 2009
			| office5            = Chủ tịch [[Ủy ban Đối ngoại Thượng viện Hoa Kỳ|Ủy ban Đối ngoại Thượng viện]]
			| term_start5         = 2001{{efn|Biden giữ chức chủ tịch từ ngày 1 đến ngày 20 tháng 1, rồi được [[Jesse Helms]] kế nhiệm cho đến ngày 6 tháng 6, sau đó lại giữ chức vụ này cho đến năm 2003.}}
			| term_end5          = 2003
			| term_start6        = 2007
			| term_end6          = 2009
			| office7            = Chủ tịch [[Ủy ban Tư pháp Thượng viện Hoa Kỳ|Ủy ban Tư pháp Thượng viện]]
			| term_start7          = 1987
			| term_end7            = 1995
			{{Collapsed infobox section end}}
			}}
			| birth_name          = Joseph Robinette Biden Jr.
			| birth_date          = {{Birth date and age|1942|11|20}}
			| birth_place   =  [[Scranton, Pennsylvania]], Hoa Kỳ
			| death_date    =
			| death_place   =
			| spouse        = {{plainlist|
			* {{marriage|[[Neilia Hunter Biden|Neilia Hunter]]|ngày 27 tháng 8 năm 1966|ngày 18 tháng 12 năm 1972|reason=died}}
			* {{marriage|[[Jill Biden|Jill Jacobs]]|ngày 17 tháng 6 năm 1977}}
			}}
			| children      = {{flatlist|
			* [[Beau Biden|Beau]]
			* [[Hunter Biden|Hunter]]
			* Naomi
			* [[Ashley Biden|Ashley]]}}
			| parents       = {{plainlist|
			* Joseph Robinette Biden Sr.
			* Catherine Eugenia Finnegan}}
			| relatives     = [[Gia đình Biden]]
			| party         = [[Đảng Dân chủ (Hoa Kỳ)|Dân chủ]] (1969–nay)
			| otherparty          = [[Chính khách độc lập|Độc lập]] (trước 1969)
			| occupation    = {{Hlist|[[Luật sư]]|[[chính trị gia]]|[[nhà văn]]}}
			| education     = {{plainlist|
			* [[Đại học Delaware]] ([[Cử nhân Nghệ thuật|BA]])
			* [[Đại học Syracuse]] ([[Tiến sĩ Luật|JD]])}}
			| awards        = [[Danh sách danh hiệu và giải thưởng của Joe Biden|Danh sách danh hiệu và giải thưởng]]
			| signature     = Joe Biden Signature.svg
			| website       = {{plainlist|
			* {{URL|joebiden.com|Trang web tranh cử}}
			* {{URL|whitehouse.gov/administration/president-biden/|Trang web ở Nhà Trắng}}
			}}
			}}
			
			'''Joseph Robinette Biden Jr.''' ({{IPAc-en|ˈ|b|aɪ|d|ən}} {{respell|BY|dən}}; sinh ngày 20 tháng 11 năm 1942) là một [[chính trị gia]] [[người Mỹ]], [[Tổng thống Hoa Kỳ|tổng thống thứ 46]] và đương nhiệm của Hoa Kỳ. Là thành viên của [[Đảng Dân chủ (Hoa Kỳ)|Đảng Dân chủ]], ông từng là [[Danh sách Phó Tổng thống Hoa Kỳ|phó tổng thống thứ 47]] từ năm 2009 đến năm 2017 dưới thời [[Barack Obama]] và [[Thượng viện Hoa Kỳ|thượng nghị sĩ]] đại diện cho [[Delaware]] từ năm 1973 đến năm 2009.`: `{{Thông tin viên chức 1
			| name          = Joe Biden
			| image         = Joe Biden presidential portrait.jpg
			| imagesize     = 220px
			| caption       = Biden vào năm 2021
			| office        = Tổng thống Hoa Kỳ
			|1namedata      = [[Kamala Harris]]
			|1blankname     = Phó Tổng thống
			| term_start          = 20 tháng 1 năm 2021
			| predecessor         = [[Donald Trump]]
			| order              = thứ 46
			| order2             = thứ 47
			| office2             = Phó Tổng thống Hoa Kỳ
			| president2          = [[Barack Obama]]
			| term_start2         = 20 tháng 1 năm 2009
			| term_end2           = 20 tháng 1 năm 2017
			| predecessor2        = [[Dick Cheney]]
			| successor2          = [[Mike Pence]]
			| jr/sr3              = [[Thượng viện Hoa Kỳ|Thượng nghị sĩ Hoa Kỳ]]
			| state3              = [[Delaware]]
			| term_start3         = 3 tháng 1 năm 1973
			| term_end3           = 15 tháng 1 năm 2009
			| predecessor3        = [[J. Caleb Boggs]]
			| successor3          = [[Ted Kaufman]]
			{{collapsed infobox section begin|last=yes|Các chức vụ khác
			|titlestyle = border:1px dashed lightgrey}}{{Infobox officeholder1|embed = yes
			| office4             = Chủ tịch [[Hội nghị của thượng viện Hoa Kỳ về kiểm soát ma túy quốc tế|Hội nghị về kiểm soát ma túy quốc tế]]
			| term_start4        = 2007
			| term_end4          = 2009
			| office5            = Chủ tịch [[Ủy ban Đối ngoại Thượng viện Hoa Kỳ|Ủy ban Đối ngoại Thượng viện]]
			| term_start5         = 2001{{efn|Biden giữ chức chủ tịch từ ngày 1 đến ngày 20 tháng 1, rồi được [[Jesse Helms]] kế nhiệm cho đến ngày 6 tháng 6, sau đó lại giữ chức vụ này cho đến năm 2003.}}
			| term_end5          = 2003
			| term_start6        = 2007
			| term_end6          = 2009
			| office7            = Chủ tịch [[Ủy ban Tư pháp Thượng viện Hoa Kỳ|Ủy ban Tư pháp Thượng viện]]
			| term_start7          = 1987
			| term_end7            = 1995
			{{Collapsed infobox section end}}
			}}
			| birth_name          = Joseph Robinette Biden Jr.
			| birth_date          = {{Birth date and age|1942|11|20}}
			| birth_place   =  [[Scranton, Pennsylvania]], Hoa Kỳ
			| death_date    =
			| death_place   =
			| spouse        = {{plainlist|
			* {{marriage|[[Neilia Hunter Biden|Neilia Hunter]]|ngày 27 tháng 8 năm 1966|ngày 18 tháng 12 năm 1972|reason=died}}
			* {{marriage|[[Jill Biden|Jill Jacobs]]|ngày 17 tháng 6 năm 1977}}
			}}
			| children      = {{flatlist|
			* [[Beau Biden|Beau]]
			* [[Hunter Biden|Hunter]]
			* Naomi
			* [[Ashley Biden|Ashley]]}}
			| parents       = {{plainlist|
			* Joseph Robinette Biden Sr.
			* Catherine Eugenia Finnegan}}
			| relatives     = [[Gia đình Biden]]
			| party         = [[Đảng Dân chủ (Hoa Kỳ)|Dân chủ]] (1969–nay)
			| otherparty          = [[Chính khách độc lập|Độc lập]] (trước 1969)
			| occupation    = {{Hlist|[[Luật sư]]|[[chính trị gia]]|[[nhà văn]]}}
			| education     = {{plainlist|
			* [[Đại học Delaware]] ([[Cử nhân Nghệ thuật|BA]])
			* [[Đại học Syracuse]] ([[Tiến sĩ Luật|JD]])}}
			| awards        = [[Danh sách danh hiệu và giải thưởng của Joe Biden|Danh sách danh hiệu và giải thưởng]]
			| signature     = Joe Biden Signature.svg
			| website       = {{plainlist|
			* {{URL|joebiden.com|Trang web tranh cử}}
			* {{URL|whitehouse.gov/administration/president-biden/|Trang web ở Nhà Trắng}}
			}}
			}}`,
	}
	for text, expected := range samples {
		infobox := GetWikiInfoboxFromText(text)
		assert.Equal(t, infobox, expected)
	}
}
