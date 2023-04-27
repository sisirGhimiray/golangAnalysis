package makeStockList

import (
	"os"
	"strings"

	makejson "example.com/golang/makeJson"
	"example.com/golang/utils"
)
type PreMarket struct {
	Symbol                      string      `json:"symbol"`
	PreviousClose               interface{} `json:"previousClose"`
	IndicativeEquillibriumPrice interface{} `json:"indicativeEquillibriumPrice"`
	Change                      interface{} `json:"change"`
	PercentageChange            interface{} `json:"percentageChange"`
	Final                       interface{} `json:"final"`
	FinalQuantity               interface{} `json:"finalQuantity"`
	Value                       interface{} `json:"value"`
	NormalMarket52WeekHigh      interface{} `json:"normalMarket52WeekHigh"`
	NormalMarket52WeekLow       interface{} `json:"normalMarket52WeekLow"`
	Secure                      string      `json:"secure"`
}

type EquityAndSmeMarket struct{
	Symbol 	string `json:"symbol"`
	Open 	interface{} `json:"open"`
	High 	interface{}	`json:"high"`
	Low		interface{}	`json:"low"`
	PrevClose interface{} `json:"prevClose"`
	LTP		interface{}	`json:"ltp"`
	Change	interface{}	`json:"change"`
	PercentChange interface{} `json:"percentChange"`
	Volume	interface{}			`json:"volume"`
	Value	interface{}		`json:"value"`
	Fifty2WeekL interface{}	`json:"fifty2WeekL"`
	Fifty2WeekH interface{}	`json:"fifty2WeekH"`

}

var (
	eqSmetempObj EquityAndSmeMarket
	preMarketObj PreMarket
	equityAndSmeMarketList []EquityAndSmeMarket
	preMarketList []PreMarket
	rawStr string
	splitedRawStr []string
	track int =0
	tdTrack int =0
	
)





//  this function makes data object we should mention that it's for niftyCurrent or preMarket 
// the symbol is where we have to put pre or nif token
func mk_Eq_Sme_Obj(str string){
	var s string
	
	
		substr:="</td>"
		substr2:="<span"
		substr3:="</a>"
		// substr2:="<span"
		
			idx:=strings.Index(str,substr)
			idx2:=strings.Index(str,substr2)
			idx3:=strings.Index(str,substr3)
			
		
	
		if(idx>0){
		
			
			s=utils.PrintReverseDataByte(str,idx-1)
			
			
		if s==""{
			track++
			if track==1{
				s=utils.PrintReverseDataByte(str,idx3-1)
				// println(s,track)
			}else{
				s=utils.PrintReverseDataByte(str,idx2-1)
				
				// println(s,track)
	
			}
		}else{
			track++
			// println(s,track)
		}





	switch track {
	case 1:
			eqSmetempObj.Symbol=s
	case 2:
			eqSmetempObj.Open=utils.StrToNumTheObj(s)
	case 3:
			eqSmetempObj.High=utils.StrToNumTheObj(s)
	case 4:
			eqSmetempObj.Low=utils.StrToNumTheObj(s)
	case 5:
			eqSmetempObj.PrevClose=utils.StrToNumTheObj(s)
	case 6:
			eqSmetempObj.LTP=utils.StrToNumTheObj(s)
	case 7:
			eqSmetempObj.Change=utils.StrToNumTheObj(s)
	case 8:
			eqSmetempObj.PercentChange=utils.StrToNumTheObj(s)
	case 9:
			eqSmetempObj.Volume=utils.StrToNumTheObj(s)
	case 10:
			eqSmetempObj.Value=utils.StrToNumTheObj(s)
	case 11:
			eqSmetempObj.Fifty2WeekL=utils.StrToNumTheObj(s)
	case 12:
			eqSmetempObj.Fifty2WeekH=utils.StrToNumTheObj(s)
			track=0
			equityAndSmeMarketList = append(equityAndSmeMarketList, eqSmetempObj)
	}







	mk_Eq_Sme_Obj(str[idx+len(substr):])
}
}




func mk_pre_mrkt_obj(str string){
	// splitting str with td
	substr:="</td>"
	
	var s string
	idx:=strings.Index(str,substr);

	if idx>0{
		// track how many td are passed
		tdTrack+=1
		
		s=utils.PrintReverseDataByte(str,idx-1)
		



	
	switch tdTrack {
	case 1:
		// print(s)
	case 2:
		if s==""{

			s=utils.PrintReverseDataByte(str,idx-5)
			preMarketObj.Symbol=s
		}
	case 3:
		preMarketObj.PreviousClose=utils.StrToNumTheObj(s)
	case 4:
		preMarketObj.IndicativeEquillibriumPrice=utils.StrToNumTheObj(s)
	case 5:
		preMarketObj.Change=utils.StrToNumTheObj(s)
	case 6:
		preMarketObj.PercentageChange=utils.StrToNumTheObj(s)
	case 7:
		preMarketObj.Final=utils.StrToNumTheObj(s)
	case 8:
		preMarketObj.FinalQuantity=utils.StrToNumTheObj(s)
	case 9:
		preMarketObj.Value=utils.StrToNumTheObj(s)
	case 10:
		preMarketObj.NormalMarket52WeekHigh=utils.StrToNumTheObj(s)
	case 11:
		preMarketObj.NormalMarket52WeekLow=utils.StrToNumTheObj(s)
		preMarketList = append(preMarketList, preMarketObj)
		
	case 12:
		tdTrack=0
	}





	mk_pre_mrkt_obj(str[idx+len(substr):])
	
}

}




func ReadFile(fileName string){
// Reading the Raw Stock Data Text take from the nse website
	file,err:=os.ReadFile(fileName)

	if err!=nil{
		panic(err)
	}

// Raw stock data save in the rawStr variable
rawStr=string(file)

//  spliting the raw str and updated into the splitedRawStr array
utils.SplitRawStr("</tr>",rawStr,&splitedRawStr)


}


func Process(symbol string){

 for _,val:=range splitedRawStr{
	
	count:=utils.CalcHowMuchTd(val)
	
  if count>10&&symbol=="nif"{
	// algorithim for making understandable object from the Raw text Nifty One
	  mk_Eq_Sme_Obj(val)
  }

  if count==12&&symbol=="pre"{
	// algorithim for making understandable object from the Raw text Premarket One
	mk_pre_mrkt_obj(val)
  }

 }



 if symbol=="nif"{
	 defer os.WriteFile("./result/niftyStocks.json",makejson.MakeJsonData(equityAndSmeMarketList),0666)
	}else if symbol=="pre"{
		defer os.WriteFile("./result/preMrkt.json",makejson.MakeJsonData(preMarketList),0666)
	}
	
}


