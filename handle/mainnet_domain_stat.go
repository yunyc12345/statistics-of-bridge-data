package handle

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"yunyc12345/statistics-of-bridge-data/utils"
)

func TriggerDomainStat(w *sync.WaitGroup, statDomains []utils.StatDomain, filePath, dataStr, fileDataStr string, hisList map[string]*sync.Map) {
	fp := filePath + "/" + dataStr + "/stat_domain"
	err := utils.CreateCsvDir(fp)
	if err != nil {
		panic(err)
	}

	w.Add(1)
	defer w.Done()
	utils.Logger.Infof("stat nft bridge sender handle start")

	wg := &sync.WaitGroup{}
	for _, domain := range statDomains {
		wg.Add(1)
		go func(d utils.StatDomain) {
			defer wg.Done()

			list := &sync.Map{}

			if hisList != nil && len(hisList) != 0 && hisList[d.Ty] != nil {
				list = hisList[d.Ty]
			}

			StatDomainHandler(d, list)

			name := d.Ty + "-" + "stat_domain" + "-" + strconv.FormatInt(d.EndTimestamp, 10) + "-" + fileDataStr
			ll, err := utils.ToCsv(list, fp, name)
			if err != nil {
				utils.Logger.Errorf("stat domain to csv err: %v", err)
			}
			utils.Logger.Infof("domain ty: %v, new domain stat list len: %v", d.Ty, ll)

		}(domain)
	}
}

func StatDomainHandler(domain utils.StatDomain, list *sync.Map) {
	j, _ := json.Marshal(domain)
	utils.Logger.Infof("domain: %v", string(j))
	//ctx := context.Background()

	uri := domain.Url + "?start=" + strconv.FormatInt(domain.StartTimestamp, 10) + "&end=" + strconv.FormatInt(domain.EndTimestamp, 10) + "&domain_type=" + domain.Ty

	utils.Logger.Infof("domain: %v, uri: %v", domain.Ty, uri)
	res, err := http.Get(uri)
	if err != nil {
		utils.Logger.Errorf("get domain url: %v, err: %v", uri, err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	var sdr utils.StateDomainResp
	err = json.Unmarshal(body, &sdr)
	if err != nil {
		utils.Logger.Errorln(err)
	}

	//utils.Logger.Infof("domain: %v, res: %v", domain.Ty, sdr)

	for _, datum := range sdr.Data {
		list.Store(strings.ToLower(strings.TrimSpace(datum.Address)), utils.Member)
		if strings.Contains(datum.SenderAddress, "0x") {
			list.Store(strings.ToLower(strings.TrimSpace(datum.SenderAddress)), utils.Member)
		}
	}

}
