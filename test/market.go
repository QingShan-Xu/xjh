package test

import "github.com/QingShan-Xu/xjh/bm"

var D database

type database struct {
	Pet Pet
}

type Pet struct {
	bm.Model

	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

var API api

type api struct {
	ReqPet       ReqBindGetUser
	ReqUpdatePet ReqUpdatePet
	ReqCreatePet ReqCreatePet
}

type ReqBindGetUser struct {
	MarketID   int    `uri:"market_id"`
	MarketName string `param:"market_name" binding:"required"`
}

type ReqUpdatePet struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ReqCreatePet struct {
	UserID int    `uri:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}
