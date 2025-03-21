package powerwall

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/ygelfand/go-powerwall/internal/powerwall/queries"
	"google.golang.org/protobuf/proto"
)

func (p *PowerwallGateway) GetConfig() *string {
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
	resp, err := p.makeTedRequest("POST", "v1", bytes.NewBuffer(reqbody))
	if err != nil {
		log.Println(err)
		return nil
	}
	pr := &ParentMessage{}
	err = proto.Unmarshal(resp, pr)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &pr.Message.Config.GetRecv().File.Text
}

func (p *PowerwallGateway) RunQuery(query string, params *string) *string {
	var reqbody string
	queryObj := queries.GetQuery(query)
	if queryObj == nil {
		log.Printf("Query: %s not found", query)
		return nil
	}
	if params == nil || *params == "" {
		if queryObj.DefaultParams != nil {
			reqbody = *queryObj.DefaultParams
		} else {
			reqbody = "{}"
		}
	} else {
		reqbody = *params
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
						Value: queries.GetQuery(query).Key(),
						Text:  queries.GetQuery(query).GetQuery(),
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
	if err != nil {
		log.Println(err)
		return nil
	}
	resp, err := p.makeTedRequest("POST", "v1", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return nil
	}

	pr := &ParentMessage{}
	err = proto.Unmarshal(resp, pr)
	if err != nil {
		log.Println(err)
		return nil
	}
	if pr.Message.Payload == nil {
		log.Printf("err: %v", pr)
		return nil
	}
	return &pr.Message.Payload.Recv.Text
}

func (p *PowerwallGateway) getDin() *string {
	resp, err := p.makeTedRequest("GET", "din", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	res := string(resp)
	return &res
}

func (p *PowerwallGateway) makeTedRequest(method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, p.Endpoint.JoinPath("tedapi", path).String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/octet-string")
	req.SetBasicAuth("Tesla_Energy_Device", p.password)
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respbody, nil
}
