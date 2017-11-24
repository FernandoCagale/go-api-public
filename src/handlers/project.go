package handlers

import (
	"net/http"

	"github.com/FernandoCagale/go-api-public/src/models"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database   = "go-api-public"
	collection = "projects"
)

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) SaveProject(c echo.Context) error {
	project := new(models.Project)

	collection, valid := getCollection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	if errors, valid := project.Validate(); !valid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	project.ID = bson.NewObjectId()

	if err := collection.Insert(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}
	return c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) UpdateProject(c echo.Context) error {
	id := bson.ObjectIdHex(c.Param("id"))
	project := new(models.Project)

	collection, valid := getCollection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	if errors, valid := project.Validate(); !valid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := collection.UpdateId(id, &project); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "NotFound",
		})
	}

	return c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) GetAllProject(c echo.Context) error {
	collection, valid := getCollection(c)
	if !valid {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	var projects []models.Project
	if err := collection.Find(nil).Sort("-start").All(&projects); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetProject(c echo.Context) error {
	collection, valid := getCollection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	id := bson.ObjectIdHex(c.Param("id"))

	project := models.Project{}
	if err := collection.FindId(id).One(&project); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "NotFound",
		})
	}

	return c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	collection, valid := getCollection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	id := bson.ObjectIdHex(c.Param("id"))

	if err := collection.Remove(bson.M{"_id": id}); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "NotFound",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted",
	})
}

func getCollection(c echo.Context) (*mgo.Collection, bool) {
	if mongo := c.Get("mongo"); mongo != nil {
		session := mongo.(*mgo.Session)
		return session.DB(Database).C(Collection), true
	}
	return nil, false
}
