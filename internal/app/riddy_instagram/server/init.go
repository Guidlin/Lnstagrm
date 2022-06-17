package server

import "riddy_instagram/internal/app/riddy_instagram/router"

func Init() {
	r := router.Router()
	r.Run(":3000")
}
