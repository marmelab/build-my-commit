package hookhandler

import "io"
import "io/ioutil"
import "net/http"
import "log"

func returnError(responseWriter http.ResponseWriter, status int, msg string) {
    responseWriter.WriteHeader(http.StatusBadRequest)
    responseWriter.Write([]byte(msg))
    return
}

func HookHandler(responseWriter http.ResponseWriter, request *http.Request){
    if request.Method != "POST" {
        returnError(responseWriter, 400, "Invalid method")
         return
    }

    body, err := ioutil.ReadAll(request.Body)

    if err != nil {
        returnError(responseWriter, 500, "")
    }

    isValid, err := ParsePayload(body)

    if err != nil {
        returnError(responseWriter, 400, "Invalid method")
    }

    if isValid {
        log.Println("Request should be parsed")
    } else {
        log.Println("Request should not be parsed")
    }

    io.WriteString(responseWriter, "foo");
}
