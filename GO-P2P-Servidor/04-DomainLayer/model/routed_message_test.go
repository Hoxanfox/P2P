package model

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestNewRoutedMessage_Success(t *testing.T) {
    msgID := uuid.New()
    peerID := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    rm, err := NewRoutedMessage(msgID, peerID, now)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if rm.MensajeID() != msgID {
        t.Errorf("MensajeID: esperado %v, obtuvo %v", msgID, rm.MensajeID())
    }
    if rm.NodoDestinoID() != peerID {
        t.Errorf("NodoDestinoID: esperado %v, obtuvo %v", peerID, rm.NodoDestinoID())
    }
    if !rm.EnrutaAt().Equal(now) {
        t.Errorf("EnrutaAt: esperado %v, obtuvo %v", now, rm.EnrutaAt())
    }
}

func TestNewRoutedMessage_Errors(t *testing.T) {
    validMsgID := uuid.New()
    validPeerID := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name    string
        fn      func() (*RoutedMessage, error)
        wantErr error
    }{
        {"MensajeID inválido", func() (*RoutedMessage, error) {
            return NewRoutedMessage(uuid.Nil, validPeerID, now)
        }, ErrRoutedMessageMensajeIDNil},

        {"NodoDestinoID inválido", func() (*RoutedMessage, error) {
            return NewRoutedMessage(validMsgID, uuid.Nil, now)
        }, ErrRoutedMessageDestinoIDNil},

        {"EnrutaAt cero", func() (*RoutedMessage, error) {
            return NewRoutedMessage(validMsgID, validPeerID, time.Time{})
        }, ErrRoutedMessageEnrutaAtZero},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := c.fn()
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
