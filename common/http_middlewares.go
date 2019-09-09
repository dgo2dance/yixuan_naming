/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file http_middlewares.go
 * @package common
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/25/2019
 */

package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/valyala/fasthttp"
)

// JwtClaims : JWT claims
type JwtClaims struct {
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

// JSONEnvelope : Generate JSON data envelope
type JSONEnvelope struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Message   string      `json:"message"`
	Links     []string    `json:"links"`
	Data      interface{} `json:"data"`
}

func parseBasicAuthorization(auth string) (username, password string) {
	c, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		return
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}

	return cs[:s], cs[s+1:]
}

func parseBearerAuthorization(auth string) (string, string) {
	var tokenString = auth[7:]
	s := strings.Split(tokenString, ".")
	if len(s) == 3 {
		return tokenString, "jwt"
	}

	return tokenString, "normal"
}

func parseJWT(tokenString, key string) *JwtClaims {
	var (
		claims = &JwtClaims{}
	)

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil || !tkn.Valid {
		return nil
	}

	return claims
}

// HTTPGlobalRuntime : Set runtime into context
func HTTPGlobalRuntime(next fasthttp.RequestHandler, r *GlobalRuntime) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.SetUserValue("_g", r)
		next(ctx)
		return
	})
}

// HTTPAuthorization : Process HTTP authorization header
func HTTPAuthorization(next fasthttp.RequestHandler, requireAuthorization string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		var (
			user        string
			pass        string
			auths       string
			bearerToken string
			bearerType  string
			authType    = "unknown"
			ok          = false
			r           *GlobalRuntime
			c           *JwtClaims
		)

		auth := ctx.Request.Header.Peek("Authorization")
		if auth == nil {
			if requireAuthorization == "none" {
				ok = true
			}
		}

		auths = string(auth)
		if strings.HasPrefix(auths, "Basic ") {
			// Basic authoriztion
			user, pass = parseBasicAuthorization(auths)
			authType = "basic"
		} else if strings.HasPrefix(string(auth), "Bearer ") {
			// Bearer (Normal Token / JWT)
			bearerToken, bearerType = parseBearerAuthorization(auths)
			authType = "bearer"
		}

		tr := ctx.UserValue("_g")
		if tr != nil {
			r = tr.(*GlobalRuntime)
		}

		if r == nil {
			// No runtime
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetUserValue("_envelope_code", 10500)
			ctx.SetUserValue("_envelope_message", "Runtime error")
		} else {
			tDebug := r.Config.GetBool("authorization_jwt_debug")
			if requireAuthorization != "none" {
				if authType == "basic" && requireAuthorization == "basic" {
					tUser := r.Config.GetString("agent_app_id")
					tPass := r.Config.GetString("agent_app_sec")
					if tUser != user || tPass != pass {
						ctx.SetStatusCode(fasthttp.StatusUnauthorized)
						ctx.SetUserValue("_envelope_code", 10401)
						ctx.SetUserValue("_envelope_message", "Invalid authorization")
					} else {
						ctx.SetUserValue("auth_type", "basic")
						ok = true
					}
				} else if tDebug == false && authType == "bearer" {
					if bearerType == "jwt" && requireAuthorization == "jwt" {
						c = parseJWT(bearerToken, r.Config.GetString("jwt_key"))
						if c == nil {
							ctx.SetStatusCode(fasthttp.StatusUnauthorized)
							ctx.SetUserValue("_envelope_code", 10401)
							ctx.SetUserValue("_envelope_message", "JWT token invalid")
						} else {
							ctx.SetUserValue("auth_type", "jwt")
							ctx.SetUserValue("jwt_claims", c)
							ok = true
						}
					} else if requireAuthorization == "token" {
						// Find token
						ctx.SetUserValue("auth_type", "bearer_token")
						ctx.SetUserValue("bearer_token", bearerToken)
						ok = true
					} else {
						// Unknown
						ctx.SetStatusCode(fasthttp.StatusBadRequest)
						ctx.SetUserValue("_envelope_code", 10400)
						ctx.SetUserValue("_envelope_message", "Invalid authorization type")
					}
				}
			} else {
				ok = true
			}
		}

		if ok {
			next(ctx)
		}

		return
	})
}

// HTTPEnvelope : Add JSON envelope to data
func HTTPEnvelope(prev fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		prev(ctx)

		ctx.ResetBody()

		envelope := &JSONEnvelope{
			Message:   "OK",
			Timestamp: time.Now().UnixNano(),
		}
		tCode := ctx.UserValue("_envelope_code")
		tLinks := ctx.UserValue("_envelope_links")
		tMessage := ctx.UserValue("_envelope_message")
		tData := ctx.UserValue("_envelope_data")

		if tCode != nil {
			envelope.Code = tCode.(int)
		}
		if tLinks != nil {
			envelope.Links = tCode.([]string)
		}
		if tMessage != nil {
			envelope.Message = tMessage.(string)
		}
		if tData != nil {
			envelope.Data = tData
		}

		// MIME
		ctx.Response.Header.Set("Content-Type", "application/json")
		// Add CORS
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Response.Header.Set("Access-Control-Max-Age", "86400")

		ret, _ := json.Marshal(envelope)
		fmt.Fprint(ctx, string(ret))

		return
	})
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
