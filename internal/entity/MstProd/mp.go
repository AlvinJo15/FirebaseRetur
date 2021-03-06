package product

import (
	"gopkg.in/guregu/null.v3/zero"
)

// MstProduct ...
type MstProduct struct {
	ProRunningID       zero.Int    `db:"Pro_RunningID" json:"pro_runningid"`
	ProCode            zero.String `db:"Pro_Code" json:"pro_code"`
	ProName            zero.String `db:"Pro_Name" json:"pro_name"`
	ProName2           zero.String `db:"Pro_Name2" json:"pro_name2"`
	ProDeptCode        zero.String `db:"Pro_DeptCode" json:"pro_deptcode"`
	ProPrinCode        zero.String `db:"Pro_PrinCode" json:"pro_princode"`
	ProCtrlCode        zero.String `db:"Pro_CtrlCode" json:"pro_ctrlcode"`
	ProStsMargin       zero.Int    `db:"Pro_StsMargin" json:"pro_stsmargin"`
	ProExpDateYN       zero.String `db:"Pro_ExpDateYN" json:"pro_expdateyn"`
	ProBuyPack         zero.Int    `db:"Pro_BuyPack" json:"pro_buypack"`
	ProSellUnit        zero.Int    `db:"Pro_SellUnit" json:"pro_sellunit"`
	ProSellPack        zero.Int    `db:"Pro_SellPack" json:"pro_sellpack"`
	ProMedUnit         zero.Int    `db:"Pro_MedUnit" json:"pro_medunit"`
	ProMedPack         zero.Int    `db:"Pro_MedPack" json:"pro_medpack"`
	ProBarcodeYN       zero.String `db:"Pro_BarcodeYN" json:"pro_barcodeyn"`
	ProMinOrdPO        zero.Float  `db:"Pro_MinOrdPO" json:"pro_minordpo"`
	ProKelipatan       zero.Float  `db:"Pro_Kelipatan" json:"pro_kelipatan"`
	ProNoRegYN         zero.String `db:"Pro_NoRegYN" json:"pro_noregyn"`
	ProNoReg           zero.String `db:"Pro_NoReg" json:"pro_noreg"`
	ProNIEDate         zero.String `db:"Pro_NIEDate" json:"pro_niedate"`
	ProKetNIE          zero.String `db:"Pro_KetNIE" json:"pro_ketnie"`
	ProNieSubmitBPOM   zero.String `db:"Pro_NieSubmitBPOM" json:"pro_niesubmitbpom"`
	ProHrgSpecialYN    zero.String `db:"Pro_HrgSpecialYN" json:"pro_hrgspecialyn"`
	ProHrgSpecial      zero.Float  `db:"Pro_HrgSpecial" json:"pro_hrgspecial"`
	ProHETYN           zero.String `db:"Pro_HETYN" json:"pro_hetyn"`
	ProHETPrice        zero.Float  `db:"Pro_HETPrice" json:"pro_hetprice"`
	ProBrandCode       zero.String `db:"Pro_BrandCode" json:"pro_brandcode"`
	ProNPDYN           zero.String `db:"Pro_NPDYN" json:"pro_npdyn"`
	ProTimbangYN       zero.String `db:"Pro_TimbangYN" json:"pro_timbangyn"`
	ProDecorYN         zero.String `db:"Pro_DecorYN" json:"pro_decoryn"`
	ProTabletYN        zero.String `db:"Pro_TabletYN" json:"pro_tabletyn"`
	ProMonograf        zero.String `db:"Pro_Monograf" json:"pro_monograf"`
	ProPatenYN         zero.String `db:"Pro_PatenYN" json:"pro_patenyn"`
	ProImportYN        zero.String `db:"Pro_ImportYN" json:"pro_importyn"`
	ProPackYN          zero.String `db:"Pro_PackYN" json:"pro_packyn"`
	ProKlasiBPOM       zero.String `db:"Pro_KlasiBPOM" json:"pro_klasibpom"`
	ProClass           zero.Int    `db:"Pro_Class" json:"pro_class"`
	ProIPA             zero.String `db:"Pro_IPA" json:"pro_ipa"`
	ProValueIPA        zero.Float  `db:"Pro_ValueIPA" json:"pro_valueipa"`
	ProIDA             zero.String `db:"Pro_IDA" json:"pro_ida"`
	ProValueIDA        zero.Float  `db:"Pro_ValueIDA" json:"pro_valueida"`
	ProGenerikID       zero.String `db:"Pro_GenerikID" json:"pro_generikid"`
	ProTerapiID        zero.String `db:"Pro_TerapiID" json:"pro_terapiid"`
	ProStrength        zero.Int    `db:"Pro_Strength" json:"pro_strength"`
	ProDsgForm         zero.Int    `db:"Pro_DsgForm" json:"pro_dsgform"`
	ProAmount          zero.Int    `db:"Pro_Amount" json:"pro_amount"`
	ProUnit            zero.Int    `db:"Pro_Unit" json:"pro_unit"`
	ProKdGenerik       zero.String `db:"Pro_KdGenerik" json:"pro_kdgenerik"`
	ProFirstActivity   zero.String `db:"Pro_FirstActivity" json:"pro_firstactivity"`
	ProSurveyYN        zero.String `db:"Pro_SurveyYN" json:"pro_surveyyn"`
	ProPrgDiscYN       zero.String `db:"Pro_PrgDiscYN" json:"pro_prgdiscyn"`
	ProCityCodeHJA     zero.Int    `db:"Pro_CityCodeHJA" json:"pro_citycodehja"`
	ProSalePrice       zero.Float  `db:"Pro_SalePrice" json:"pro_saleprice"`
	ProSalePriceBox    zero.Float  `db:"Pro_SalePriceBox" json:"pro_salepricebox"`
	ProSalePriceNon    zero.Float  `db:"Pro_SalePriceNon" json:"pro_salepricenon"`
	ProSalePriceNonBox zero.Float  `db:"Pro_SalePriceNonBox" json:"pro_salepricenonbox"`
	ProHPP             zero.Float  `db:"Pro_HPP" json:"pro_hpp"`
	ProScore           zero.Float  `db:"Pro_Score" json:"pro_score"`
	ProMarginComp      zero.Float  `db:"Pro_MarginComp" json:"pro_margincomp"`
	ProKeterangan      zero.Int    `db:"Pro_Keterangan" json:"pro_keterangan"`
	ProBuyGrp          zero.Float  `db:"Pro_BuyGrp" json:"pro_buygrp"`
	ProBuyGrpBox       zero.Float  `db:"Pro_BuyGrpBox" json:"pro_buygrpbox"`
	ProActiveYN        zero.String `db:"Pro_ActiveYN" json:"pro_activeyn"`
	ProUserID          zero.String `db:"Pro_UserID" json:"pro_userid"`
	ProLastUpdate      zero.String `db:"Pro_LastUpdate" json:"pro_lastupdate"`
	ProDataAktifYN     zero.String `db:"Pro_DataAktifYN" json:"pro_dataaktifyn"`
	ProHomeBrandYN     zero.String `db:"Pro_HomeBrandYN" json:"pro_homebrandyn"`
	ProLeadTime        zero.Int    `db:"Pro_LeadTime" json:"pro_leadtime"`
	ProBufferYN        zero.String `db:"Pro_BufferYN" json:"pro_bufferyn"`
	ProStorageTemp     zero.Int    `db:"Pro_StorageTemp" json:"pro_storagetemp"`
	ProClsProduct      zero.String `db:"Pro_ClsProduct" json:"pro_clsproduct"`
	ProClsMargin       zero.String `db:"Pro_ClsMargin" json:"pro_clsmargin"`
	ProInHealth        zero.String `db:"Pro_InHealth" json:"pro_inhealth"`
	ProCodeSup         zero.String `db:"Pro_CodeSup" json:"pro_codesup"`
	ProPointMedPack    zero.Float  `db:"Pro_PointMedPack" json:"pro_pointmedpack"`
	ProPointSellPack   zero.Float  `db:"Pro_PointSellPack" json:"pro_pointsellpack"`
	ProPointDyn        zero.Float  `db:"Pro_PointDyn" json:"pro_pointdyn"`
	ProPointDynX       zero.Float  `db:"Pro_PointDynX" json:"pro_pointdynx"`
	ProPointDynO       zero.Float  `db:"Pro_PointDynO" json:"pro_pointdyno"`
	ProPointDynV       zero.Float  `db:"Pro_PointDynV" json:"pro_pointdynv"`
	ProBundleYN        zero.String `db:"Pro_BundleYN" json:"pro_bundleyn"`
	ProBolehBundleYN   zero.String `db:"Pro_BolehBundleYN" json:"pro_bolehbundleyn"`
	ProNonGenerikYN    zero.String `db:"Pro_NonGenerikYN" json:"pro_nongenerikyn"`
	ProNomorIjinPro    zero.String `db:"Pro_NomorIjinPro" json:"pro_nomorijinpro"`
	ProKategoriPro     zero.String `db:"Pro_KategoriPro" json:"pro_kategoripro"`
	ProHalalYN         zero.String `db:"Pro_HalalYN" json:"pro_halalyn"`
	ProNoSertHalal     zero.String `db:"Pro_NoSertHalal" json:"pro_noserthalal"`
	ProHalalDate       zero.String `db:"Pro_HalalDate" json:"pro_halaldate"`
	ProHalal           zero.String `db:"Pro_Halal" json:"pro_halal"`
}

// JSONTerima ...
type JSONTerima struct {
	Data      MstProduct   `json:"data"`
	DataArray []MstProduct `json:"data_array"`
}
