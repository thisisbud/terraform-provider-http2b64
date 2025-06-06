package httpprovider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/thisisbud/terraform-provider-http2b64/client"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*http2b64Resource)(nil)

type http2b64Resource struct {
	provider http2b64Provider //nolint:unused
}

func NewResource() resource.Resource {
	return &http2b64Resource{}
}

func (e *http2b64Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "http2b64"
}

func (e *http2b64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configurable_attribute": schema.StringAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Required: true,
			},
			"status_code": schema.StringAttribute{
				Computed: true,
			},
			"response_body_base64": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

type http2b64ResourceData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
	Url                   types.String `tfsdk:"url"`
	StatusCode            types.String `tfsdk:"status_code"`
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
	inputUrl := data.Url.ValueString()

	responseBody, statusCode, err := client.GetFile(inputUrl)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %v", err),
		)
		return
	}
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %v", err),
		)
		return
	}
	emptyParams := url.Values{}
	parsedUrl.RawQuery = emptyParams.Encode()
	data.Id = types.StringValue(parsedUrl.String())
	data.Url = types.StringValue(inputUrl)
	data.StatusCode = types.StringValue(strconv.Itoa(statusCode))
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

	//// Create resource using 3rd party API.
	//inputUrl := data.Url.ValueString()
	//
	//responseBody, statusCode, err := client.GetFile(inputUrl)
	//if err != nil {
	//	//return fmt.Errorf("error Getting resource '%v'", err)
	//	resp.Diagnostics.AddError(
	//		"API Error Creating Resource",
	//		fmt.Sprintf("... details ... %v", err),
	//	)
	//	return
	//}
	//parsedUrl, err := url.Parse(inputUrl)
	//if err != nil {
	//	resp.Diagnostics.AddError(
	//		"API Error Creating Resource",
	//		fmt.Sprintf("... details ... %v", err),
	//	)
	//	return
	//}
	//emptyParams := url.Values{}
	//parsedUrl.RawQuery = emptyParams.Encode()
	//data.Id = types.StringValue(parsedUrl.String())
	//data.Url = types.StringValue(inputUrl)
	//data.StatusCode = types.StringValue(strconv.Itoa(statusCode))
	//data.ResponseBodyBas64 = types.StringValue(responseBody)

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
	inputUrl := data.Url.ValueString()

	responseBody, statusCode, err := client.GetFile(inputUrl)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %v", err),
		)
		return
	}
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		resp.Diagnostics.AddError(
			"API Error Creating Resource",
			fmt.Sprintf("... details ... %v", err),
		)
		return
	}
	emptyParams := url.Values{}
	parsedUrl.RawQuery = emptyParams.Encode()
	data.Id = types.StringValue(parsedUrl.String())
	data.Url = types.StringValue(inputUrl)
	data.StatusCode = types.StringValue(strconv.Itoa(statusCode))
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
