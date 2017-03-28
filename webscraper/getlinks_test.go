package webscraper

import (
	"fmt"
	"testing"
)

func TestGetLinksWithDivClass(t *testing.T) {
	// GIVEN
	pageStr := `
	<html>
	<head>
	<title>THIS IS THE TITLE</title>
	</head>
	<body>
	<div class="product ">
		<div class="productInner">
			<div class="productInfoWrapper">
				<div class="productInfo">
					<h3>
	                                    <a href="http://www.sainsburys.co.uk/shop/gb/groceries/ripe---ready/sainsburys-avocado--ripe---ready-x2" >
	                                        Sainsbury's Avocado, Ripe & Ready x2
		                                <img src="http://c2.sainsburys.co.uk/wcsstore7.20.1.145/ExtendedSitesCatalogAssetStore/images/catalog/productImages/22/0000001600322/0000001600322_M.jpeg" alt="" />
		                             </a>
					</h3>
				</div>
			</div>
		</div>
	</div>
	<div class="product ">
		<div class="productInner">
			<div class="productInfoWrapper">
				<div class="productInfo">
					<h3>
						<a href="http://www.sainsburys.co.uk/shop/gb/groceries/ripe---ready/sainsburys-avocados--ripe---ready-x4" >
						Sainsbury's Avocados, Ripe & Ready x4
						<img src="http://c2.sainsburys.co.uk/wcsstore7.20.1.145/ExtendedSitesCatalogAssetStore/images/catalog/productImages/15/0000000184915/0000000184915_M.jpeg" alt="" />
						</a>
					</h3>
				</div>
			</div>
		</div>
	</div>
	</body>
	</html>
	`
	//page := []byte(pageStr)
	// WHEN
	urls := GetLinksWithDivClass([]byte(pageStr), "productInfo")
	fmt.Println(len(urls))
	for _, url := range urls {
		fmt.Println("urls", url)
	}

	// THEN
	expects := []string{
		"http://www.sainsburys.co.uk/shop/gb/groceries/ripe---ready/sainsburys-avocado--ripe---ready-x2",
		"http://www.sainsburys.co.uk/shop/gb/groceries/ripe---ready/sainsburys-avocados--ripe---ready-x4",
	}

	for i, expect := range expects {
		if urls[i] != expect {
			t.Errorf("Test %d failed, Expected '%s', got '%s'", i, expect, urls[i])
		}
	}
}
