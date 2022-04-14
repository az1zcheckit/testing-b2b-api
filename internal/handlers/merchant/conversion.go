package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"encoding/json"
	"net/http"
)

// Conversion    godoc
// @Summary      Conversion route
// @Description  Conversion
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Param        conversion body  models.ConversionRequest  true  "conversion"  Format(json)
// @Success      200  {object}   models.ConversionResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/conversion [post]
func (mh merchHandler) Conversion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")

		if len(token) == 0 {
			response.ToJson(w, http.StatusBadRequest, response.TokenIsEmpty)
		}
		conv := &merchant.Conversion{}
		if err := json.NewDecoder(r.Body).Decode(conv); err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		documentId, err := mh.svc.MerchRepositoryInstance().Conversion(ctx, token, "EN", *conv)

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"documentId": documentId,
			"msg":        err,
		})
	}
}
