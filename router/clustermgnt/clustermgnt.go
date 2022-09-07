package clustermgnt

import (
	"etcd-visual/router"
	"net/http"
)

var Routes = router.ApiGroup{
	Prefix: "/cluster",
	Routes: []router.Route{
		{
			Path:        "/",
			Method:      http.MethodGet,
			HandlerFunc: list,
		},
		{
			Path:        "/:id",
			Method:      http.MethodGet,
			HandlerFunc: detail,
		},
	},
}
