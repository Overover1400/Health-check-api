package store

import (
	"github.com/overover1400/api-health-check/entity"
	"net/http"
	"time"
)

func (m MySqlStore) FindClientApiToHealthCheck() error {
	var clientUrl []entity.ClientApi
	if err := m.db.Select("Select listener_url,id").Where("subdate(now(), interval health_check_time second) >= date_interval").
		Find(&clientUrl).Error; err != nil {
		return err
	}

	//clientUrlEntities := make([]entity.ClientApi, len(clientUrl))
	//for i := range clientUrl {
	//	clientUrlEntities[i].ID = clientUrl[i]
	//}

	//TODO separate http.get to delivery layer
	if len(clientUrl) != 0 {
		for i := 0; i < len(clientUrl); i++ {
			resp, err := http.Get(clientUrl[i].Url)
			//TODO defer should not use inside "for" loop for resource leak
			defer resp.Body.Close()
			if err != nil {
				return err
			}
			m.db.Model(&ClientApi{}).Where("id = ?", clientUrl[i].ID).
				Update("interval_date=?", time.Now().String()[:19])

			if resp.StatusCode != 200 {

				resp, err = http.Get(clientUrl[i].ListenerUrl)
				//TODO defer should not use inside "for" loop for resource leak
				defer resp.Body.Close()
				if err != nil {
					return err
				}

			}

		}
	}

	return nil
}
