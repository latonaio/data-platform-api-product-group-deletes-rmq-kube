package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-product-group-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-group-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) ProductGroup(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductGroup {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE productGroup.ProductGroup = \"%s\", input.ProductGroup.ProductGroup"),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	productGroup.ProductGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_type_product_type_data as productGroup 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductGroup(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
