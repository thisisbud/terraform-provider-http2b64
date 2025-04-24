//package provider
//
//import (
//	"fmt"
//	"github.com/hashicorp/terraform-plugin-framework/resource"
//	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
//	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
//	"github.com/hashicorp/terraform-plugin-framework/types"
//	"github.com/hashicorp/terraform-plugin-log/tflog"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//
//	"github.com/ackers-bud/terraform-provider-http2b64/client"
//)
//
//func resourcehttp2b64() *schema.Resource {
//	return &schema.Resource{
//		Create: Create,
//		Update: Update,
//		Read:   ReadUrl,
//		Delete: Delete,
//
//		Schema: map[string]*schema.Schema{
//			"url": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//
//			"id": {
//				Description: "The ID of this resource.",
//				Type:        schema.TypeString,
//				Computed:    true,
//			},
//
//			"status_code": {
//				Description: "the returned http status code",
//				Type:        schema.TypeInt,
//				Computed:    true,
//			},
//
//			"response_body_base64": {
//				Description: "The Returned body base64 encoded",
//				Type:        schema.TypeString,
//				Computed:    true,
//			},
//		},
//	}
//}
//
////nolint:ineffassign
//func Create(d *schema.ResourceData, meta interface{}) error {
//
//	url := d.Get("url").(string)
//
//	responseBody, statusCode, err := client.GetFile(url)
//	if err != nil {
//		return fmt.Errorf("error Getting resource '%v'", err)
//	}
//
//	err = d.Set("status_code", statusCode)
//	if err != nil {
//		return fmt.Errorf("error Setting status_code '%v'", err)
//	}
//
//	err = d.Set("response_body_base64", responseBody)
//	if err != nil {
//		return fmt.Errorf("error Setting response_body_base64 '%v'", err)
//	}
//
//	d.SetId(url)
//	return nil
//}
//
//func Update(d *schema.ResourceData, meta interface{}) error {
//
//	url := d.Get("url").(string)
//
//	responseBody, statusCode, err := client.GetFile(url)
//	if err != nil {
//		return fmt.Errorf("error Getting resource '%v'", err)
//	}
//
//	err = d.Set("status_code", statusCode)
//	if err != nil {
//		return fmt.Errorf("error Setting status_code '%v'", err)
//	}
//
//	err = d.Set("response_body_base64", responseBody)
//	if err != nil {
//		return fmt.Errorf("error Setting response_body_base64 '%v'", err)
//	}
//
//	d.SetId(url)
//
//	return nil
//}
//
//func ReadUrl(d *schema.ResourceData, meta interface{}) error {
//
//	return nil
//}
//
//func Delete(d *schema.ResourceData, meta interface{}) error {
//
//	d.SetId("")
//
//	return nil
//}

// BREAK

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/thisisbud/terraform-provider-http2b64/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*http2b64Resource)(nil)

type http2b64Resource struct {
	provider http2b64Provider
}

func NewResource() resource.Resource {
	return &http2b64Resource{}
}

func (e *http2b64Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (e *http2b64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url":                  schema.StringAttribute{},
			"status_code":          schema.StringAttribute{},
			"response_body_base64": schema.StringAttribute{},
		},
	}
}

type http2b64ResourceData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
	Url                   types.String `tfsdk:"url"`
	StatusCode            types.Int32  `tfsdk:"status_code"`
	ResponseBodyBas64     types.String `tfsdk:"response_body_base64"`
}

func (e *http2b64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data http2b64ResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create resource using 3rd party API.
	url := data.Url.String()

	responseBody, statusCode, err := client.GetFile(url)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(url)
	data.Url = types.StringValue(url)
	data.StatusCode = types.Int32Value(int32(statusCode))
	data.ResponseBodyBas64 = types.StringValue(responseBody)

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *http2b64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data http2b64ResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create resource using 3rd party API.
	url := data.Url.String()

	responseBody, statusCode, err := client.GetFile(url)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(url)
	data.Url = types.StringValue(url)
	data.StatusCode = types.Int32Value(int32(statusCode))
	data.ResponseBodyBas64 = types.StringValue(responseBody)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *http2b64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data http2b64ResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create resource using 3rd party API.
	url := data.Url.String()

	responseBody, statusCode, err := client.GetFile(url)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(url)
	data.Url = types.StringValue(url)
	data.StatusCode = types.Int32Value(int32(statusCode))
	data.ResponseBodyBas64 = types.StringValue(responseBody)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *http2b64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data http2b64ResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete resource using 3rd party API.
}
