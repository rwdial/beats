package dns

import (
	"github.com/rwdial/beats/libbeat/logp"

	"github.com/rwdial/beats/packetbeat/procs"
	"github.com/rwdial/beats/packetbeat/protos"
)

func (dns *Dns) ParseUdp(pkt *protos.Packet) {
	defer logp.Recover("Dns ParseUdp")

	logp.Debug("dns", "Parsing packet addressed with %s of length %d.",
		pkt.Tuple.String(), len(pkt.Payload))

	dnsPkt, err := decodeDnsData(TransportUdp, pkt.Payload)
	if err != nil {
		// This means that malformed requests or responses are being sent or
		// that someone is attempting to the DNS port for non-DNS traffic. Both
		// are issues that a monitoring system should report.
		logp.Debug("dns", "%s", err.Error())
		return
	}

	dnsTuple := DnsTupleFromIpPort(&pkt.Tuple, TransportUdp, dnsPkt.Id)
	dnsMsg := &DnsMessage{
		Ts:           pkt.Ts,
		Tuple:        pkt.Tuple,
		CmdlineTuple: procs.ProcWatcher.FindProcessesTuple(&pkt.Tuple),
		Data:         dnsPkt,
		Length:       len(pkt.Payload),
	}

	if dnsMsg.Data.Response {
		dns.receivedDnsResponse(&dnsTuple, dnsMsg)
	} else /* Query */ {
		dns.receivedDnsRequest(&dnsTuple, dnsMsg)
	}
}
