/*
 * @Author: zzzzztw
 * @Date: 2023-07-02 15:11:36
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-02 16:18:07
 * @FilePath: /Gee/gee/router.go
 */
package gee

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newrouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusBadRequest, "404 not found: %q", c.Path)
	}
}
