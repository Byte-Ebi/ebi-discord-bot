package drawing

import (
	"math/rand"
)

var (
	RemoveCommands = true
)

var holomens = []string{
	"ときのそら",
	"ロボ子さん",
	"さくらみこ",
	"星街すいせい",
	"AZKi",
	"夜空メル",
	"アキ・ローゼンタール",
	"赤井はあと",
	"白上フブキ",
	"夏色まつり",
	"湊あくあ",
	"紫咲シオン",
	"百鬼あやめ",
	"癒月ちょこ",
	"大空スバル",
	"大神ミオ",
	"猫又おかゆ",
	"戌神ころね",
	"兎田ぺこら",
	"不知火フレア",
	"白銀ノエル",
	"宝鐘マリン",
	"天音かなた",
	"桐生ココ",
	"角巻わため",
	"常闇トワ",
	"姫森ルーナ",
	"雪花ラミィ",
	"桃鈴ねね",
	"獅白ぼたん",
	"尾丸ポルカ",
	"ラプラス・ダークネス",
	"鷹嶺ルイ",
	"博衣こより",
	"沙花叉クロヱ",
	"風真いろは",
	"友人A",
	"春先のどか",
	"Ayunda Risu",
	"Moona Hoshinova",
	"Airani Iofifteen",
	"Kureiji Ollie",
	"Anya Melfissa",
	"Pavolia Reine",
	"Vestia Zeta",
	"Kaela Kovalskia",
	"Kobo Kanaeru",
	"Mori Calliope",
	"Takanashi Kiara",
	"Ninomae Ina'nis",
	"Gawr Gura",
	"Watson Amelia",
	"IRyS",
	"Tsukumo Sana",
	"Ceres Fauna",
	"Ouro Kronii",
	"Nanashi Mumei",
	"Hakos Baelz",
}

var quotes = map[string][]string{
	"ときのそら":            {},
	"ロボ子さん":            {},
	"さくらみこ":            {},
	"星街すいせい":           {},
	"AZKi":             {},
	"夜空メル":             {},
	"アキ・ローゼンタール":       {},
	"赤井はあと":            {},
	"白上フブキ":            {},
	"夏色まつり":            {},
	"湊あくあ":             {},
	"紫咲シオン":            {},
	"百鬼あやめ":            {},
	"癒月ちょこ":            {},
	"大空スバル":            {},
	"大神ミオ":             {},
	"猫又おかゆ":            {},
	"戌神ころね":            {},
	"兎田ぺこら":            {},
	"不知火フレア":           {},
	"白銀ノエル":            {},
	"宝鐘マリン":            {},
	"天音かなた":            {},
	"桐生ココ":             {},
	"角巻わため":            {},
	"常闇トワ":             {},
	"姫森ルーナ":            {},
	"雪花ラミィ":            {},
	"桃鈴ねね":             {},
	"獅白ぼたん":            {},
	"尾丸ポルカ":            {},
	"ラプラス・ダークネス":       {},
	"鷹嶺ルイ":             {},
	"博衣こより":            {},
	"沙花叉クロヱ":           {},
	"風真いろは":            {},
	"友人A":              {},
	"春先のどか":            {},
	"Ayunda Risu":      {},
	"Moona Hoshinova":  {},
	"Airani Iofifteen": {},
	"Kureiji Ollie":    {},
	"Anya Melfissa":    {},
	"Pavolia Reine":    {},
	"Vestia Zeta":      {},
	"Kaela Kovalskia":  {},
	"Kobo Kanaeru":     {},
	"Mori Calliope":    {},
	"Takanashi Kiara":  {},
	"Ninomae Ina'nis":  {},
	"Gawr Gura":        {},
	"Watson Amelia":    {},
	"IRyS":             {},
	"Tsukumo Sana":     {},
	"Ceres Fauna":      {},
	"Ouro Kronii":      {},
	"Nanashi Mumei":    {},
	"Hakos Baelz":      {},
}

func getHolomen(seed int64) string {
	s := rand.NewSource(seed)
	r := rand.New(s)
	num := r.Intn(len(holomens))

	holomen := holomens[num]
	return holomen
}

func getQuote(seed int64, holomen string) string {
	s := rand.NewSource(seed)
	r := rand.New(s)
	num := r.Intn(len(quotes[holomen]))

	quote := quotes[holomen][num]
	return quote
}
