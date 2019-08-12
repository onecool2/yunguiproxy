package main

import (
        "fmt"
	"log"
	"time"
	"context"
        //"gopkg.in/mgo.v2"
        //"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"
	"math/big"
	"io/ioutil"
	"crypto/ecdsa"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	yunguiContract "github.com/onecool2/yunguiProxy/contract"
)

type Person struct {
        Name string
        Phone string
}
/*
type Orders struct {
	Order_no	 string
	User_id		 uint64
	Amount		 float
	Account_amount	 float
	Thirdpay_amount	 float
	Refund_amount	 float
	Coupon_amount	 float
	User_coupon_id	 uint64
	Bank_type	 string
	Pay_status	 uint8
	Remark		 
	Success_time
	Create_time
	Update_time
	User_name
	Mobile
	Fridge_id
	Partner_id
	Calc_fail
	Order_pay_no
	Open_door_time
	Open_door_weight
	Close_door_weight
	Last_close_door_weight
	Wmdiscount
	Nodiscountam
	Diff_door_weight
	Has_zero_layer
	Origin_amount
	Before_discount_amount

}
*/

type Fridge struct {
	Order_no		string
	Partner_id		uint64
	Business_area_id	uint64
	Business_district_id	uint64
	Warehouse_id		uint64
	Fridge_id		uint64
	Real_type		int
	Order_type		int
	Status			int
	Export_status		int
	Create_user_id		uint64
	Warehouse_order_no	string
	Create_time		string
	Update_time		string
}

type Transaction struct {
	_id string
	Hash string
	BlockHash string
	BlockNumber string
	CumulativeGasUsed string
	Fn_name string
	From string
	Gas string
	GasPrice string
	GasUsed string
	Input string
	Log string
	LogBloom string
	Nonce string
	R string
	S string
	Ss string
	To string
	TransactionIndex string
	V string
	Value string
}
type RetMsg struct {
	code    int    //400
	message string //"具体错误信息",
	data    string //"data":null
}

type SenderBuffer struct {
	Table	string
	Ss	string
}


const (
	RPC_HOST           string = "http://39.100.153.59:32000"
	WS_HOST            string = "ws://127.0.0.1:32002"
	CONTRACT_ADDRESS   string = "0xe751788bc2ce3cd885f027e22190a11203893659"
	OWNER_PUBLIC_KEY   string = ""
	OWNER_PRIVATE_KEY  string = "41177460dc4d1760832fcbfcaae19e3f734c968edafd00e2947e8bee3da59801"
	EMPTY              int    = 0
	READY              int    = 1
	SENT               int    = 2
	MAX_ON_CHAIN_DELAY        = "-30s"
)

var (
	EthClient  *ethclient.Client
	Contract   *yunguiContract.YunguiContract
	privateKey *ecdsa.PrivateKey
	publicKey  common.Address
	owner      common.Address
	SendChan   chan	SenderBuffer 	
)


