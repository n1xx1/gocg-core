package ocgcore

import (
	"math/rand"
	"ocgcore/lib"
	"sync"
)

type OcgDuel struct {
	handle lib.Duel

	messageCh  chan Message
	incomingCh chan []byte

	aliveLock sync.Mutex
}

func newDuel(d lib.Duel) *OcgDuel {
	return &OcgDuel{handle: d}
}

func (d *OcgDuel) Destroy() {
	close(d.incomingCh)
	lib.DestroyDuel(d.handle)
}

func (d *OcgDuel) Start() <-chan Message {
	d.messageCh = make(chan Message)
	d.incomingCh = make(chan []byte)

	go d.run()
	return d.messageCh
}

func (d *OcgDuel) run() {
	d.readMessages()
	lib.StartDuel(d.handle)

outer:
	for {
		status := lib.DuelProcess(d.handle)
		d.readMessages()
		switch status {
		case lib.ProcessorFlagEnd:
			break outer
		case lib.ProcessorFlagWaiting:
			d.messageCh <- MessageWaitingResponse{}
			r, ok := <-d.incomingCh
			if !ok {
				return
			}
			if r != nil {
				lib.DuelSetResponse(d.handle, r)
			}
		case lib.ProcessorFlagContinue:
			continue
		default:
			panic("invalid status")
		}
	}

	d.aliveLock.Lock()
	defer d.aliveLock.Unlock()

	close(d.messageCh)
	d.messageCh = nil
}

func (d *OcgDuel) readMessages() {
	d.aliveLock.Lock()
	defer d.aliveLock.Unlock()

	messages := duelGetMessage(d.handle)
	for _, message := range messages {
		if d.messageCh != nil {
			d.messageCh <- readMessage(message)
		}
	}
}

func (d *OcgDuel) SendResponse(r Response) {
	d.incomingCh <- r.responseWrite()
}

func (d *OcgDuel) SetupDeck(player int, mainDeck []uint32, extraDeck []uint32, shuffle bool) {
	if shuffle {
		rand.Shuffle(len(mainDeck), func(i, j int) {
			mainDeck[i], mainDeck[j] = mainDeck[j], mainDeck[i]
		})
	}

	var cardInfo lib.NewCardInfo
	cardInfo.Duelist = 0
	cardInfo.Team = uint8(player)
	cardInfo.Controller = uint8(player)
	cardInfo.Position = lib.PositionFaceDownDefense

	cardInfo.Location = lib.LocationDeck
	for _, card := range mainDeck {
		cardInfo.Code = card
		lib.DuelNewCard(d.handle, cardInfo)
	}

	cardInfo.Location = lib.LocationExtra
	for _, card := range extraDeck {
		cardInfo.Code = card
		lib.DuelNewCard(d.handle, cardInfo)
	}
}
