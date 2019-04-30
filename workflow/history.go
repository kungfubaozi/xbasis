package workflow

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
)

type history struct {
	conn    redis.Conn
	session *mgo.Session
	client  *indexutils.Client
}

type nodeOperateHistory struct {
	InstanceId string `json:"instance_id"`
	NodeId     string `json:"node_id"`
	Ok         bool   `json:"ok"`
}

//直接从es里查找，速度要快
//gs-flow-node-status-his主要对节点操作进行记录
func (h *history) isCompleted(instanceId, nodeId string) (*gs_commons_dto.State, error) {
	var nh *nodeOperateHistory
	ok, err := h.client.QueryFirst("gs-flow-node-status-his", map[string]interface{}{
		"instanceId": instanceId,
		"nodeId":     nodeId,
	}, nh)
	if err != nil {
		return nil, err
	}
	if ok && nh.Ok {
		return errstate.Success, nil
	}
	return ErrInvalid, nil
}

func (h *history) getAllInstanceStatusHistory(instanceId string) ([]*nodeOperateHistory, error) {

}

func (h *history) finished(instanceId, nodeId string) (*gs_commons_dto.State, error) {

}
