package core

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
)

type ImapClient struct {
	Server string //服务端地址
	Port   int    //端口
	User   string //用户
	Pwd    string //密码或授权码
	c      *client.Client
}

func (this *ImapClient) Connect() {
	if this.Server != "" && this.User != "" && this.Pwd != "" {
		var err error
		this.c, err = client.Dial(fmt.Sprintf("%s:%d", this.Server, this.Port))
		if err != nil {
			log.Println("connect error ", err.Error())
			return
		}

		if err := this.c.Login(this.User, this.Pwd); err != nil {

		}

	}
}

func (this *ImapClient) Disconnect() {
	if this.c != nil {
		err := this.c.Logout()
		if err != nil {
			log.Println("logout error ", err.Error())
		}

		this.c.Close()
	}
}

func (this *ImapClient) GetBoxList() []string {
	if this.c != nil {
		mailboxes := make(chan *imap.MailboxInfo, 10)
		done := make(chan error, 1)
		go func() {
			done <- this.c.List("", "*", mailboxes)
		}()

		log.Println("Mailboxes:")
		var list []string
		for m := range mailboxes {
			list = append(list, m.Name)
			log.Println("* " + m.Name)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		return list
	}
	return nil
}

func (this *ImapClient) GetMessage(box string) {
	if this.c != nil {
		mbox, err := this.c.Select(box, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Flags for INBOX:", mbox.Flags)

		// Get the last 4 messages
		from := uint32(1)
		to := mbox.Messages
		if mbox.Messages > 3 {
			// We're using unsigned integers here, only substract if the result is > 0
			from = mbox.Messages - 3
		}
		seqset := new(imap.SeqSet)
		seqset.AddRange(from, to)

		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)
		go func() {
			done <- this.c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
		}()

		log.Println("Last 4 messages:")
		for msg := range messages {
			log.Println("* " + msg.Envelope.Subject)

		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		log.Println("Done!")
	}
}
