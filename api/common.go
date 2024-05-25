package api

import (
	"fmt"
	"strconv"

	"github.com/vrsalazar/pertcpm/models"
)

func AddPhase(phaseNum string, projectId uint, k string, ot, mlt, pt int, workName, workRate, workQty, equipName, equipRate, equipQty []interface{}) uint {
	// log.Printf("Adding Phase: phaseNum=%s, projectID=%d, k=%s, values=%d,%d,%d\n", phaseNum, projectId, k, ot, mlt, pt)
	phase := models.Phase{}
	phase.No = phaseNum
	phase.ProjectID = projectId
	phase.Name = k
	phase.OptimisticTime = ot
	phase.MostLikelyTime = mlt
	phase.PessimisticTime = pt
	phase.Save()
	AddWorker(projectId, phase.ID, workName, workRate, workQty)
	fmt.Println("worker names: ", workName)
	fmt.Println("worker rates: ", workRate)
	fmt.Println("worker quantity: ", workQty)
	AddEquipment(projectId, phase.ID, equipName, equipRate, equipQty)
	fmt.Println("equip names: ", equipName)
	fmt.Println("equip cost: ", equipRate)
	fmt.Println("equip quantity: ", equipQty)
	return phase.ID
}

func AddActivity(actNum string, projectId, phaseId uint, k string, ot, mlt, pt int) uint {
	// log.Printf("Adding Activity: actNum=%s, projectID=%d, phaseID=%d, k=%s, values=%d,%d,%d\n", actNum, projectId, phaseId, k, ot, mlt, pt)
	activity := models.Activity{}
	activity.No = actNum
	activity.Name = k
	activity.OptimisticTime = ot
	activity.MostLikelyTime = mlt
	activity.PessimisticTime = pt
	activity.ProjectID = projectId
	activity.PhaseID = phaseId
	activity.Save()
	return activity.ID
}

func AddSubActivity(subActNum string, projectId, actId, phaseId uint, k string, ot, mlt, pt int) {
	// log.Printf("Adding SubActivity: subActNum=%s, projectID=%d, activityID=%d, phaseID=%d, k=%s, values=%d,%d,%d\n", subActNum, projectId, actId, phaseId, k, ot, mlt, pt)
	subActivity := models.SubActivity{}
	subActivity.No = subActNum
	subActivity.Name = k
	subActivity.OptimisticTime = ot
	subActivity.MostLikelyTime = mlt
	subActivity.PessimisticTime = pt
	subActivity.ProjectID = projectId
	subActivity.PhaseID = phaseId
	subActivity.ActivityID = actId
	subActivity.Save()
}

func AddWorker(projectId, phaseId uint, workName, workRate, workQty []interface{}) {
	for i := range workName {
		worker := models.Worker{}
		worker.ProjectID = projectId
		worker.PhaseID = phaseId
		worker.Name = string(workName[i].(string))
		worker.Rate, _ = strconv.Atoi(workRate[i].(string))
		worker.Quantity, _ = strconv.Atoi(workQty[i].(string))
		worker.Save()
	}
}

func AddEquipment(projectId, phaseId uint, equipmentName, equipmentRate, equipmentQty []interface{}) {
	for i := range equipmentName {
		equipment := models.Equipment{}
		equipment.ProjectID = projectId
		equipment.PhaseID = phaseId
		equipment.Name = string(equipmentName[i].(string))
		equipment.Cost, _ = strconv.Atoi(equipmentRate[i].(string))
		equipment.Quantity, _ = strconv.Atoi(equipmentQty[i].(string))
		equipment.Save()
	}
}
