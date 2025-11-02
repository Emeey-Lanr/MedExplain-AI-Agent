package config

import (
	"ai-agent/models"
	"sync"
)

type HistoryDB struct {
	Mu sync.Mutex
	History map[string][]models.ContentData
}
var History  *HistoryDB

func Init (){
	 History = &HistoryDB{History: make(map[string][]models.ContentData)}

}

func (d *HistoryDB) AddHistory (ctxid string, role string, text string){
   d.Mu.Lock()
   defer d.Mu.Unlock()
   d.History[ctxid] = append(d.History[ctxid], models.ContentData{Role: role, Parts:append([]models.TextData{}, models.TextData{Text: text}) } )
}

func (d *HistoryDB) GetHistory (ctxid string) []models.ContentData{
d.Mu.Lock()
defer d.Mu.Unlock()
return d.History[ctxid]
}