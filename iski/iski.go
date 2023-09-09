package iski

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Shortages struct {
	Data []struct {
	  IlceKodu             string   `json:"ilceKodu"`
	  IlceAdi              string   `json:"ilceAdi"`
	  ArizaAdedi           int      `json:"arizaAdedi"`
	  EtkilenenMahalleAdedi string   `json:"etkilenenMahalleAdedi"`
	  Detail            []struct {
		  ArizaNo              string `json:"arizaNo"`
		  IlceKodu             string `json:"ilceKodu"`
		  IlceAdi              string `json:"ilceAdi"`
		  MahalleKodu          string `json:"mahalleKodu"`
		  MahalleAdi           string `json:"mahalleAdi"`
		  ArizaNeviAciklamasi  string `json:"arizaNeviAciklamasi"`
		  BaslamaTarihi        string `json:"baslamaTarihi"`
		  TahminiBitisTarihi   string `json:"tahminiBitisTarihi"`
	  } `json:"details"`
  } `json:"data"`
}

type District struct {
	Data []struct {
        ID         int    `json:"id"`
        Attributes struct {
            Name       string `json:"title"`
            Code    string `json:"ilceKodu"`
        } `json:"attributes"`
    } `json:"data"`
}

const (
	ISKI_ARIZA_API_URL = "https://iskiapi.iski.gov.tr/api/iski/arizakesinti";
	ISKI_ILCE_API_URL = "https://iskiapi.iski.istanbul/api/ilceler?pagination[limit]=100";
)

func GetAllShortage() Shortages {
	
	response, err := http.Get(ISKI_ARIZA_API_URL);
	if err != nil {
		fmt.Println("Error making incidents API request:", err)
	}

	defer response.Body.Close();
	body, _ := io.ReadAll(response.Body)

	var incidentsResponse Shortages
	json.Unmarshal(body, &incidentsResponse)

	return incidentsResponse;
}

func GetDistiricts() {
	response , err := http.Get(ISKI_ILCE_API_URL);


	if err != nil {
		fmt.Println("Error making API request:", err)
	}

	defer response.Body.Close();


	body, _ := io.ReadAll(response.Body)

	var districtApiResp District
	json.Unmarshal(body, &districtApiResp)
}