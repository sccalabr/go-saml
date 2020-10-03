package examples

import (
	"fmt"
	"github.com/LoginRadius/go-saml"
	"github.com/LoginRadius/go-saml/examples/utitlity"
)

func MetadataExample() {

	idp := saml.IdentityProvider{
		IsIdpInitiated:       false,
		Issuer:               "https://identity-provider.com/",
		IDPCert:              utitlity.GetIdpCertificate(),
		NameIdentifierFormat: saml.AttributeFormatUnspecified,
	}

	idp.AddSingleSignOnService(saml.MetadataBinding{
		Binding:  saml.HTTPPostBinding,
		Location: "https://identity-provider.com/saml/post",
	})

	idp.AddSingleSignOutService(saml.MetadataBinding{
		Binding:  saml.HTTPPostBinding,
		Location: "https://identity-provider.com/saml/post/logout",
	})

	idp.AddContactPerson(saml.MetadataBinding{
		Binding:  saml.HTTPPostBinding,
		Location: "https://identity-provider.com/saml/post",
	})

	idp.AddOrganization(saml.MetadataBinding{
		Binding:  saml.HTTPPostBinding,
		Location: "https://identity-provider.com/saml/post",
	})

	// Generate xml for IDP Metadata
	metadata, err := idp.MetaDataResponse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------Metadata------------------\n", metadata)
}
