package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/mitchellh/mapstructure"

	"github.com/udistrital/terceros_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

// InfoComplementariaTerceroController operations for InfoComplementariaTercero
type InfoComplementariaTerceroController struct {
	beego.Controller
}

// URLMapping ...
func (c *InfoComplementariaTerceroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("PostPadre", c.PostPadre)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create InfoComplementariaTercero
// @Param	body		body 	models.InfoComplementariaTercero	true		"body for InfoComplementariaTercero content"
// @Success 201 {int} models.InfoComplementariaTercero
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *InfoComplementariaTerceroController) Post() {
	var v models.InfoComplementariaTercero
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = time_bogota.TiempoBogotaFormato()
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if _, err := models.AddInfoComplementariaTercero(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// PostPadre ...
// @Title PostPadre
// @Description create InfoComplementariaTercero Padre
// @Param	body		body 	models.InfoComplementariaTercero	true		"body for InfoComplementariaTercero content"
// @Success 201 {int} models.InfoComplementariaTercero
// @Failure 400 the request contains incorrect syntax
// @router /padre [post]
func (c *InfoComplementariaTerceroController) PostPadre() {
	var v models.InfoComplementariaTercero
	var resultado map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &resultado); err == nil {

		datoJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["Dato"])
		//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
		//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
		infoComplementariaPadre := map[string]interface{}{
			"TerceroId":            resultado["TerceroId"].(map[string]interface{}),
			"InfoComplementariaId": resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"],
			"Activo":               true,
			"Dato":                 string(datoJson),
			"FechaCreacion":        time_bogota.TiempoBogotaFormato(),
			"FechaModificacion":    time_bogota.TiempoBogotaFormato(),
		}
		err := mapstructure.Decode(infoComplementariaPadre, &v)
		if err != nil {
			panic(err)
		}

		if _, err := models.AddInfoComplementariaTercero(&v); err == nil {
			var va models.InfoComplementariaTercero
			datoJson2, _ := json.Marshal(resultado["FechaInicio"].(map[string]interface{})["Dato"])
			//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
			//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
			infoComplementariaPadre2 := map[string]interface{}{
				"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
				"InfoComplementariaId":     resultado["FechaInicio"].(map[string]interface{})["InfoComplementaria"],
				"Activo":                   true,
				"Dato":                     string(datoJson2),
				"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
				"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
				"InfoCompleTerceroPadreId": v,
			}
			err := mapstructure.Decode(infoComplementariaPadre2, &va)
			if err != nil {
				panic(err)
			}
			if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

				var va models.InfoComplementariaTercero
				datoJson3, _ := json.Marshal(resultado["FechaFin"].(map[string]interface{})["Dato"])
				//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
				//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
				infoComplementariaPadre3 := map[string]interface{}{
					"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
					"InfoComplementariaId":     resultado["FechaFin"].(map[string]interface{})["InfoComplementaria"],
					"Activo":                   true,
					"Dato":                     string(datoJson3),
					"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
					"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
					"InfoCompleTerceroPadreId": v,
				}
				err := mapstructure.Decode(infoComplementariaPadre3, &va)
				if err != nil {
					panic(err)
				}

				if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

					var va models.InfoComplementariaTercero
					datoJson4, _ := json.Marshal(resultado["TituloTrabajoGrado"].(map[string]interface{})["Dato"])
					//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
					//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
					infoComplementariaPadre3 := map[string]interface{}{
						"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
						"InfoComplementariaId":     resultado["TituloTrabajoGrado"].(map[string]interface{})["InfoComplementaria"],
						"Activo":                   true,
						"Dato":                     string(datoJson4),
						"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
						"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
						"InfoCompleTerceroPadreId": v,
					}
					err := mapstructure.Decode(infoComplementariaPadre3, &va)
					if err != nil {
						panic(err)
					}
					if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

						var va models.InfoComplementariaTercero
						datoJson4, _ := json.Marshal(resultado["DesTrabajoGrado"].(map[string]interface{})["Dato"])
						//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
						//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
						infoComplementariaPadre3 := map[string]interface{}{
							"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
							"InfoComplementariaId":     resultado["DesTrabajoGrado"].(map[string]interface{})["InfoComplementaria"],
							"Activo":                   true,
							"Dato":                     string(datoJson4),
							"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
							"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
							"InfoCompleTerceroPadreId": v,
						}
						err := mapstructure.Decode(infoComplementariaPadre3, &va)
						if err != nil {
							panic(err)
						}
						if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

							var va models.InfoComplementariaTercero
							datoJson4, _ := json.Marshal(resultado["DocumentoId"].(map[string]interface{})["Dato"])
							//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
							//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
							infoComplementariaPadre3 := map[string]interface{}{
								"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
								"InfoComplementariaId":     resultado["DocumentoId"].(map[string]interface{})["InfoComplementaria"],
								"Activo":                   true,
								"Dato":                     string(datoJson4),
								"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
								"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
								"InfoCompleTerceroPadreId": v,
							}
							err := mapstructure.Decode(infoComplementariaPadre3, &va)
							if err != nil {
								panic(err)
							}
							if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

								var va models.InfoComplementariaTercero
								datoJson4, _ := json.Marshal(resultado["NitUniversidad"].(map[string]interface{})["Dato"])
								//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
								//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
								infoComplementariaPadre3 := map[string]interface{}{
									"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
									"InfoComplementariaId":     resultado["NitUniversidad"].(map[string]interface{})["InfoComplementaria"],
									"Activo":                   true,
									"Dato":                     string(datoJson4),
									"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
									"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
									"InfoCompleTerceroPadreId": v,
								}
								err := mapstructure.Decode(infoComplementariaPadre3, &va)
								if err != nil {
									panic(err)
								}
								if _, err := models.AddInfoComplementariaTercero(&va); err == nil {

									var va models.InfoComplementariaTercero
									datoJson4, _ := json.Marshal(resultado["NivelFormacion"].(map[string]interface{})["Dato"])
									//terceroJson,_ := json.Marshal( resultado["TerceroId"].(map[string]interface{}))
									//infoComplemenJson, _ := json.Marshal(resultado["ProgramaAcademico"].(map[string]interface{})["InfoComplementaria"])
									infoComplementariaPadre3 := map[string]interface{}{
										"TerceroId":                resultado["TerceroId"].(map[string]interface{}),
										"InfoComplementariaId":     resultado["NivelFormacion"].(map[string]interface{})["InfoComplementaria"],
										"Activo":                   true,
										"Dato":                     string(datoJson4),
										"FechaCreacion":            time_bogota.TiempoBogotaFormato(),
										"FechaModificacion":        time_bogota.TiempoBogotaFormato(),
										"InfoCompleTerceroPadreId": v,
									}
									err := mapstructure.Decode(infoComplementariaPadre3, &va)
									if err != nil {
										panic(err)
									}
									if _, err := models.AddInfoComplementariaTercero(&va); err == nil {
										c.Ctx.Output.SetStatus(201)
										c.Data["json"] = v

									} else {

										logs.Error(err)
										//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = err
										c.Abort("400")
										fmt.Println(c.Data["system"])
										fmt.Println(err)
									}

								} else {

									logs.Error(err)
									//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = err
									c.Abort("400")
									fmt.Println(c.Data["system"])
									fmt.Println(err)
								}

							} else {

								logs.Error(err)
								//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = err
								c.Abort("400")
								fmt.Println(c.Data["system"])
								fmt.Println(err)
							}

						} else {

							logs.Error(err)
							//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = err
							c.Abort("400")
							fmt.Println(c.Data["system"])
							fmt.Println(err)
						}

					} else {

						logs.Error(err)
						//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = err
						c.Abort("400")
						fmt.Println(c.Data["system"])
						fmt.Println(err)
					}

				} else {

					logs.Error(err)
					//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = err
					c.Abort("400")
					fmt.Println(c.Data["system"])
					fmt.Println(err)
				}

			} else {

				logs.Error(err)
				//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = err
				c.Abort("400")
				fmt.Println(c.Data["system"])
				fmt.Println(err)
			}

		} else {

			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
			fmt.Println(c.Data["system"])
			fmt.Println(err)
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get InfoComplementariaTercero by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.InfoComplementariaTercero
// @Failure 404 not found resource
// @router /:id [get]
func (c *InfoComplementariaTerceroController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetInfoComplementariaTerceroById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get InfoComplementariaTercero
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.InfoComplementariaTercero
// @Failure 404 not found resource
// @router / [get]
func (c *InfoComplementariaTerceroController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllInfoComplementariaTercero(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the InfoComplementariaTercero
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.InfoComplementariaTercero	true		"body for InfoComplementariaTercero content"
// @Success 200 {object} models.InfoComplementariaTercero
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *InfoComplementariaTerceroController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.InfoComplementariaTercero{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		infoAd, _ := models.GetInfoComplementariaTerceroById(id)
		if infoAd != nil {
			v.FechaCreacion = time_bogota.TiempoCorreccionFormato(infoAd.FechaCreacion)
			v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		}
		if err := models.UpdateInfoComplementariaTerceroById(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the InfoComplementariaTercero
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *InfoComplementariaTerceroController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteInfoComplementariaTercero(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
