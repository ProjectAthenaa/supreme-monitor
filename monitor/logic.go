package monitor

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/prometheus/common/log"
	"strconv"
	"strings"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (tk *Task) iteration() error {
	req, err := tk.NewRequest("GET", "https://www.supremenewyork.com/mobile_stock.json", nil)
	if err != nil {
		return err
	}

	res, err := tk.Do(req)
	if err != nil {
		return err
	}

	stockArray, err := parseMobileStock(res.Body)
	if err != nil{
		return err
	}

	log.Info(tk.keywords)

	outer:
	for _, data := range stockArray.ProductsAndCategories[tk.category]{
		for _, substr := range tk.keywords.Positive{
			if !strings.Contains(strings.ToLower(substr), data.Name){
				continue outer
			}
		}
		for _, substr := range tk.keywords.Negative{
			if strings.Contains(strings.ToLower(substr), data.Name){
				continue outer
			}
		}
		tk.GetProduct(data)
	}

	return nil
}


func (tk *Task) GetProduct(dataIn ProductData) error{
	req, err := tk.NewRequest("GET", fmt.Sprintf("https://www.supremenewyork.com/shop/%d.json", dataIn.ID), nil)
	if err != nil{
		return err
	}
	res, err := tk.Do(req)
	if err != nil{
		return err
	}

	var productList *ProductResponse
	json.Unmarshal(res.Body, &productList)

	for _, data := range productList.Styles{
		for _, size := range data.Sizes{
			if size.StockLevel != 0{
				tk.Monitor.Channel <- map[string]interface{}{
					"color":data.Name,
					"size": size.Name,
					"sizeId": strconv.Itoa(size.ID),
					"pid": dataIn.ID,
				}
			}
		}
	}
	return nil
}

//HD6QH:H2ZP10D3:194.163.222.98:8682
//HD6QH:H2ZP10D3:194.163.222.99:8683
//HD6QH:H2ZP10D3:194.163.222.100:8684
//HD6QH:H2ZP10D3:194.163.222.101:8685
//HD6QH:H2ZP10D3:194.163.222.102:8686