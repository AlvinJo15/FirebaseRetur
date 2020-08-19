package boot

import (
	"firebase/internal/data/auth"
	"firebase/pkg/firebaseclient"
	"firebase/pkg/httpclient"
	"log"
	"net/http"

	"firebase/internal/config"

	firebaseData "firebase/internal/data/firebase"
	firebaseServer "firebase/internal/delivery/http"
	firebaseHandler "firebase/internal/delivery/http/firebase"
	firebaseService "firebase/internal/service/firebase"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s         firebaseServer.Server    // HTTP Server Object
		fc        *firebaseclient.Client   // Firebase initiation
		firebaseD firebaseData.Data        // Domain data layer
		firebaseS firebaseService.Service  // Domain service layer
		firebaseH *firebaseHandler.Handler // Domain handler
		cfg       *config.Config           // Configuration object
		httpc     *httpclient.Client
		ad        auth.Data
		cred      map[string]string
	)

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg, cred = config.Get()

	// // Open MySQL DB Connection
	// db, err := sqlx.Open("mysql", cfg.Database.Master)
	// if err != nil {
	// 	log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	// }

	fc, err = firebaseclient.NewClient(cfg, cred)
	if err != nil {
		return err
	}

	httpc = httpclient.NewClient()
	ad = auth.New(httpc, cfg.API.Auth)

	// Diganti dengan domain yang anda buat
	firebaseD = firebaseData.New(fc)
	firebaseS = firebaseService.New(firebaseD, ad)
	firebaseH = firebaseHandler.New(firebaseS)

	s = firebaseServer.Server{
		Firebase: firebaseH,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
