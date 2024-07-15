package main

import (
    "net/http"

    "github.com/antonlindstrom/pgstore"
    "github.com/gin-contrib/sessions"
    gsessions "github.com/gorilla/sessions"
)

type PGSessionStore struct {
    *pgstore.PGStore
}

func NewPGSessionStore(pgStore *pgstore.PGStore) *PGSessionStore {
    return &PGSessionStore{PGStore: pgStore}
}

func (p *PGSessionStore) Options(options sessions.Options) {
    p.PGStore.Options = &gsessions.Options{
        Path:     options.Path,
        Domain:   options.Domain,
        MaxAge:   options.MaxAge,
        Secure:   options.Secure,
        HttpOnly: options.HttpOnly,
        SameSite: options.SameSite,
    }
}

func (p *PGSessionStore) Get(req *http.Request, name string) (*gsessions.Session, error) {
    store := p.PGStore
    return store.Get(req, name)
}

func (p *PGSessionStore) New(req *http.Request, name string) (*gsessions.Session, error) {
    store := p.PGStore
    return store.New(req, name)
}

func (p *PGSessionStore) Save(req *http.Request, w http.ResponseWriter, session *gsessions.Session) error {
    store := p.PGStore
    return store.Save(req, w, session)
}
