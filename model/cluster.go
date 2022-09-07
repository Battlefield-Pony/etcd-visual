package model

import (
	"etcd-visual/conf"
)

const (
	ClusterTableName         = "cluster"
	ClusterEndpointTableName = "cluster_endpoint"
)

type Cluster struct {
	BaseEntity
	ClusterName      string            `json:"clusterName,omitempty"`
	ClusterEndpoints []ClusterEndpoint `json:"clusterEndpoints,omitempty"`
}

type ClusterEndpoint struct {
	BaseEntity
	Endpoint  string `json:"endpoint,omitempty"`
	ClusterId string `json:"clusterId"`
}

func ListCluster(cluster Cluster) ([]Cluster, error) {
	var clusters []Cluster
	err := conf.SvcCtx.Db.Table(ClusterTableName).Where(cluster).Find(&clusters).Error
	return clusters, err
}

func GetClusterWithAssociation(id int) (Cluster, error) {
	cluster := &Cluster{
		BaseEntity: BaseEntity{
			Id: id,
		},
	}
	err := conf.SvcCtx.Db.Model(&Cluster{}).Preload("ClusterEndpoints").Find(cluster).Error
	return *cluster, err
}
