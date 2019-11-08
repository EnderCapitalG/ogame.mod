package ogame

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getNbrV7(doc *goquery.Document, name string) int {
	val, _ := strconv.Atoi(doc.Find("span."+name+" span").First().AttrOr("data-value", "0"))
	return val
}

func extractFacilitiesFromDocV7(doc *goquery.Document) (Facilities, error) {
	res := Facilities{}
	res.RoboticsFactory = getNbrV7(doc, "roboticsFactory")
	res.Shipyard = getNbrV7(doc, "shipyard")
	res.ResearchLab = getNbrV7(doc, "researchLaboratory")
	res.AllianceDepot = getNbrV7(doc, "allianceDepot")
	res.MissileSilo = getNbrV7(doc, "missileSilo")
	res.NaniteFactory = getNbrV7(doc, "naniteFactory")
	res.Terraformer = getNbrV7(doc, "terraformer")
	res.SpaceDock = getNbrV7(doc, "repairDock")
	res.LunarBase = getNbrV7(doc, "lunarBase")         // TODO: ensure name is correct
	res.SensorPhalanx = getNbrV7(doc, "sensorPhalanx") // TODO: ensure name is correct
	res.JumpGate = getNbrV7(doc, "jumpGate")           // TODO: ensure name is correct
	return res, nil
}

func extractDefenseFromDocV7(doc *goquery.Document) (DefensesInfos, error) {
	res := DefensesInfos{}
	res.RocketLauncher = getNbrV7(doc, "rocketLauncher")
	res.LightLaser = getNbrV7(doc, "laserCannonLight")
	res.HeavyLaser = getNbrV7(doc, "laserCannonHeavy")
	res.GaussCannon = getNbrV7(doc, "gaussCannon")
	res.IonCannon = getNbrV7(doc, "ionCannon")
	res.PlasmaTurret = getNbrV7(doc, "plasmaCannon")
	res.SmallShieldDome = getNbrV7(doc, "shieldDomeSmall")
	res.LargeShieldDome = getNbrV7(doc, "shieldDomeLarge")
	res.AntiBallisticMissiles = getNbrV7(doc, "missileInterceptor")
	res.InterplanetaryMissiles = getNbrV7(doc, "missileInterplanetary")
	return res, nil
}