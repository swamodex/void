package void

import (
	"fmt"
	"void/upgrades"
	v1 "void/upgrades/v1"

	upgradetypes "cosmossdk.io/x/upgrade/types"
)

// Upgrades list of chain upgrades
var Upgrades = []upgrades.Upgrade{v1.Upgrade}

// RegisterUpgradeHandlers registers the chain upgrade handlers
func (app *VoidApp) RegisterUpgradeHandlers() {
	// register all upgrade handlers
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.ModuleManager,
				app.Configurator(),
			),
		)
	}

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	// register store loader for current upgrade
	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
			break
		}
	}
}
