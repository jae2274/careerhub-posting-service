package main

import (
	"log"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgapp"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgrepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/awscfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
	"github.com/jae2274/goutils/cchan"
)

func main() {
	quitChan := make(chan bgapp.QuitSignal)
	processedChan, errChan := initApp().Run(quitChan)

	timeoutQuit := make(chan bgapp.QuitSignal)
	errorQuit := make(chan bgapp.QuitSignal)
	go cchan.Timeout(10*time.Minute, 10*time.Minute, processedChan, timeoutQuit)
	go cchan.TooMuchError(10, 10*time.Minute, errChan, errorQuit)

	select {
	case <-errorQuit:
		close(quitChan)
		log.Fatal("Too much error")
	case <-timeoutQuit:
		close(quitChan)
		log.Fatal("Timeout")
	case <-quitChan:
		close(errorQuit)
		close(timeoutQuit)
		return
	}
}

func initApp() *bgapp.BackgroundApp {
	envVars, err := vars.Variables()
	checkErr(err)

	cfg, err := awscfg.Config()
	checkErr(err)

	jobPostingQ, err := queue.NewSQS(cfg, envVars.SqsEndpoint, envVars.JobPostingQueue)
	checkErr(err)

	companyQ, err := queue.NewSQS(cfg, envVars.SqsEndpoint, envVars.CompanyQueue)
	checkErr(err)

	closedQ, err := queue.NewSQS(cfg, envVars.SqsEndpoint, envVars.ClosedQueue)
	checkErr(err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName)
	checkErr(err)

	jpModel := (&jobposting.JobPostingInfo{})
	jpCol := db.Collection(jpModel.Collection())
	err = mongocfg.CheckIndexViaCollection(jpCol, jpModel.IndexModels())
	checkErr(err)

	return bgapp.NewBackgroundApp(
		bgrepo.NewJobPostingRepo(jpCol),
		companyQ, jobPostingQ, closedQ)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
