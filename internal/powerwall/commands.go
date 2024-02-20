package powerwall

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ygelfand/go-powerwall/internal/powerwall/queries"
	"google.golang.org/protobuf/proto"
)

type PowerwallGateway struct {
	endpoint   string
	password   string
	httpClient *http.Client
	Din        string
}

func NewPowerwallGateway() *PowerwallGateway {
	pwr := &PowerwallGateway{}
	pwr.httpClient = pwr.getClient()
	val, present := os.LookupEnv("POWERWALL_PASSWORD")
	if !present {
		log.Fatal("Please set POWERWALL_PASSWORD environment variable to full powerwall password")
	}
	pwr.password = val
	val, present = os.LookupEnv("POWERWALL_ENDPOINT")
	if present {
		pwr.endpoint = val
	} else {
		pwr.endpoint = "https://192.168.91.1/tedapi"
	}
	pwr.Din = pwr.getDin()
	return pwr
}

func (p *PowerwallGateway) GetConfig() string {
	pm := &ParentMessage{
		Message: &MessageEnvelope{
			Config: &ConfigType{
				Config: &ConfigType_Send{
					Send: &PayloadConfigSend{
						Num:  1,
						File: "config.json",
					},
				},
			},
			DeliveryChannel: 1,
			Sender: &Participant{
				Id: &Participant_Local{
					Local: 1,
				},
			},
			Recipient: &Participant{
				Id: &Participant_Din{
					Din: p.Din,
				},
			},
		},
		Tail: &Tail{
			Value: 1,
		},
	}
	reqbody, err := proto.Marshal(pm)
	resp, err := p.httpClient.Do(p.getRequest("POST", "v1", bytes.NewBuffer(reqbody)))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	pr := &ParentMessage{}
	proto.Unmarshal(body, pr)
	return pr.Message.Config.GetRecv().File.Text

}

func (p *PowerwallGateway) RunQuery(query string, params interface{}) string {
	var reqbody string
	queryObj := queries.GetQuery(query)
	if queryObj == nil {
		log.Fatalf("Query: %s not found", query)
	}
	if params == nil {
		if queryObj.DefaultParams != nil {
			reqbody = *queryObj.DefaultParams
		} else {
			reqbody = "{}"
		}
	} else {
		obj, err := json.Marshal(params)
		if err != nil {
			log.Fatalln(err)
		}
		reqbody = string(obj)
	}
	pm := &ParentMessage{
		Message: &MessageEnvelope{
			DeliveryChannel: 1,
			Sender: &Participant{
				Id: &Participant_Local{
					Local: 1,
				},
			},
			Recipient: &Participant{
				Id: &Participant_Din{
					Din: p.Din,
				},
			},
			Payload: &QueryType{
				Send: &PayloadQuerySend{
					RequestFormat: Format_Json,
					Signature:     queries.GetQuery(query).Sig(),
					Payload: &PayloadString{
						Value: 1,
						Text:  queries.GetQuery(query).Query,
					},
					Body: &StringValue{
						Value: reqbody,
					},
				},
			},
		},
		Tail: &Tail{
			Value: 1,
		},
	}

	body, err := proto.Marshal(pm)
	resp, err := p.httpClient.Do(p.getRequest("POST", "v1", bytes.NewBuffer(body)))
	if err != nil {
		log.Fatalln(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	pr := &ParentMessage{}
	err = proto.Unmarshal(body, pr)
	if err != nil {
		log.Fatalln(err)
	}
	return pr.Message.Payload.Recv.Text
}
func (p *PowerwallGateway) getDin() string {
	resp, err := p.httpClient.Do(p.getRequest("GET", "din", nil))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func (p *PowerwallGateway) getRequest(method, path string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", p.endpoint, path), body)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-type", "application/octet-string")
	req.SetBasicAuth("Tesla_Energy_Device", p.password)
	return req
}

func (p *PowerwallGateway) getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
