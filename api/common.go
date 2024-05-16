package api

import (
	"github.com/vrsalazar/pertcpm/models"
)

func AddPhase(phaseNum string, projectId uint, k string, ot, mlt, pt int) uint {
	// log.Printf("Adding Phase: phaseNum=%s, projectID=%d, k=%s, values=%d,%d,%d\n", phaseNum, projectId, k, ot, mlt, pt)
	phase := models.Phase{}
	phase.No = phaseNum
	phase.ProjectID = projectId
	phase.Name = k
	phase.OptimisticTime = ot
	phase.MostLikelyTime = mlt
	phase.PessimisticTime = pt
	phase.Save()
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
