package main

import (
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type x11Writer struct {
	conn       *xgb.Conn
	root       xproto.Window
	atomWMName xproto.Atom
	atomNetWM  xproto.Atom
	atomUTF8   xproto.Atom
}

var (
	writerOnce sync.Once
	writerX11  *x11Writer
	writerErr  error
)

func initX11Writer() {
	conn, err := xgb.NewConn()
	if err != nil {
		writerErr = err
		return
	}
	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	atom := func(name string) (xproto.Atom, error) {
		cookie := xproto.InternAtom(conn, true, uint16(len(name)), name)
		reply, err := cookie.Reply()
		if err != nil {
			return 0, err
		}
		return reply.Atom, nil
	}

	atomWMName, err := atom("WM_NAME")
	if err != nil {
		writerErr = err
		conn.Close()
		return
	}
	atomNetWM, err := atom("_NET_WM_NAME")
	if err != nil {
		writerErr = err
		conn.Close()
		return
	}
	atomUTF8, err := atom("UTF8_STRING")
	if err != nil {
		writerErr = err
		conn.Close()
		return
	}

	writerX11 = &x11Writer{
		conn:       conn,
		root:       root,
		atomWMName: atomWMName,
		atomNetWM:  atomNetWM,
		atomUTF8:   atomUTF8,
	}
}

func (w *x11Writer) writeStatus(s string) error {
	bs := []byte(s)
	// Old-style WM_NAME (STRING) + modern _NET_WM_NAME (UTF8_STRING).
	if err := xproto.ChangePropertyChecked(w.conn, xproto.PropModeReplace, w.root, w.atomWMName, xproto.AtomString, 8, uint32(len(bs)), bs).Check(); err != nil {
		return err
	}
	if err := xproto.ChangePropertyChecked(w.conn, xproto.PropModeReplace, w.root, w.atomNetWM, w.atomUTF8, 8, uint32(len(bs)), bs).Check(); err != nil {
		return err
	}
	w.conn.Sync()
	return nil
}

func write(s string) error {
	writerOnce.Do(initX11Writer)
	if writerErr == nil && writerX11 != nil {
		if err := writerX11.writeStatus(s); err == nil {
			return nil
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return exec.CommandContext(ctx, "xsetroot", "-name", s).Run()
}
