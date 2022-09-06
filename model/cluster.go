package model

import "etcd-visual/conf"

type Cluster struct {
	BaseEntity
	ClusterName string            `json:"clusterName,omitempty"`
	Endpoints   []ClusterEndpoint `json:"endpoints,omitempty"`
}

type ClusterEndpoint struct {
	BaseEntity
	Endpoint  string `json:"endpoint,omitempty"`
	ClusterId string `json:"clusterId"`
}

func ListCluster(cluster Cluster) ([]Cluster, error) {
	var clusters []Cluster
	err := conf.SvcCtx.Db.Table("Cluster").Where(cluster).Find(&clusters).Error
	return clusters, err
}

func GetClusterWithAssociation(id string) (Cluster, error) {
	cluster := &Cluster{
		BaseEntity: BaseEntity{
			Id: id,
		},
	}
	err := conf.SvcCtx.Db.Table("Cluster").Preload("ClusterEndpoint").Find(cluster).Error
	return *cluster, err
}
