package handler

import "github.com/nporsche/np-golang-logger"
import "net/http"
import "storage"
import "time"
import "util"
import "model"
import "encoding/json"
import "def"
import "strconv"
import "runtime/debug"

//import "io/ioutil"
type reportLocationResponse struct {
	ErrNo   int            `json:"errno"`
	ErrMsg  string         `json:"errmsg"`
	Players []model.Player `json:"players"`
	Beans   []model.Bean   `json:"beans"`
}

func ReportLocation(rw http.ResponseWriter, req *http.Request) {
	errNo := 0
	errMsg := ""
	players := []model.Player{}
	beans := []model.Bean{}
	startTick := time.Now()

	session := util.Session()

	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case *def.BusinessException:
				errNo = x.(*def.BusinessException).ErrNo
				errMsg = x.(*def.BusinessException).ErrMsg
			default:
				errMsg = def.UnExpectedErrMsg
				errNo = def.UnExpectedErrNo
				logger.Errorf("session=[%s] unexpected exception=[%v] stack=[%s]", session, x, string(debug.Stack()))
			}
		}
		response := reportLocationResponse{errNo, errMsg, players, beans}

		encoder := json.NewEncoder(rw)
		encoder.Encode(response)

		endTick := time.Now()
		logger.Infof("Access ReportLocation session=[%s] errno=[%d] errmsg=[%s] duration=[%d] id=[%s] longitude=[%s] latitude=[%s]",
			session,
			errNo,
			errMsg,
			int64(endTick.Sub(startTick)/time.Millisecond),
			req.FormValue("id"),
			req.FormValue("longitude"),
			req.FormValue("latitude"))
	}()

	id, err1 := strconv.ParseUint(req.FormValue("id"), 10, 64)
	longitude, err2 := strconv.ParseFloat(req.FormValue("longitude"), 64)
	latitude, err3 := strconv.ParseFloat(req.FormValue("latitude"), 64)

	if err1 == nil && err2 == nil && err3 == nil {
		storage.Ele.PlayerReport(id, longitude, latitude)
		players = storage.Ele.Players
		beans = storage.Ele.Beans
	} else {
		panic(def.ParamException)
	}
}
