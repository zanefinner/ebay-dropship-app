package api

const  (
	oAuth = `Bearer v^1.1#i^1#r^0#I^3#p^1#f^0#t^H4sIAAAAAAAAAOVYW2wUVRjutN2acpMaRMTarANIIpnduezsZWA32V6gq73ZLRWKhMzlTHvY3Zlhzqy7i4BNg5jqi8EQiTGhD1w0pkKJhgdDFDGEaIwXqqISFFIVG3lAjPii8czuUraVtEg32MR92cx//vOf7/++/z/nzNC9FZWP7GrcdW0ucVfpQC/dW0oQzGy6ssKxYl5Z6WJHCV3gQAz0Lu0t7yu7tAqJibghtANk6BoCznQiriEhawySSVMTdBFBJGhiAiDBkoVouLlJYF20YJi6pct6nHRG6oNkgGElBbB+hZV8PMMy2Kpdj9mhB0mv6uMVlmM4TuJlnvfjcYSSIKIhS9SsIMnSLE3RPMX6OxiPwPMCz7m8nLeLdHYCE0Fdwy4umgxl4QrZuWYB1smhiggB08JByFAkvDraGo7UN7R0rHIXxArleYhaopVE45/qdAU4O8V4Eky+DMp6C9GkLAOESHcot8L4oEL4OpjbgJ+lWmYkL6v6OQ4A2gu8UlGoXK2bCdGaHIdtgQqlZl0FoFnQykzFKGZD2gxkK//UgkNE6p323+NJMQ5VCMwg2VAbXh9uayNDXaIGVkNNoxrSBjRhAlBt7fWUjwWSJ+DxcJQsB1RZZQL5hXLR8jRPWKlO1xRok4acLbpVCzBqMJEbTwE32KlVazXDqmUjKvTzXueQDXTZouZUTFo9mq0rSGAinNnHqRUYm21ZJpSSFhiLMHEgS1GQFA0DKuTEwWwt5ssnjYJkj2UZgtudSqVcKc6lm91ulqYZ97rmpqjcAxIiiX3tXs/5w6knUDCbigzwTAQFK2NgLGlcqxiA1k2GPIyHD/jyvI+HFZpo/YehIGf3+I4oVod4FNXD+1QPkETASZKnGB0Syhep28aBA2eohGjGgGXERRlQMq6zZAKYUBE4XmU5vwooxRtQKU9AVSmJV7wUo+J+BUCS5ID//9Qot1rqUSCbwCpKrRetzt2elvpYrFOFjz7RnqQb1qRam5kMV5eJJhtjrVI7jgbWp72cDNc1B2+1G26evKwboE2PQzlTBAbsXi8iC5yptImmlYmCeBwbppUoshOdWSLb8xEOIBrQZTe2S9YTbl3EO7pt2pRFPK2cw4YRSSSSlijFQaQ4u/l/tJPfND2I7zozKiesX05IqOQuKa6smi70lOwyAdKTJr6fuVrtM7tDjwEN74CWqcfjwOxkpi30ndbX7vUp+PiXh8Xt5V68m8pMqm05DnEJbZppmd0RRaE4w05jhg/QXp/Px3PTyqsuq2lHZqadQ406soAyWWrla27zWu0e/5IfKsn+mD7ibbqPGColCNpNL2OW0A9VlK0tL5uzGEELuKCouhDs1vC7qwlcMZAxRGiWVhDGWnF0WcFnhYGN9KKxDwuVZczsgq8MdPWNEQdz931zMSU868eNyfNcF73kxmg5s7B8wUhk5aHLbNXS/uHNj0VYc88bu71f03PHnAjCUVLeR5QsPjx64p5ZV698uDe0tPHdknRTT9vW/d2Oh+994dm91ecWXXyg79SW5/svOrhvXyEOjlYfjTUNyAv/6n8zpdS+dhyeOLat5q2R6vmdT1ftyFQyg73f7/jmR/+DJ5kfzinlQ2uQtBl5v1oxfOjksiPuJy/4R4ccny24OnT+1ZENy6s+uHZgO/feudMv/Tan59SGA01X3983650r+37/eOfA/MHDB/uPn/wk5D2f+m7LR7/+8nLsGEmcHdwTJRY9VxrkF56ZcyZ95VJy3f2vD/6x3urf/vP+dI24f+cXGytXfPl5bPjFbT8N1DQG3A8dGP5z/jyhxFN7NLycoM9eWLnzSM/uT6tTI8/Ql/tqtlZdO52T72+nh2NG8BEAAA==`
	id    = "ZaneFinn-Expirime-PRD-72eb49443-cc9fcf19"
)
//Use ...
func Use(r http.ResponseWriter, w *http.Request /*uri*/) {
	var (
		resp *http.Response
		err error
		req *http.Request
		body []byte
	)
	action := "show-item"//derived uri1
	param := "drone" //derived uri2
	client := &http.Response{}
	if action == "show-item" {
		resp, err = client.Get("https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search?q=" + param + "&limit=3")
		// ...

		req, err = http.NewRequest("GET", "https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search?q=drone&limit=3", nil)
		// ...
		req.Header.Add("Authorization", oAuthToken)
		req.Header.Add("Content-Type", `application/json`)
		req.Header.Add("X-EBAY-C-ENDUSERCTX", `affiliateCampaignId=<ePNCampaignId>,affiliateReferenceId=<referenceId>`)
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		r.Write(body)

	} else {
		//show 404
	}
}
