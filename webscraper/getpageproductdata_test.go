package webscraper

import (
	"testing"
)

func TestGetPageProductData(t *testing.T) {
	// GIVEN
	pageStr := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Title</title>
</head>
<body>
<div class="productSummary">
    <div class="productTitleDescriptionContainer">
        <h1>Sainsbury's Avocado, Ripe & Ready x2</h1>
    </div>
    <div class="addToTrolleytabBox" >
        <div class="addToTrolleytabContainer addItemBorderTop">
            <div class="pricingAndTrolleyOptions">
                <div class="priceTab activeContainer priceTabContainer" id="addItem_138041"> <!-- CachedProductOnlyDisplay.jsp -->
                    <div class="pricing">
                        <p class="pricePerUnit">
                            £1.90<abbr title="per">/</abbr><abbr title="unit"><span class="pricePerUnitUnit">unit</span></abbr>
                        </p>
                        <p class="pricePerMeasure">£0.95<abbr
                                title="per">/</abbr><abbr
                                title="each"><span class="pricePerMeasureMeasure">ea</span></abbr>
                        </p>
                    </div>
                </div>
            </div><!-- End pricingAndTrolleyOptions -->
        </div><!-- End addToTrolleytabContainer -->
    </div>
</div>
<div class="mainProductInfoWrapper">
<h3 class="productDataItemHeader">Description</h3>
<div class="productText">
<p>Avocados</p>
<p>
<p></p>
</p>
</div>
<h3 class="productDataItemHeader">Manufacturer</h3>
<div class="productText">
<p>We are happy to replace this item if it is not satisfactory</p>
</div>
<div id="additionalItems_138041" class="additionalProductInfo">
    <div class="crossSell">
        <div class="crossSellContent">
            <div class="crossSellInfo">
                <h3 class="crossSellName">
                    <span class="access">Try this product with </span>
                    <a href="http://www.sainsburys.co.uk/shop/gb/groceries/sainsburys-11l-klip-lock-storage-set-3pk">
                        Sainsbury's Klip Lock Storage Set 1.1L x3
                        <img src="/wcsstore7.20.1.145/ExtendedSitesCatalogAssetStore/images/catalog/productImages/61/5053217772661/5053217772661_S.jpeg" alt="" />
                    </a>
                </h3>
            </div>
            <div class="pricingAndTrolleyFormWrapper">
                <div class="pricingReviews">
                    <div class="pricing">

                        <p class="pricePerUnit">
                            £9.00<abbr title="per">/</abbr><abbr title="unit"><span class="pricePerUnitUnit">unit</span></abbr>
                        </p>

                        <p class="pricePerMeasure">£9.00<abbr
                                title="per">/</abbr><abbr
                                title="each"><span class="pricePerMeasureMeasure">ea</span></abbr>
                        </p>

                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</div>
</body>
</html>
	`

	// WHEN
	productData := GetPageProductData([]byte(pageStr))

	// THEN
	expectedData := ProductData{
		Title:       "Sainsbury's Avocado, Ripe & Ready x2",
		Size:        "2kb",
		UnitPrice:   1.9,
		Description: "Avocados",
	}
	if productData.Title != expectedData.Title {
		t.Errorf("Test failed, Title: expected '%s', got '%s'", expectedData.Title, productData.Title)
	}
	if productData.Size != expectedData.Size {
		t.Errorf("Test failed, Size: expected '%s', got '%s'", expectedData.Size, productData.Size)
	}
	if productData.UnitPrice != expectedData.UnitPrice {
		t.Errorf("Test failed, UnitPrice: expected '%.2f', got '%.2f'", expectedData.UnitPrice, productData.UnitPrice)
	}
	if productData.Description != expectedData.Description {
		t.Errorf("Test failed, Description: expected '%s', got '%s'", expectedData.Description, productData.Description)
	}

}
