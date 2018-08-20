/*
Sniperkit-Bot
- Status: analyzed
*/

package e2e

import (
	"os"

	"github.com/sniperkit/snk.fork.klone/pkg/klone"
)

func IdempotentKlone(path, query string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	err := klone.Klone(query)
	if err != nil {
		return err
	}
	return nil
}
