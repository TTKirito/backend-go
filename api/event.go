package api

import (
	"database/sql"
	"net/http"

	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEventRequest struct {
	Title        string           `json:"title"`
	StartTime    int64            `json:"start_time" binding:"required"`
	EndTime      int64            `json:"end_time" binding:"required"`
	IsEmegency   bool             `json:"is_emegency"`
	Owner        int64            `json:"owner" binding:"required,min=1"`
	Note         string           `json:"note"`
	Type         string           `json:"type" binding:"required,eventType"`
	VisitType    string           `json:"visit_type" binding:"required,visitType"`
	Meeting      string           `json:"meeting" binding:"required"`
	Location     db.Location      `json:"location" binding:"required"`
	Participants []db.Participant `json:"participants" binding:"required"`
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEventTxParams{
		Title:        req.Title,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		IsEmegency:   req.IsEmegency,
		Owner:        req.Owner,
		Note:         req.Note,
		Type:         db.EventTypes(req.Type),
		VisitType:    db.VisitTypes(req.VisitType),
		Meeting:      req.Meeting,
		Location:     req.Location,
		Participants: req.Participants,
	}

	event, err := server.store.CreateEventTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)

}

type getEventRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event, err := server.store.GetEvent(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type listEventRequest struct {
	PageID    int32 `form:"page_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=10"`
	StartTime int64 `form:"start_time" binding:"required"`
	EndTime   int64 `form:"end_time" binding:"required"`
}

func (server *Server) listEvent(ctx *gin.Context) {
	var req listEventRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEventParams{
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	events, err := server.store.ListEvent(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, events)

}
