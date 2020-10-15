package ocgcore

import (
	"math/rand"
	"sync"
)

type OcgDuel struct {
	c *OcgCore
	h duelHandle

	messageCh  chan Message
	incomingCh chan []byte

	aliveLock sync.Mutex
}

func (d *OcgDuel) Destroy() {
	close(d.incomingCh)
	d.c.destroyDuel(d.h)
}

func (d *OcgDuel) Start() <-chan Message {
	d.messageCh = make(chan Message)
	d.incomingCh = make(chan []byte)

	go d.run()
	return d.messageCh
}

func (d *OcgDuel) run() {
	d.readMessages()
	d.c.startDuel(d.h)

outer:
	for {
		status := d.c.duelProcess(d.h)
		d.readMessages()
		switch status {
		case processorFlagEnd:
			break outer
		case processorFlagWaiting:
			d.messageCh <- MessageWaitingResponse{}
			r, ok := <-d.incomingCh
			if !ok {
				return
			}
			if r != nil {
				d.c.duelSetResponse(d.h, r)
			}
		case processorFlagContinue:
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

	messages := d.c.duelGetMessage(d.h)
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

	var cardInfo typeNewCardInfo
	cardInfo.duelist = 0
	cardInfo.team = uint8(player)
	cardInfo.con = uint8(player)

	cardInfo.pos = uint32(corePositionFaceDownDefense)

	cardInfo.loc = uint32(coreLocationDeck)
	for _, card := range mainDeck {
		cardInfo.code = card
		d.c.duelNewCard(d.h, cardInfo)
	}
	cardInfo.loc = uint32(coreLocationExtra)
	for _, card := range extraDeck {
		cardInfo.code = card
		d.c.duelNewCard(d.h, cardInfo)
	}
}