func init() {
	var contractAddress common.Address
	var err error
	EthClient, err = ethclient.Dial(RPC_HOST)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress = common.HexToAddress(CONTRACT_ADDRESS)
	Contract, err = yunguiContract.NewYunguiContract(contractAddress, EthClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = crypto.HexToECDSA(OWNER_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	owner = crypto.PubkeyToAddress(*publicKeyECDSA)
	SendChan = make (chan SenderBuffer, 500)
	go LoopAndSendTx()
}

func LoopAndSendTx() {
	var tx *types.Transaction 
	var err error
	var count int

	limiter := time.Tick(time.Duration(1)*time.Second)
	nonce, err := EthClient.PendingNonceAt(context.Background(), owner)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case   <-limiter:
			count = 0
			fmt.Println("clear timer limiter")
		case  senderBuffer := <-SendChan:
			auth := bind.NewKeyedTransactor(privateKey)
	    	 	auth.Nonce = big.NewInt(int64(nonce))
	    		auth.Value = big.NewInt(0)     // in wei
	     		auth.GasLimit = uint64(600000) // in units
			if senderBuffer.Table == "Building" {
				tx, err = Contract.Building(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "Business_district" {
				tx, err = Contract.BusinessDistrict(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "Fridge_storage" {	
				tx, err = Contract.FridgeStorage(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "Orders" {	
				tx, err = Contract.Orders(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "Fridge" {	
				tx, err = Contract.Fridge(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "Business_area" {	
				tx, err = Contract.BusinessArea(auth, senderBuffer.Ss)
			}else if senderBuffer.Table == "City_info" {	
				tx, err = Contract.CityInfo(auth, senderBuffer.Ss)
			}
		        fmt.Println("table: ss: nonce:", senderBuffer.Table, senderBuffer.Ss, nonce)
		        if err != nil {
			    glog.Warning("send tx:", err)
		            nonce, err = EthClient.PendingNonceAt(context.Background(), owner)
			    if err != nil {
			        log.Fatal("!!!!!!!!!!!!!!!:", err)
			    }
			}else {
			    nonce++
		            fmt.Println("###########:", tx.Hash(), nonce)
			    count++
				if count > 20 {
				    time.Sleep(time.Duration(1)*time.Second)
				    count = 0
				}
			}
		}
	}
}


func setBuilding(Data map[string]interface{}) {
	fmt.Println("setBuilding:", Data)

	var err error
	tableType := Data["type"]
	var senderBuffer SenderBuffer	

	if (tableType == "building"){
		id := int64(Data["id"].(float64))
		name := Data["name"].(string)
		business_district_id := int64(Data["business_district_id"].(float64))
		warehouse_id := int64(Data["warehouse_id"].(float64))
		status := int64(Data["status"].(float64))
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))

		ss := fmt.Sprintf("id|%d^name|%s^business_district_id|%d^warehouse_id|%d^status|%d^create_time|%d^update_time|%d", id, name, business_district_id, warehouse_id, status, create_time, update_time)
		senderBuffer.Ss = ss
	        senderBuffer.Table = "Building"
		SendChan <- senderBuffer

		if err != nil {
			glog.Warning("send tx:", err)
		}
	}else{
		fmt.Println("tableType", tableType)
	}
}

func setBusiness_district(Data map[string]interface{}) {
	fmt.Println("setBusiness_district:", Data)

	tableType := Data["type"]
	var senderBuffer SenderBuffer	

	if (tableType == "business_district"){
		id := int64(Data["id"].(float64))
		name := Data["name"].(string)
		business_area_id := int64(Data["business_area_id"].(float64))
		status := int64(Data["status"].(float64))
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))

		ss := fmt.Sprintf("id|%d^name|%s^business_area_id|%d^status|%d^create_time|%d^update_time|%d", id, name, business_area_id, status, create_time, update_time)
		senderBuffer.Ss = ss
	        senderBuffer.Table = "Business_district"
		SendChan <- senderBuffer

	}else{
		fmt.Println("tableType", tableType)
	}
}

func setFridge_storage(Data map[string]interface{}) {
	fmt.Println("setFridge_storage:", Data)
	tableType := Data["type"]

	var senderBuffer SenderBuffer	
	if (tableType == "fridge_storage"){
		fridge_id := int64(Data["fridge_id"].(float64))
		partner_id := int64(Data["partner_id"].(float64))
		warehouse_id := int64(Data["warehouse_id"].(float64))
		goods_id := int64(Data["goods_id"].(float64))
		goods_name := Data["goods_name"].(string)
		goods_code := Data["goods_code"].(string)
		quantity := int64(Data["quantity"].(float64))
		weight :=  Data["weight"].(float64)
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))

		ss := fmt.Sprintf("fridge_id|%d^partner_id|%d^warehouse_id|%d^goods_id|%d^goods_name|%s^goods_code|%s^quantity|%d^weight|%f^create_time|%d^update_time|%d", fridge_id, partner_id, warehouse_id, goods_id, goods_name, goods_code, quantity, weight, create_time, update_time)
		senderBuffer.Ss = ss
	        senderBuffer.Table = "Fridge_storage"
		SendChan <- senderBuffer
	} else{
		fmt.Println("tableType", tableType)
	}
}

func setOrders(Data map[string]interface{}) {
	fmt.Println("setOrders:", Data)

	var senderBuffer SenderBuffer	
	tableType := Data["type"]

	if (tableType == "orders"){
		order_no := Data["order_no"].(string)
		user_id := int64(Data["user_id"].(float64))
		amount := Data["amount"].(float64)
		account_amount := Data["account_amount"].(float64)
		thirdpay_amount := Data["thirdpay_amount"].(float64)
		refund_amount := Data["refund_amount"].(float64)
		coupon_amount := Data["coupon_amount"].(float64)
		user_coupon_id := int64(Data["user_coupon_id"].(float64))
		bank_type := Data["bank_type"].(string)
		pay_status := int64(Data["pay_status"].(float64))
		remark := Data["remark"].(string)
		success_time := int64(Data["success_time"].(float64))
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))
		user_name := Data["user_name"].(string)
		mobile := Data["mobile"].(string)
		fridge_id := int64(Data["fridge_id"].(float64))
		partner_id := int64(Data["partner_id"].(float64))
		calc_fail := int64(Data["calc_fail"].(float64))
		order_pay_no := Data["order_pay_no"].(string)
		open_door_time := int64(Data["open_door_time"].(float64))
		open_door_weight := Data["open_door_weight"].(float64)
		close_door_weight := Data["close_door_weight"].(float64)
		last_close_door_weight := Data["last_close_door_weight"].(float64)
		wmdiscount := Data["wmdiscount"].(float64)
		nodiscountam := Data["nodiscountam"].(float64)
		diff_door_weight := Data["diff_door_weight"].(float64)
		has_zero_layer := int64(Data["has_zero_layer"].(float64))
		origin_amount := Data["origin_amount"].(float64)

		ss := fmt.Sprintf("order_no|%s^user_id|%d^amount|%f^account_amount|%f^thirdpay_amount|%f^refund_amount|%f^coupon_amount|%f^user_coupon_id|%d^bank_type|%s^pay_status|%d^remark|%s^success_time|%d^create_time|%d^update_time|%d^user_name|%s^mobile|%s^fridge_id|%d^partner_id|%d^calc_fail|%d^order_pay_no|%s^open_door_time|%d^open_door_weight|%f^close_door_weight|%f^last_close_door_weight|%f^wmdiscount|%f^nodiscountam|%f^diff_door_weight|%f^has_zero_layer|%d^origin_amount|%f", order_no,user_id,amount,account_amount,thirdpay_amount,refund_amount,coupon_amount,user_coupon_id,bank_type,pay_status,remark,success_time,create_time,update_time,user_name,mobile,fridge_id,partner_id,calc_fail,order_pay_no,open_door_time,open_door_weight,close_door_weight,last_close_door_weight,wmdiscount,nodiscountam,diff_door_weight, has_zero_layer, origin_amount)
		senderBuffer.Ss = ss
	        senderBuffer.Table = "Orders"
		SendChan <- senderBuffer
	}else{
		fmt.Println("tableType", tableType)
	}

}
func setFridge(Data map[string]interface{}) {

	fmt.Println("fridge:", Data)
	tableType := Data["type"]
	var senderBuffer SenderBuffer	

	if (tableType == "fridge"){
		id := int64(Data["id"].(float64))
		partner_id := int64(Data["partner_id"].(float64))
		terminal_no := Data["terminal_no"].(string)
		name := Data["name"].(string)
		fridge_no := Data["fridge_no"].(string)

		business_area_id := int64(Data["business_area_id"].(float64))
		business_district_id := int64(Data["business_district_id"].(float64))
		warehouse_id := int64(Data["warehouse_id"].(float64))
		building_id := int64(Data["building_id"].(float64))

		goods_table_id := int64(Data["goods_table_id"].(float64))
		build_goods_table_status := int64(Data["build_goods_table_status"].(float64))
		build_goods_table_time := int64(Data["build_goods_table_time"].(float64))
		max_weight := Data["max_weight"].(float64)
		address := Data["address"].(string)
		status := int64(Data["status"].(float64))
		
		create_time := uint64(Data["create_time"].(float64))
		update_time := uint64(Data["update_time"].(float64))
		generate_order_interval := int64(Data["generate_order_interval"].(float64))
		dot_type := int64(Data["dot_type"].(float64))
		grade := Data["grade"].(string)
		user_count := int64(Data["user_count"].(float64))
		has_sell_data := int64(Data["has_sell_data"].(float64))
		dec_weight_error := Data["dec_weight_error"].(float64)
		inc_weight_error := Data["inc_weight_error"].(float64)


		ss := fmt.Sprintf("id|%d^partner_id|%d^terminal_no|%s^name|%s^fridge_no|%s^business_area_id|%d^business_district_id|%d^warehouse_id|%d^building_id|%d^goods_table_id|%d^build_goods_table_status|%d^build_goods_table_time|%d^max_weight|%f^address|%s^status|%d^Create_time|%d^Update_time|%d^generate_order_interval|%d^dot_type|%d^grade|%s^user_count|%d^has_sell_data|%d^dec_weight_error|%f^inc_weight_error|%f", id,partner_id,terminal_no,name,fridge_no,business_area_id,business_district_id,warehouse_id,building_id,goods_table_id,build_goods_table_status,build_goods_table_time,max_weight,address,status,create_time,update_time,generate_order_interval,dot_type,grade,user_count,has_sell_data,dec_weight_error,inc_weight_error)
		senderBuffer.Ss = ss
	        senderBuffer.Table = "Fridge"
		SendChan <- senderBuffer
	}else{
		fmt.Println("tableType", tableType)
	}
}

func setBusiness_area(Data map[string]interface{}) {

	fmt.Println("business_area:", Data)
	var senderBuffer SenderBuffer	
	tableType := Data["type"]

	if (tableType == "business_area"){
		id := int64(Data["id"].(float64))
		name := Data["name"].(string)
		province_code := Data["province_code"].(string)
		city_code := Data["city_code"].(string)
		partner_id := int64(Data["partner_id"].(float64))
		status := int64(Data["status"].(float64))
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))

		ss := fmt.Sprintf("id|%d^name|%s^province_code|%s^city_code|%s^partner_id|%d^status|%d^create_time|%d^update_time|%d", id, name, province_code, city_code, partner_id, status, create_time, update_time)

		senderBuffer.Ss = ss
	        senderBuffer.Table = "Business_area"
		SendChan <- senderBuffer
	}else{
		fmt.Println("tableType", tableType)
	}
}

