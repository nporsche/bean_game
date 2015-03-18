package handler

import "github.com/nporsche/np-golang-logger"
import "net/http"
import "time"
import "encoding/json"
import "strconv"
import "runtime/debug"

//import "io/ioutil"
type reportLocationResponse struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}

func FetchMsgs(rw http.ResponseWriter, req *http.Request) {
	errNo := 0
	errMsg := ""
	startTick := time.Now()

	session := util.Session()

	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case *def.BusinessException:
				errNo = x.(*def.BusinessException).ErrNo
				errMsg = x.(*def.BusinessException).ErrMsg
			default:
				errMsg = def.UnexpectedInternalErrMsg
				errNo = def.UnexpectedInternalErrNo
				logger.Errorf("session=[%s] unexpected exception=[%v] stack=[%s]", session, x, string(debug.Stack()))
			}
		}
		response := reportLocationResponse{errNo, errMsg}

		encoder := json.NewEncoder(rw)
		encoder.Encode(fetchResponse)

		endTick := time.Now()
		logger.Infof("Access FetchMsgs session=[%s] errno=[%d] errmsg=[%s] duration=[%d] token=[%s] phone=[%s] city_id=[%s] begin=[%s] count=[%s]",
			session,
			errNo,
			errMsg,
			int64(endTick.Sub(startTick)/time.Millisecond),
			req.FormValue("token"),
			req.FormValue("phone"),
			req.FormValue("city_id"),
			req.FormValue("begin"),
			req.FormValue("count"))
	}()

	if !util.ValidateTicket(session, req.FormValue("token")) {
		panic(def.TokenException)
	}

	begin, err1 := strconv.Atoi(req.FormValue("begin"))
	count, err2 := strconv.Atoi(req.FormValue("count"))
	phone, err3 := util.DecodeCellphone(req.FormValue("phone"))
	cityId, err4 := strconv.Atoi(req.FormValue("city_id"))

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		msgs = db.FetchMsgs(session, phone, cityId, begin, count)
	} else {
		panic(def.ParamException)
	}
}
