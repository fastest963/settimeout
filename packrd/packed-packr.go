// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "989baf8dff96c1772b9ce45468da44bd"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"0d0c2b18cf90a0cffcd5d496655c0398": "1f8b08000000000000ffecc9a10d80401044d1b78204790a8da434fa6f60c901091a3f4f4de653ca186ee7c28603f3dacdbe3eb17cdedddd2222e2bf2b0000ffffe1e947677e050000",
		"717fdd01e40cead73ba7b6551033613c": "1f8b08000000000000ff8c94416be4381085ef81fc87470ecb2e847627210b9b759ccd2e0bc39c06ba07668eb25db694d82a8daa9c9efef783d4ee49d29384f4a1b190ead5d3f70a9556c7a13a3e2a2d99b63a3e0280529d0e5409e9da8dc493c2080c56141f5c4365b1db3e3e2a8bb9a8acb9ddee8bedf9eb95f67c7f2acc1fe9b70ad4b86e0bb50433f2e415dc41dd48f8bda5ce4c830a9431ba6170420dfb56fec096276c8cd75c26141f28a643321005d4d471244492c0be75be5f947544f1b4e9140247157c58af3f21d2b78944051ba7362be6f63539dfe765306a0f25bef284c67898411842be45c771348a915bd7398a82c1dd136eee84fd296e1a91f46f86a136cdfdb592e869da3cc58df3e9b377dd62e65324407ba217d56721fc6784a42cecc54b0cf75e2621a875999710c1756828aa711ec1a4eb7297c84598106022a11eb8b97f09d06d9bb8c1a011c934529cca3010dd0e04357dee289637390cd7edf7c412291a9e8676a7bf6b294e095de411039b1c0a5666cc7e7dff48dee04e0ecde8e338e5b86303e7214d7441939305d696903883833af6d95a4d8947bb93befd78fb05097e6693e0ef5c45f22d458ca496db9fdcd2a5d4d2365b7d299450fde66b097f1f2695c7e9ffef660cc3615a56355c158590ced759382e4ad1c8beaf2e97cb6559cc0b6c32bbdd2c771c91769fcdffaeed7b84cf969226ef75edb325665118df22924ed12397bcbb87bd7cab016c9e38dfe212cffcbfc532542b9e624357280d6ca4eefa64f6d03bb553bd68782c3a2329c9bffebc78e2eca47ae7c1b230d5b34bfebbfda5996e9c2ac52c726746126ba2bafea4fae7c9ea51a82ce6b7b02ce697f547000000ffffc08f49f662050000",
		"7c52f5bf7077f8f322c76466a0e819f2": "1f8b08000000000000ff0a2d4e2dd24d4c4fcd2bb152d0e2e572ccc9c92fb752d057e1e572c92c4e84f278b900010000ffff482d5d2b27000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("assets", "./assets")
		b.SetResolver("favicon.ico", packr.Pointer{ForwardBox: gk, ForwardPath: "0d0c2b18cf90a0cffcd5d496655c0398"})
		b.SetResolver("index.html", packr.Pointer{ForwardBox: gk, ForwardPath: "717fdd01e40cead73ba7b6551033613c"})
		b.SetResolver("robots.txt", packr.Pointer{ForwardBox: gk, ForwardPath: "7c52f5bf7077f8f322c76466a0e819f2"})
	}()

	return nil
}()
