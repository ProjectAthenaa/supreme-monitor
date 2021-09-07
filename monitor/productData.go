package monitor

import "strings"

type MobileStock struct {
	ProductsAndCategories map[string][]ProductData `json:"products_and_categories"`
}

type ProductData struct {
	Name string `json:"name"`
	ID   int `json:"id"`
	Category string `json:"category_name"`
}

func (m *ProductData) containsNegative(negativeKWs []string) bool {
	for _, kw := range negativeKWs {
		if strings.Contains(strings.TrimSpace(strings.ToLower(m.Name)), kw) {
			return true
		}
	}
	return false
}

func (m *ProductData) containsPositive(positiveKWs []string) bool {
	for _, kw := range positiveKWs {
		if strings.Contains(strings.TrimSpace(strings.ToLower(m.Name)), kw) {
			return true
		}
	}
	return false
}

func parseMobileStock(resp []byte) (*MobileStock, error) {
	var mobileStock MobileStock

	if err := json.Unmarshal(resp, &mobileStock); err != nil {
		return nil, err
	}

	return &mobileStock, nil
}

type ProductResponse struct {
	Styles []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Sizes []struct {
			Name       string `json:"name"`
			ID         int    `json:"id"`
			StockLevel int    `json:"stock_level"`
		} `json:"sizes"`
	} `json:"styles"`
}