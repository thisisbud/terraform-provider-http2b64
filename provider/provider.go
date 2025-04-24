//package provider
//
//import (
//	"context"
//	"github.com/hashicorp/terraform-plugin-framework/datasource"
//	"github.com/hashicorp/terraform-plugin-framework/provider"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//)
//
//func New(version string) func() *schema.Provider {
//	return func() *schema.Provider {
//		p := &schema.Provider{
//
//			ResourcesMap: map[string]*schema.Resource{
//				"http2b64": resourcehttp2b64(),
//			},
//		}
//
//		p.ConfigureContextFunc = configure(version, p)
//
//		return p
//	}
//}
//
//type apiClient struct {
//}
//
//func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
//	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
//		return &apiClient{}, nil
//	}
//}

// Break

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	_ "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	_ "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*http2b64Provider)(nil)

//var _ provider.ProviderWithMetadata = (*http2b64Provider)(nil)

type http2b64Provider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &http2b64Provider{}
	}
}

func (p *http2b64Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *http2b64Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "thisisbud"
}

func (p *http2b64Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		//NewDataSource,
	}
}

func (p *http2b64Provider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}

func (p *http2b64Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}