func setCity_info(Data map[string]interface{}) {

	fmt.Println("city_info:", Data)
	var senderBuffer SenderBuffer	
	tableType := Data["type"]

	if (tableType == "city_info"){
		id := int64(Data["id"].(float64))
		city_code := Data["city_code"].(string)
		city_name := Data["city_name"].(string)
		status := int64(Data["status"].(float64))
		create_time := int64(Data["create_time"].(float64))
		update_time := int64(Data["update_time"].(float64))

		ss := fmt.Sprintf("id|%d^city_code|%s^city_name|%s^status|%d^create_time|%d^update_time|%d", id, city_code, city_name, status, create_time, update_time)

		senderBuffer.Ss = ss
	        senderBuffer.Table = "City_info"
		SendChan <- senderBuffer
	}else{
		fmt.Println("tableType", tableType)
	}
}

func replyMsg(writer http.ResponseWriter, code int, data string) {

	if bs, err := json.Marshal(code); err == nil {
		writer.Header().Set("Content-type", "application/x-www-form-urlencoded")
		writer.WriteHeader(200)
		writer.Write([]byte(bs))

		fmt.Println(bs)
	}

}

func handlerHttp(writer http.ResponseWriter, request *http.Request){
	body, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	body_str := string(body)
	fmt.Println("--------------------------------------")
	fmt.Println("JSON:", body_str)
	fmt.Println("--------------------------------------")
      
/**************************************************/
	var reqBody map[string]interface{}
	if err := json.Unmarshal(body, &reqBody); err == nil {
		fmt.Println("type:", reqBody["type"])
		if (reqBody["type"] == "fridge"){
		    setFridge(reqBody)
		}else if(reqBody["type"] == "orders"){
		    setOrders(reqBody)
		}else if(reqBody["type"] == "fridge_storage"){
		    setFridge_storage(reqBody)
		}else if(reqBody["type"] == "city_info"){
		    setCity_info(reqBody)
		}else if(reqBody["type"] == "business_area"){
		    setBusiness_area(reqBody)
		}else if(reqBody["type"] == "business_district"){
		    setBusiness_district(reqBody)
		}else if(reqBody["type"] == "building"){
		    setBuilding(reqBody)
		}else{
	            fmt.Println("can't find any path:", reqBody["type"])
		    replyMsg(writer, 400, err.Error())
		}
	}else{
		fmt.Println("Unmarshal:", err)
			replyMsg(writer, 400, err.Error())
	}

	fmt.Println("end handlerHttp")
}

