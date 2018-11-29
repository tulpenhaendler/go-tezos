package goTezos

import (
	"encoding/json"
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"time"
)

func (this *GoTezos) GetPeers() ([]models.Peer, error) {
	bres, err := this.GetResponse("/network/peers/", "")
	tmp := [][]json.RawMessage{}

	json.Unmarshal(bres.Bytes, &tmp)

	res := []models.Peer{}
	for _, a := range tmp {
		peer := models.Peer{}
		json.Unmarshal(a[1], &peer)
		peer.Addr = peer.ReachableAt.Addr
		peer.Port = peer.ReachableAt.Port
		peer.TotalSent = peer.Stat.TotalSent
		peer.TotalRecv = peer.Stat.TotalRecv
		peer.CurrentInflow = peer.Stat.CurrentInflow
		peer.CurrentOutflow = peer.Stat.CurrentOutflow

		if peer.State == "running" {
			peer.LastRunning = time.Now()
		}
		res = append(res, peer)
	}

	if err != nil {
		return res, nil
	}

	return res, nil
}
