package service

//
//import (
//	"online_judge/dao"
//	"online_judge/model"
//)
//
//func AddTestcase(t model.Testcase) (err error) {
//	err = dao.InsertTestcase(t)
//	return
//}
//
//func SearchTestcase(uid string, pid string) (testcases []model.Testcase, err error) {
//	testcases, err = dao.SearchTestcase(uid, pid)
//	return
//}
//
//func UpdateTestcase(t model.Testcase) (err error) {
//	err = dao.UpdateTestcase(t)
//	return
//}
//
//func DeleteTestcase(uid string, tid string) (err error) {
//	err = dao.DeleteTestcase(uid, tid)
//	return
//}
//
//func SearchPendingCode() (submissions []model.Submission, err error) {
//	submissions, err = dao.GetPendingCode()
//	return
//}
//
//func SearchTestcasesByPid(pid string) (testcases []model.Testcase, err error) {
//	testcases, err = dao.GetTestcasesByPid(pid)
//	return
//}
//
//func UpdateStatus(status string, sid string) (err error) {
//	err = dao.UpdateStatus(status, sid)
//	return
//}
