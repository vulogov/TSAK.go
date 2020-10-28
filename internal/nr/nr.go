package nr

import (
  "bytes"
  "compress/gzip"
  "net/http"
)

func Logs(nrikey string, url string, compress bool, _payload []byte) bool {
  var payload []byte
  var b bytes.Buffer
  if compress {
    w := gzip.NewWriter(&b)
    w.Write([]byte(_payload))
    w.Close()
    payload = []byte(b.Bytes())
  } else {
    payload = []byte(_payload)
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
  if err != nil {
    return false
  }
  req.Header.Set("Api-Key", nrikey)
  if compress {
    req.Header.Set("Content-Type", "application/gzip")
    req.Header.Set("Content-Encoding", "gzip")
  } else {
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Content-Encoding", "json")
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  return true
}