func main() {
/************************** connect mongoDB ***************************************/
/*
        session, err := mgo.Dial("localhost:32017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("quorum").C("transactions")
        //err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//               &Person{"Cla", "+55 53 8402 8510"})
        //if err != nil {
        //        log.Fatal(err)
        //}

        result := []Transaction{}
        err = c.Find(bson.M{"to": "0x2d186eccb98f157669dc6a0bf2f328df318c363a"}).All(&result)
        if err != nil {
                log.Fatal(err)
        }

	fmt.Println("len", len(result))
	for i:=0; i< len(result); i++ {
		fmt.Println("ss:", result[i].Ss)
	}
*/
/*************************************************************************/

/************************* start http server *****************************************/
	r := mux.NewRouter()

	r.HandleFunc("/yungui/fridge", handlerHttp)
	http.Handle("/yungui/fridge", r)

	r.HandleFunc("/yungui/orders", handlerHttp)
	http.Handle("/yungui/orders", r)

	r.HandleFunc("/yungui/fridge_storage", handlerHttp)
	http.Handle("/yungui/fridge_storage", r)
	
	r.HandleFunc("/yungui/city_info", handlerHttp)
	http.Handle("/yungui/city_info", r)

	r.HandleFunc("/yungui/business_area", handlerHttp)
	http.Handle("/yungui/business_area", r)

	r.HandleFunc("/yungui/business_district", handlerHttp)
	http.Handle("/yungui/business_district", r)
	
	r.HandleFunc("/yungui/building", handlerHttp)
	http.Handle("/yungui/building", r)


	fmt.Println("Side Car proxy start listening on 3000...")
	http.ListenAndServe(":3000", r)
/*************************************************************************/
}
