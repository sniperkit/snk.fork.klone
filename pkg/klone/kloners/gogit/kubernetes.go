/*
Sniperkit-Bot
- Status: analyzed
*/

package gogit

import (
	"fmt"

	"github.com/sniperkit/snk.fork.klone/pkg/provider"
)

// Kubernetes does not follow the traditional path logic, so we have to hard code it
func repoToKubernetesPath(repo provider.Repo) string {
	path := fmt.Sprintf("%s/src/%s/%s", Gopath(), "k8s.io", repo.Name())
	return path
}
